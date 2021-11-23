package utils

import (
	"fmt"
	b64 "encoding/base64"
	"encoding/base32"
	"crypto/sha512"
	"crypto/ed25519"
	"golang.org/x/crypto/hkdf"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)


func GenerateAccountList(seed string, index int, amount int)([]string){
	var accountArray []string;
	var decodedKey []byte;
	var err error;

	if amount == 0 {
		amount = 1
	}

	if len(seed) > 50 {
		masterKey := MasterToDerivation(seed)
		fmt.Println(masterKey)
		
		decodedKey, err = b64.StdEncoding.DecodeString(masterKey)
		
		if err != nil {
			fmt.Println(err )
		}
	} else {  
		decodedKey, err = b64.StdEncoding.DecodeString(seed)
		if err != nil {
			fmt.Println(err )
		}
	}
	
	for i := index; i < index + amount; i++ {
		info := []byte(fmt.Sprintf("AlgorandDeterministicKey-%d", i))

		hash := sha512.New512_256

		keystream := hkdf.Expand(hash, decodedKey, info)

		pub, priv, err := ed25519.GenerateKey(keystream)

		if err != nil{
			// Just to ignore the priv warning
			fmt.Println(priv)
			fmt.Println(err)
		}

		chksum := sha512.Sum512_256(pub[:])

		checksumAddress := append(pub[:], chksum[28:]...)
		
		accountArray = append(accountArray, base32.StdEncoding.EncodeToString(checksumAddress)[:58])
	}

	return accountArray
}


func GenerateAccountMnemonic(seed string, index int)([]string){
	var accountArray []string;
	var decodedKey []byte;
	var err error;

	if len(seed) > 50 {
		masterKey := MasterToDerivation(seed)

		decodedKey, err = b64.StdEncoding.DecodeString(masterKey)
		
		if err != nil {
			fmt.Println(err )
		}
	} else {  
		decodedKey, err = b64.StdEncoding.DecodeString(seed)
		if err != nil {
			fmt.Println(err )
		}
	}
	
	for i := index; i < index + 1; i++ {
		info := []byte(fmt.Sprintf("AlgorandDeterministicKey-%d", i))

		hash := sha512.New512_256

		keystream := hkdf.Expand(hash, decodedKey, info)

		pub, priv, err := ed25519.GenerateKey(keystream)

		if err != nil{
			// Just to ignore the priv warning
			fmt.Println(pub)
			fmt.Println(err)
		}

		mnemonic, err := mnemonic.FromPrivateKey(priv)
		accountArray = append(accountArray, mnemonic)
	}

	return accountArray
}