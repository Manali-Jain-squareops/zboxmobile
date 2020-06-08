package zbox

import (
	"encoding/hex"
	"encoding/json"

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

func GetNetwork() (string, error) {
	networkDetails := sdk.GetNetwork()
	networkDetailsBytes, err := json.Marshal(networkDetails)
	if err != nil {
		return "", err
	}
	return string(networkDetailsBytes), nil
}

func GetBlobbers() (string, error) {
	blobbers, err := sdk.GetBlobbers()
	if err != nil {
		return "", err
	}

	blobbersBytes, err := json.Marshal(blobbers)
	if err != nil {
		return "", err
	}
	return string(blobbersBytes), nil
}
