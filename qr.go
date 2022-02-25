package fatoora

import (
	"bytes"
	"encoding/base64"

	qr "github.com/skip2/go-qrcode"
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

func (c *QRData) Base64PNG() (string, error) {
	png, err := qr.Encode(c.Base64(), qr.Medium, 512)
	if err != nil {
		return "", err
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(png), nil
}
