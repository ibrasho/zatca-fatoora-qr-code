package fatoora

import (
	"bytes"
)

type TagType int

const (
	SellerName            TagType = 1
	VATRegistrationNumber TagType = 2
	InvoiceTimestamp      TagType = 3
	InvoiceTotal          TagType = 4
	VATTotal              TagType = 5
	InvoiceHash           TagType = 6
	Signature             TagType = 7
	PublicKey             TagType = 8
	StampSignature        TagType = 9
)

type Tag struct {
	Type  TagType
	Value string
}

func NewTag(tag TagType, value string) Tag {
	return Tag{
		Type:  tag,
		Value: value,
	}
}

func (t Tag) TLV() []byte {
	b := &bytes.Buffer{}

	// fmt.Sprintf("%02X%02X%s", t.Type, t.Length(), t.Value)
	b.WriteByte(byte(t.Type))
	b.WriteByte(byte(len(t.Value)))
	b.Write([]byte(t.Value))

	return b.Bytes()
}

func NewSellerNameTag(v string) Tag            { return NewTag(SellerName, v) }
func NewVATRegistrationNumberTag(v string) Tag { return NewTag(VATRegistrationNumber, v) }
func NewInvoiceTimestampTag(v string) Tag      { return NewTag(InvoiceTimestamp, v) }
func NewInvoiceTotalTag(v string) Tag          { return NewTag(InvoiceTotal, v) }
func NewVATTotalTag(v string) Tag              { return NewTag(VATTotal, v) }
func NewInvoiceHashTag(v string) Tag           { return NewTag(InvoiceHash, v) }
func NewSignatureTag(v string) Tag             { return NewTag(Signature, v) }
func NewPublicKeyTag(v string) Tag             { return NewTag(PublicKey, v) }
func NewStampSignatureTag(v string) Tag        { return NewTag(StampSignature, v) }
