# ZATCA FATOORA QR Code

A Go library to encode data for ZATCA FATOORA-compliant QR Codes.

## Documentation

[![godoc](https://godoc.org/github.com/ibrasho/zatca-fatoora-qr-code?status.png)](https://godoc.org/github.com/ibrasho/zatca-fatoora-qr-code)

## Example Code

```go
import "github.com/ibrasho/zatca-fatoora-qr-code"

func main() {
	qr := fatoora.NewQRData(
		sellerName,
		vatRegistrationNumber,
		invoiceTimestamp,
        strconv.FormatFloat(invoiceTotal, 'f', 2, 64),
        strconv.FormatFloat(vatTotal, 'f', 2, 64),
	)

    // return LTV-encoded []byte
    qr.LTV()

    // return base64-encoded string (use this to generate the QR)
    qr.Base64()
}
```
