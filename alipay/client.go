package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/teatak/config/sections"
)

type client struct {
	gateway    string
	appID      string
	privateKey string
	publicKey  string
}

//创建默认client
func NewDefault() *client {
	d := sections.Alipay["default"]
	return &client{
		gateway:    d.Gateway,
		appID:      d.AppID,
		privateKey: d.PrivateKey,
		publicKey:  d.PublicKey,
	}
}

//创建client
func New(gateway, appID, privateKey, publicKey string) *client {
	return &client{
		gateway:    gateway,
		appID:      appID,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

//Excute
func (s *client) Excute(request Request) *Result {
	result := &Result{
		err:  nil,
		data: []byte{},
	}
	p, err := s.buildValues(request)
	if err != nil {
		result.err = err
		return result
	}
	buf := strings.NewReader(p.Encode())
	req, err := http.NewRequest("POST", s.gateway, buf)
	if err != nil {
		result.err = err
		return result
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := http.DefaultClient.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		result.err = err
		return result
	}

	//verify data
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.err = err
		return result
	}
	if len(s.publicKey) > 0 {
		var dataStr = string(data)
		var rootNodeName = strings.Replace(request.APIName(), ".", "_", -1) + "_response"
		var rootIndex = strings.LastIndex(dataStr, rootNodeName)
		var errorIndex = strings.LastIndex(dataStr, "error_response")
		var content string
		var sign string
		if rootIndex > 0 {
			content, sign = parserJSONSource(dataStr, rootNodeName, rootIndex)
		} else if errorIndex > 0 {
			content, sign = parserJSONSource(dataStr, "error_response", errorIndex)
		} else {
			result.err = errors.New("error format")
			return result
		}
		if ok, err := verifyResponseData([]byte(content), sign, s.publicKey); !ok {
			result.err = err
			return result
		}
	}
	result.data = data
	return result
}

func (s *client) buildValues(request Request) (value url.Values, err error) {
	var p = url.Values{}
	p.Add("app_id", s.appID)
	p.Add("method", request.APIName())
	p.Add("format", "JSON")
	p.Add("charset", "utf-8")
	p.Add("sign_type", "RSA2")
	p.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Add("version", "1.0")

	bytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	p.Add("biz_content", string(bytes))

	var ps = request.Params()

	for key, value := range ps {
		p.Add(key, value)
	}

	sign, err := sign(p, []byte(s.privateKey))
	if err != nil {
		return nil, err
	}

	p.Add("sign", sign)
	return p, nil
}

func sign(param url.Values, privateKey []byte) (s string, err error) {
	var keys = make([]string, 0)
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var p = make([]string, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			p = append(p, key+"="+value)
		}
	}
	var src = strings.Join(p, "&")

	sig, err := signPKCS1v15([]byte(src), privateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func parserJSONSource(rawData string, nodeName string, nodeIndex int) (content string, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\"sign\"")
	var dataEndIndex = signIndex - 1

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	var signStartIndex = signIndex + len("sign") + 4
	sign = rawData[signStartIndex:]
	var signEndIndex = strings.LastIndex(sign, "\"}")
	sign = sign[:signEndIndex]

	return content, sign
}

func verifyResponseData(data []byte, sign string, key string) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	err = verifyPKCS1v15(data, signBytes, []byte(key), crypto.SHA256)
	if err != nil {
		return false, err
	}
	return true, nil
}

func signPKCS1v15(src, key []byte, hash crypto.Hash) ([]byte, error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var pri *rsa.PrivateKey
	pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed)
}

func verifyPKCS1v15(src, sig, key []byte, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
}

type Request interface {
	// 用于提供访问的 method
	APIName() string
	// 返回参数列表
	Params() map[string]string
}

type Response interface {
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type Result struct {
	err  error
	data []byte
}

func (s *Result) Decode(v interface{}) error {
	if s.err != nil {
		return s.err
	}
	err := json.Unmarshal(s.data, v)
	return err
}
