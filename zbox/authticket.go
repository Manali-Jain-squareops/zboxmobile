package zbox

import (
	"github.com/0chain/gosdk/zboxcore/sdk"
)

type AuthTicket struct {
	sdkAuthTicket *sdk.AuthTicket
}

func InitAuthTicket(authTicket string) *AuthTicket {
	at := &AuthTicket{}
	at.sdkAuthTicket = sdk.InitAuthTicket(authTicket)
	return at
}

func (at *AuthTicket) IsDir() (bool, error) {
	return at.sdkAuthTicket.IsDir()
}

func (at *AuthTicket) GetFilename() (string, error) {
	return at.sdkAuthTicket.GetFileName()
}
