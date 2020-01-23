package zbox

import "github.com/0chain/gosdk/zboxcore/sdk"

func GetClientEncryptedPublicKey() (string, error) {
	return sdk.GetClientEncryptedPublicKey()
}
