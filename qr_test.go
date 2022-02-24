package fatoora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQRCode_TLV(t *testing.T) {
	for name, tc := range map[string]struct {
		sellerName            string
		vatRegistrationNumber string
		invoiceTimestamp      string
		invoiceTotal          string
		vatTotal              string
		encoded               []byte
	}{
		"1": {
			"Bobs Records",
			"310122393500003",
			"2022-04-25T15:30:00Z",
			"1000.00",
			"150.00",
			hexToBytes("010c426f6273205265636f726473020F3331303132323339333530303030330314323032322d30342d32355431353a33303a30305a0407313030302e303005063135302e3030"),
		},
	} {
		qr := NewQRData(
			tc.sellerName,
			tc.vatRegistrationNumber,
			tc.invoiceTimestamp,
			tc.invoiceTotal,
			tc.vatTotal,
		)

		assert.Equal(t, tc.encoded, qr.TLV(), "%s: encoded data is invalid", name)
	}
}

func TestQRCode_Base64(t *testing.T) {
	for name, tc := range map[string]struct {
		sellerName            string
		vatRegistrationNumber string
		invoiceTimestamp      string
		invoiceTotal          string
		vatTotal              string
		encoded               string
	}{
		"1": {
			"Bobs Basement Records",
			"100025906700003",
			"2022-04-25T15:30:00Z",
			"2100100.99",
			"315015.15",
			"ARVCb2JzIEJhc2VtZW50IFJlY29yZHMCDzEwMDAyNTkwNjcwMDAwMwMUMjAyMi0wNC0yNVQxNTozMDowMFoECjIxMDAxMDAuOTkFCTMxNTAxNS4xNQ==",
		},
		"2": {
			"Saudi Business Machines",
			"100025906700003",
			"2021-12-15T15:00:00Z",
			"103.50",
			"13.50",
			"ARdTYXVkaSBCdXNpbmVzcyBNYWNoaW5lcwIPMTAwMDI1OTA2NzAwMDAzAxQyMDIxLTEyLTE1VDE1OjAwOjAwWgQGMTAzLjUwBQUxMy41MA==",
		},
		"3": {
			"مغسلتي",
			"000012321132",
			"2022-01-27T13:20:54Z",
			"3.45",
			"0.45",
			"AQzZhdi62LPZhNiq2YoCDDAwMDAxMjMyMTEzMgMUMjAyMi0wMS0yN1QxMzoyMDo1NFoEBDMuNDUFBDAuNDU=",
		},
	} {
		qr := NewQRData(
			tc.sellerName,
			tc.vatRegistrationNumber,
			tc.invoiceTimestamp,
			tc.invoiceTotal,
			tc.vatTotal,
		)

		assert.Equal(t, tc.encoded, qr.Base64(), "%s: encoded data is invalid", name)
	}
}
