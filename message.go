package superpay

import (
	"encoding/json"
	"github.com/shop-r1/royalpay"
	"reflect"
)

type CreateOrderRequest struct {
	BizMchId          string            `json:"bizMchId"`
	MchId             string            `json:"mchId"`
	Desc              string            `json:"desc"`
	Money             string            `json:"money"`
	Currency          royalpay.Currency `json:"currency"`
	ThirdpartyTrxId   string            `json:"thirdparty_trx_id"`
	Callback          string            `json:"callback"`
	ServerCallbackUrl string            `json:"serverCallbackUrl"`
}

func (e CreateOrderRequest) toMap() map[string]string {
	elem := reflect.ValueOf(&e).Elem()
	return relFromTagJson(elem)

}

func relFromTagJson(elem reflect.Value) map[string]string {
	m := make(map[string]string)
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return m
}

type CreateOrderResponse struct {
	RspData
	ThirdpartyTrxId string `json:"thirdparty_trx_id"`
	PartnerTransId  string `json:"partner_trans_id"`
	QrcodeContent   string `json:"qrcodeContent"`
}

type RspData struct {
	Success   bool        `json:"success"`
	ErrorCode json.Number `json:"error_code"`
	Message   string      `json:"message"`
}
