package superpay

import (
	"encoding/json"
	"fmt"
	"github.com/shop-r1/royalpay"
	"io/ioutil"
	"net/http"
)

const Url = "https://tradeserver.supaytechnology.com/v1"

type Pay struct {
	bizMchId string
	mchId    string
	key      string
	query    string
}

//创建订单
func NewPay(bizMchId, mchId, key string) *Pay {
	return &Pay{
		bizMchId: bizMchId,
		mchId:    mchId,
		key:      key,
	}
}

func (p *Pay) CreateOrder(money, desc, thirdpartyTrxId, callback, serverCallbackUrl string, currency royalpay.Currency) (*CreateOrderResponse, error) {
	req := CreateOrderRequest{
		BizMchId:          p.bizMchId,
		MchId:             p.mchId,
		Desc:              desc,
		Money:             money,
		Currency:          currency,
		ThirdpartyTrxId:   thirdpartyTrxId,
		Callback:          callback,
		ServerCallbackUrl: serverCallbackUrl,
	}
	data := req.toMap()
	p.query = sign(data, p.key)
	rsp, err := http.Get(fmt.Sprintf("%s/testSign?%s", Url, p.query))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var rb []byte
	rb, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var r CreateOrderResponse
	err = json.Unmarshal(rb, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
