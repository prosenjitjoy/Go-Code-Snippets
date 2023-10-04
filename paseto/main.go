package main

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

func main() {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))

	token.SetString("user-id", uuid.NewString())

	key := paseto.NewV4SymmetricKey()
	fmt.Println(key.ExportHex())

	encrypted := token.V4Encrypt(key, nil)
	fmt.Println(encrypted)

	parser := paseto.NewParser()
	newToken, err := parser.ParseV4Local(key, encrypted, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newToken.ClaimsJSON()))

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public()
	fmt.Println(publicKey.ExportHex())

	signed := token.V4Sign(secretKey, nil)
	fmt.Println(signed)

	publicKey, err = paseto.NewV4AsymmetricPublicKeyFromHex("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	if err != nil {
		fmt.Println(err)
	}
	signed = "v4.public.eyJkYXRhIjoidGhpcyBpcyBhIHNpZ25lZCBtZXNzYWdlIiwiZXhwIjoiMjAyMi0wMS0wMVQwMDowMDowMCswMDowMCJ9v3Jt8mx_TdM2ceTGoqwrh4yDFn0XsHvvV_D0DtwQxVrJEBMl0F2caAdgnpKlt4p7xBnx1HcO-SPo8FPp214HDw.eyJraWQiOiJ6VmhNaVBCUDlmUmYyc25FY1Q3Z0ZUaW9lQTlDT2NOeTlEZmdMMVc2MGhhTiJ9"

	parser = paseto.NewParserWithoutExpiryCheck()
	newToken, err = parser.ParseV4Public(publicKey, signed, nil)
	fmt.Println(newToken)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newToken.ClaimsJSON()))
}
