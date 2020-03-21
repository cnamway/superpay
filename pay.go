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
	sign     string
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
	p.sign = sign(data, p.key)
	p.query = mapToUrl(data, p.key)
	fmt.Println(p.query)
	rsp, err := http.Get(fmt.Sprintf("%s/testSign?%s&sign=%s", Url, p.query, p.sign))
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
