package alipay

import (
	"testing"
)

const (
	appID      = "2021003144638015"
	gateway    = "https://openapi.alipay.com/gateway.do"
	privateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEoQIBAAKCAQEAkUYqcdB9kxVDGNUq6h9mNBxW3vFu5STlglZddCBeP8iDERpz7FzTKJZly05rufnR8SiUD+UTCLpobG/Kk3BfYDTfTUsWwm5u6Nm7u38jIBuyI1MGGzMvvOUu2VaYx5NHbHxgGS9JU8v+Ndyk2qyM1gA7fK0j0Pr66zFsQ+3K7SRlPO4/eo19iDXYwCoRRvpGrZPv+Zb9238m1E9//V5imf9jjhkVbEfPisTLYfVBWOxIFvm9f6dato14dEM3XIQSdT1uhKWOKA9PP3lae522+udD2Sqo3VZY6o/QWJPZQplJTS/I9t6xq24IOmMH68FT5K1qRAg4a5uebRCNt9XbaQIDAQABAoIBAD7rBLcGdVyeeNDogg2kS7+GBpnINx9HO5+XiCIjYq59SUa/DjXwfgO2H6BN9po1eJfZ20aHEBdXpKJSSIpSwPwJ6iyvuMu8yIxQwN4T0kK1vBFQg/SD1skY+ZFn0AOf5e82buh5cmi4kQEC7uNGmDiXISx+UGTSTppObXV+J8VMsXuVzhB6ZXW5wFDWN0HcsyyW6mgZVBAwmcQ+VmsUZxzBnOGCe1KkYZCVWsIx2I0M+dNK1XEiucXNRsRbQuwIYQDFSgAHWq6QgdbwkxGQBdnX373hQaAP4oibYuoFXpw0nq/ySpVIMeAFgBzbZMzvoY7PO5njoM4jGNVXwzf/wAECgYEA+TEsBHFO6qBy0XcBPKT4dKUZjQ2+1RZZNbmhjPfX6P3blLKTMU27PYAK0nG22JsVsa/Q8QMevRW48ppccvsfixce2UjH5qIdC9eLZjBS27ZnGW/vJ+ObdBmrLhtfXq5RBtcK0EC1IHHh9KmpTbEqnQixZvrpbYsnDm/50nV2FSkCgYEAlT4zR62tffrRQqUGH75mW1RK2K8zry2Vi5zqUVAttOjVEvEsvPdM5K9RZGD6WAMOOlU5K8PkS9+vLapDDsbh9T8uXApbP7ae/qYamVyU0aeD+bpS4s2z8ZUrPXbaWodo60JUWnuxcImSItcwC0EK5Q0vlKoZ7xTyn6RVt116HEECgYEAzCpEdTNcxBFHYWUOZ2CGjBNE9vN/GdkgqEr6GKRriuKKyK7yuB+GvGGu/DLc62VNBfMVzpGO2r70dBiSjEZB1tTPmpjt54GFKNGlzcjj6k3s9MNn/5eVIy9wt+sHXWyWyHLh0jRh8j48MhfbteLlFhosN0J2hCDRlBucZJmGcNkCgYBsqtCd5WkI7OMHuq8d1/4hi9u/sLpaWD+mW4gF0vKk4k+bLj5cDVlQvhLumupNwVPPx5QMOON3F5UvE+Ul0kxezTVNUUElwCw+0OVXO6Ekkuu2nyOQ/ySByBKm+258MlKD0lUnWR9XCOR+N99EooRzSUmW3tBpVABhmz/KdubdwQJ/NTKMNa1Fo9M6+wxGSjfcKeluCO6aqAqlTKi3eNBrfHWh+xIDztDBx2b8zHMlCQxy2Yoschm2h6CD2t33w0zw7JlesjUsYwUptjFOmA2AFwuwxM7c6V/tVbzHkuLNqtPhKFLi/0F98mAQyaKZkLStFrqOtGKA3jTnQ8Gw6Sc2Ww==\n-----END RSA PRIVATE KEY-----"
	publicKey  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoIfKR77varhJ9wf6tVmLeqqca+HIPDOawH7idVwKLq4R02v8jNQDfYaEHYf6sW4xubHwJHZpbjcSz9R2p41Rd1qvrZKLZsahWFKQuJlkDCOjV/SWqYGMQLoHhs1bet6IQZpcNpvE2ZRMIwoApLzLx03z1mbzQwBZaBlke/zkZADkD0WepFvPMfATpM7Yz+dVnWLdedi0k/LwPXQRP8bop3tjSQBO08Kv/cDyeI+KRpIIeN7jYrgdvXpc1ZL0CJo35aIRBIJTJLgO7GesjCJ4O0PiyAf4NM/zBfAi5BODqKNjjnEYgO0Eo2N6HriNYUfx2ws8x4CoFnTw9UA/NoIWEQIDAQAB\n-----END PUBLIC KEY-----"
)

func TestAuth(t *testing.T) {
	client := New(gateway, appID, privateKey, publicKey)
	request := &OauthTokenRequest{
		GrantType: "authorization_code",
		Code:      "a3af0ec1143e42868c567d72b22bHX46",
	}
	response := &OauthTokenResponse{}
	err := client.Excute(request).Decode(response)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
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
		t.Fatal(err)
	}
	t.Log(response)
}

func TestTeadeQuery(t *testing.T) {
	client := New(gateway, appID, privateKey, publicKey)
	request := &TradeQueryRequest{

		OutTradeNo: "3",
	}
	response, err := client.TradeQuery(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
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
		t.Fatal(err)
	}
	t.Log(data)
}

// p.NotifyURL = "http://203.86.24.181:3000/alipay"
// 	p.Body = "body"
// 	p.Subject = "商品标题"
// 	p.OutTradeNo = "01010101"
// 	p.TotalAmount = "100.00"
// 	p.ProductCode = "p_1010101"
