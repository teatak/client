package alipay

type GoodsDetailItem struct {
	GoodsId       string `json:"goods_id"`
	AliPayGoodsId string `json:"alipay_goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category"`
	Body          string `json:"body"`
	ShowUrl       string `json:"show_url"`
}

func (s *client) TradeAppPay(request *TradeAppPayRequest) (string, error) {
	p, err := s.buildValues(request)
	if err != nil {
		return "", err
	}
	return p.Encode(), nil
}

func (s *client) TradeCreate(request *TradeCreateRequest) (*TradeCreateResponse, error) {
	response := &TradeCreateResponse{}
	err := s.Excute(request).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *client) TradeQuery(request *TradeQueryRequest) (*TradeQueryResponse, error) {
	response := &TradeQueryResponse{}
	err := s.Excute(request).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
