package zbox

import (
	"encoding/hex"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
)

func GetClientEncryptedPublicKey() (string, error) {
	return sdk.GetClientEncryptedPublicKey()
}

func Encrypt(key, text string) (string, error) {
	keyBytes := []byte(key)
	textBytes := []byte(text)
	response, err := zboxutil.Encrypt(keyBytes, textBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(response), nil
}

func Decrypt(key, text string) (string, error) {
	keyBytes := []byte(key)
	textBytes, _ := hex.DecodeString(text)
	response, err := zboxutil.Decrypt(keyBytes, textBytes)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
