package alipay

import (
	"fmt"
	"testing"
)

const (
	appID      = "2017112200093974"
	gateway    = "https://openapi.alipay.com/gateway.do"
	privateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAjKdrqX5X4QNjYAkRtojrUDa9coD6NCzcjbjVUDlpbqVmG5rqxLYB0fqBVuFIhHzYgRG4WL5xlrr7lpBIAitG4LF7mOyGuxpkkifiOJ1/4B5JSDuoVM2i85TWZHrjTqXn6nZMUZLHLq+tlGkLrjn/W6hx6IHGqDG7aasQGzZ5n+WmYIMsOF/UcK2mNIbOSIDmp7NMKasoYywRjadPfSmMaUL8QY7YUtASXlF3b8CN/BdpGxeKEQgID3D7HNr/fMROvmOIXnQtJzGQ+BnecGehFnl6awn/NOBMP5r0F+H+uGbFM9N6iaiZUqJCVYvT/URbIgAlwHzW5WthI28pFFNinQIDAQABAoIBAEC0gFXP1pOgIGY8pdZoNICJPivf47p+7NORtSgb0UHiew+1+8yUwk+bw0Z5iwfP+zWdNkY1DkI+MKE6LLY6vHd/jpFyNiT65ktKZ8qNwhcTN7kwIKcqlo0h6mL9GXbD4eW9mCxqrpDuppbAoGv9KFk9K3G/yej0P/hDhrxOGjORHc/yYm3k3rpW88bnmwLP7lfNcUjaCsyhgBatDcFgTXcAUzBp/KIeQoJJMZMfQPqeLiUIVc+1KYvn5YVEzbBaP2cfhMyAzQWpPnPHT2dZkcSy9EbfuO9FGt0MsLuU2HsIgDVhS5KtFBcaEp0dizOiHAH/ThzfzaEVmMsVrjYaFQECgYEA0LpUKPoao1BJcDFkoSTTSvdNoVGE6+cYOe/8SS2cINUqhmc/YQpXjsfq7QzJSa3vdhlblG8pTq8oNmAa38QuuzoIjcs2P0gyCT8n1WL35LiUA7+IwKmhqn7SQKgolod/lumWt3HkE7AvjdoQVU3nN5rKTI9cRl4DMdmW2wuiQ90CgYEArIJKMrtgYkEttIpNEN5mNpHw8EqBatULm7NlTv3OuiLZyRXZU/nX8NsJgQ1XrLJb3+ziZE5EKnX3h6rfs/EABgRk7l1cQ331e/HsU+erAvfv1Nr4b7Uy+xiCjH5gBSAs9Yk8y+PSWykF6QHPrGzmn0kyXbhtEvrGQmNFnhT8DcECgYA3oWGq5jXRTyWzlEaIuGLkORNkqnBt21LWpkKsUk6pZYcG+V+oAnShpmDTuzOuwn/vDwUcg6ATy5VBIM4o04XM+sOq3v1fUmObmUyJj+4X9KXmddcB2nQvF8v71ZWzwPtdgnZcACvUn4GdIgB2a+PzKvs9+nuSeZ4KXD7btLcTtQKBgCZ7ukmu+0Vv/N41rBODSZIK5JD2TgCjHj/RuVvxythDRgpR0XWmWsHkWy5q55AjCdPeaHKv8wz13A5r1BO0Q6kT+eAi+54iz0yjH3FoUQIMhXgakSl6ZrEVOqpU1t7N2rm7r9BcNmUXtbXRLZ7tfFU0YwP9THfbcbcFplYpoEEBAoGBAL29pwZwzzvp++HtOKkFHgxlXQQVP9RyJpZRo80mu5GGwZVFC9Nr7lUKLNJjngUhj+AuhjiJIZKiMVA0w94vM7lrHek5D8owm1C1wbm9t23oPbjwdZV1BBy1QGAqXfpKP5ynssVJAZXRafEeDCD8PsSrxN76tNVZZd5dM2x68s1p\n-----END RSA PRIVATE KEY-----"
	publicKey  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiXYlHnIZsl1TBRxOgIT1wg/ultsl1fzyEIXA3QsAAICR/kOAO4+A2hK5/fi3470+vSwSVyVuYasgriqnJATsa/ByFIEmhNPlgN2d3CYM/j9nFGX4YJh4/V6Tr0Qjx3SvTMWea/zILEJ1ODkZZo2Am6hsLpK9tK2KMZA8dydePBf5CGZVvfLjs4Xn2z+OZqqKu7CtrJNRqb+WqPGOeEUhEwYYpoMElBWnAVs9sTr3TLp6yWS1HvxBz7YYSZ+B/OeujcqPL0Fije74R9u5HInZgtQKz0/R03IK4dOZZQzFxa///t6rcr0dXbg92v21br7/scOaHXHJ/UUPgF1OlDxsGQIDAQAB\n-----END PUBLIC KEY-----"
)

func TestAuth(t *testing.T) {
	// client := New(gateway, appID, privateKey, publicKey)
	// request := &OauthTokenRequest{
	// 	GrantType: "authorization_code",
	// 	Code:      "a3af0ec1143e42868c567d72b22bHX46",
	// }
	// response := &OauthTokenResponse{}
	// err := client.Excute(request).Decode(response)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println(response)
}

func TestTeadeCreate(t *testing.T) {
	client := New(gateway, appID, privateKey, publicKey)
	request := &TradeCreateRequest{
		BuyerId:        "2088702437003466",
		Subject:        "测试",
		OutTradeNo:     "3",
		TimeoutExpress: "30m",
		TotalAmount:    0.01,
	}

	response, err := client.TradeCreate(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
}

func TestTradeAppPay(t *testing.T) {
	client := New(gateway, appID, privateKey, publicKey)
	request := &TradeAppPayRequest{
		NotifyURL:      "http://222.86.24.181:3000/alipay",
		Subject:        "测试",
		OutTradeNo:     "5",
		TimeoutExpress: "30m",
		TotalAmount:    -0.01,
	}

	data, err := client.TradeAppPay(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

// p.NotifyURL = "http://203.86.24.181:3000/alipay"
// 	p.Body = "body"
// 	p.Subject = "商品标题"
// 	p.OutTradeNo = "01010101"
// 	p.TotalAmount = "100.00"
// 	p.ProductCode = "p_1010101"
