package alipay

//统一收单交易创建接口
//https://docs.open.alipay.com/api_1/alipay.trade.create/

type TradeCreateRequest struct {
	NotifyURL    string `json:"notify_url,omitempty"`
	ReturnURL    string `json:"return_url,omitempty"`
	AppAuthToken string `json:"app_auth_token,omitempty"` // 可选

	//biz content，这四个参数是必须的
	Subject     string  `json:"subject"`      // 订单标题
	OutTradeNo  string  `json:"out_trade_no"` // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	TotalAmount float64 `json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode string  `json:"product_code"` // 销售产品码，与支付宝签约的产品码名称。 参考官方文档, App 支付时默认值为 QUICK_MSECURITY_PAY

	Body                string                 `json:"body,omitempty"`                  // 订单描述
	BusinessParams      string                 `json:"business_params,omitempty"`       // 商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，格式为json格式
	DisablePayChannels  string                 `json:"disable_pay_channels,omitempty"`  // 禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注，与enable_pay_channels互斥
	EnablePayChannels   string                 `json:"enable_pay_channels,omitempty"`   // 可用渠道，用户只能在指定渠道范围内支付  当有多个渠道时用“,”分隔 注，与disable_pay_channels互斥
	ExtendParams        map[string]interface{} `json:"extend_params,omitempty"`         // 业务扩展参数，详见下面的“业务扩展参数说明”
	AgreementSignParams interface{}            `json:"agreement_sign_params,omitempty"` // 签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。 周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。
	GoodsType           string                 `json:"goods_type,omitempty"`            // 商品主类型：0—虚拟类商品，1—实物类商品 注：虚拟类商品不支持使用花呗渠道
	InvoiceInfo         string                 `json:"invoice_info,omitempty"`          // 开票信息
	PassbackParams      string                 `json:"passback_params,omitempty"`       // 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	PromoParams         string                 `json:"promo_params,omitempty"`          // 优惠参数 注：仅与支付宝协商后可用
	RoyaltyInfo         string                 `json:"royalty_info,omitempty"`          // 描述分账信息，json格式，详见分账参数说明
	SellerId            string                 `json:"seller_id,omitempty"`             // 收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	SettleInfo          string                 `json:"settle_info,omitempty"`           // 描述结算信息，json格式，详见结算参数说明
	SpecifiedChannel    string                 `json:"specified_channel,omitempty"`     // 指定渠道，目前仅支持传入pcredit  若由于用户原因渠道不可用，用户可选择是否用其他渠道支付。  注：该参数不可与花呗分期参数同时传入
	StoreId             string                 `json:"store_id,omitempty"`              // 商户门店编号。该参数用于请求参数中以区分各门店，非必传项。
	SubMerchant         string                 `json:"sub_merchant,omitempty"`          // 间连受理商户信息体，当前只对特殊银行机构特定场景下使用此字段
	TimeoutExpress      string                 `json:"timeout_express,omitempty"`       // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	TimeExpire          string                 `json:"time_expire,omitempty"`           // 该笔订单绝对超时时间，格式为yyyy-MM-dd HH:mm:ss

	//trade create
	DiscountableAmount float64            `json:"discountable_amount,omitempty"` // 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位
	BuyerId            string             `json:"buyer_id"`
	GoodsDetail        []*GoodsDetailItem `json:"goods_detail,omitempty"`
	OperatorId         string             `json:"operator_id"`
	TerminalId         string             `json:"terminal_id"`
}

func (s *TradeCreateRequest) APIName() string {
	return "alipay.trade.create"
}

func (s *TradeCreateRequest) Params() map[string]string {
	var m = make(map[string]string)
	if s.AppAuthToken != "" {
		m["app_auth_token"] = s.AppAuthToken
	}
	return m
}

type TradeCreateResponse struct {
	TradeCreate struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		TradeNo    string `json:"trade_no"` // 支付宝交易号
		OutTradeNo string `json:"out_trade_no"`
	} `json:"alipay_trade_create_response,omitempty"`
	ErrorResponse `json:"error_response,omitempty"`
	Sign          string `json:"sign"`
}
