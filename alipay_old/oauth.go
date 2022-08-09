package alipay

//https://docs.open.alipay.com/api_9/alipay.system.oauth.token

type OauthTokenRequest struct {
	GrantType    string `json:"-"` // 必须 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"-"` // 可选 授权码，用户对应用授权后得到。
	RefreshToken string `json:"-"` // 可选 刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
	response     *OauthTokenResponse
}

func (s *OauthTokenRequest) Method() string {
	return "alipay.system.oauth.token"
}

func (s *OauthTokenRequest) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = s.GrantType
	if s.Code != "" {
		m["code"] = s.Code
	}
	if s.RefreshToken != "" {
		m["refresh_token"] = s.RefreshToken
	}
	return m
}

func (s *OauthTokenRequest) Name() string {
	return ""
}

func (s *OauthTokenRequest) JSON() string {
	return ""
}

func (s *OauthTokenRequest) GetResponse() Response {
	if s.response == nil {
		s.response = &OauthTokenResponse{}
	}
	return s.response
}

type OauthTokenResponse struct {
	OauthTokenResponse struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`       // 支付宝用户的唯一userId
		AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn    int    `json:"expires_in"`    // 访问令牌的有效时间，单位是秒。
		RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn  int    `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。

	} `json:"alipay_system_oauth_token_response,omitempty"`
	ErrorResponse `json:"error_response,omitempty"`
	Sign          string `json:"sign"`
}

func (s *OauthTokenResponse) IsSuccess() bool {
	return s.OauthTokenResponse.UserId != ""
}
