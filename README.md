# GoRoitel ( Roitel golang kütüphanesi www.roitel.com.tr )

## Özellikler
* Sms Gönderim

## Kurulum
```
go get github.com/semihs/goroitel
```

## Kullanım

### Sms Gönderim

```go
package main

import (
	"github.com/semihs/goroitel"
	"fmt"
)


func main() {
	roitelClient := goroitel.NewRoitelClient("kullanıcı adınız", "şifreniz", "sms başlığı", "karakter kodlaması (türkçe için turkish)")
	if err := roitelClient.SendSms("SMS HEADER", "5005005050", "mesajınız"); err != nil {
            fmt.Errorf("an error occurred while sms sending %s", err)
        }
}
```
