package superpay

import (
	"reflect"
	"testing"

	"github.com/shop-r1/royalpay"
)

func TestCreateOrderRequest_toMap(t *testing.T) {
	type fields struct {
		BizMchId          string
		MchId             string
		Desc              string
		Money             string
		Currency          royalpay.Currency
		ThirdpartyTrxId   string
		Callback          string
		ServerCallbackUrl string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			"test01",
			fields{
				BizMchId:          "2071",
				MchId:             "59810",
				Desc:              "KFC coffee tea",
				Money:             "12.11",
				Currency:          "AUD",
				ThirdpartyTrxId:   "165655566661",
				Callback:          "http://xxx.com/callbak",
				ServerCallbackUrl: "http://xxx.com/serverCallbackUrl.php",
			},
			map[string]string{
				"bizMchId":          "2071",
				"mchId":             "59810",
				"desc":              "KFC coffee tea",
				"money":             "12.11",
				"currency":          "AUD",
				"thirdparty_trx_id": "165655566661",
				"callback":          "http://xxx.com/callbak",
				"serverCallbackUrl": "http://xxx.com/serverCallbackUrl.php",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := CreateOrderRequest{
				BizMchId:          tt.fields.BizMchId,
				MchId:             tt.fields.MchId,
				Desc:              tt.fields.Desc,
				Money:             tt.fields.Money,
				Currency:          tt.fields.Currency,
				ThirdpartyTrxId:   tt.fields.ThirdpartyTrxId,
				Callback:          tt.fields.Callback,
				ServerCallbackUrl: tt.fields.ServerCallbackUrl,
			}
			if got := e.toMap(true); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.toMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
