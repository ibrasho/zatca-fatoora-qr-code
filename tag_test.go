package fatoora

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag_Encode(t *testing.T) {
	tcs := []struct {
		tag     Tag
		encoded []byte
	}{
		{
			NewSellerNameTag("Bobs Records"),
			hexToBytes("010c426f6273205265636f726473"),
		},
		{
			NewVATRegistrationNumberTag("310122393500003"),
			hexToBytes("020F333130313232333933353030303033"),
		},
		{
			NewInvoiceTimestampTag("2022-04-25T15:30:00Z"),
			hexToBytes("0314323032322d30342d32355431353a33303a30305a"),
		},
		{
			NewInvoiceTotalTag("1000.00"),
			hexToBytes("0407313030302e3030"),
		},
		{
			NewVATTotalTag("150.00"),
			hexToBytes("05063135302e3030"),
		},
		{
			NewSellerNameTag("Bobs Basement Records"),
			hexToBytes("0115426F627320426173656D656E74205265636F726473"),
		},
		{
			NewVATRegistrationNumberTag("100025906700003"),
			hexToBytes("020f313030303235393036373030303033"),
		},
		{
			NewInvoiceTimestampTag("2022-04-25T15:30:00Z"),
			hexToBytes("0314323032322D30342D32355431353A33303A30305A"),
		},
		{
			NewInvoiceTotalTag("2100100.99"),
			hexToBytes("040a323130303130302E3939"),
		},
		{
			NewVATTotalTag("215015.15"),
			hexToBytes("05093231353031352E3135"),
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.encoded, tc.tag.TLV(), "encoded data is invalid")
	}
}

func hexToBytes(h string) []byte {
	data, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}

	return data
}
