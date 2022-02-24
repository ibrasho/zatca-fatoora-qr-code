package fatoora

import (
	"bytes"
	"encoding/base64"
)

type QRData struct {
	Tags []Tag
}

func NewQRData(sellerName, vatRegistrationNumber, invoiceTimestamp, invoiceTotal, vatTotal string) *QRData {
	return &QRData{
		Tags: []Tag{
			NewSellerNameTag(sellerName),
			NewVATRegistrationNumberTag(vatRegistrationNumber),
			NewInvoiceTimestampTag(invoiceTimestamp),
			NewInvoiceTotalTag(invoiceTotal),
			NewVATTotalTag(vatTotal),
		},
	}
}

func (c *QRData) TLV() []byte {
	b := &bytes.Buffer{}

	for _, t := range c.Tags {
		b.Write(t.TLV())
	}

	return b.Bytes()
}

func (c *QRData) Base64() string {
	return base64.StdEncoding.EncodeToString(c.TLV())
}
