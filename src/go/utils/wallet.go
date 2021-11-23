package utils

import(
	b64 "encoding/base64"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func MasterToDerivation(_mnemonic string) (string){
	key, err := mnemonic.ToMasterDerivationKey(_mnemonic)

    if err != nil{
		return err.Error()
	}

	return b64.StdEncoding.EncodeToString(key[:])
}

func DerivationToMaster(_key string)(string){
	var arr [32]byte
	copy(arr[:], _key)
	
	key, err := mnemonic.FromMasterDerivationKey(arr);
	
	if err != nil{
		return err.Error()
	}

	return key
}
