package zbox

import (
	"encoding/hex"
	"encoding/json"

	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
)

// GetClientEncryptedPublicKey - getting client encrypted pub key
func GetClientEncryptedPublicKey() (string, error) {
	return sdk.GetClientEncryptedPublicKey()
}

// Encrypt - encrypting text with key
func Encrypt(key, text string) (string, error) {
	keyBytes := []byte(key)
	textBytes := []byte(text)
	response, err := zboxutil.Encrypt(keyBytes, textBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(response), nil
}

// Decrypt - decrypting text with key
func Decrypt(key, text string) (string, error) {
	keyBytes := []byte(key)
	textBytes, _ := hex.DecodeString(text)
	response, err := zboxutil.Decrypt(keyBytes, textBytes)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// GetNetwork - get current network
func GetNetwork() (string, error) {
	networkDetails := sdk.GetNetwork()
	networkDetailsBytes, err := json.Marshal(networkDetails)
	if err != nil {
		return "", err
	}
	return string(networkDetailsBytes), nil
}

// GetBlobbers - get list of blobbers
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

// Sign - sign hash
func Sign(hash string) (string, error) {
	return client.Sign(hash)
}

// VerifySignature - verify message with signature
func VerifySignature(signature string, msg string) (bool, error) {
	return client.VerifySignature(signature, msg)
}
