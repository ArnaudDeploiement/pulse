package fn

import (
	"encoding/base64"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
)

type GroupKey struct{
	GroupName string
	PeerId string
	PrivKey string
}


func FnCreate(flag string) string {
	priv,_, _:= crypto.GenerateKeyPair(crypto.Ed25519, -1)
	privBytes,_ :=crypto.MarshalPrivateKey(priv)
	encoded:=base64.StdEncoding.EncodeToString(privBytes)





	return fmt.Sprintf("Le groupe %s a été crée avec succès", flag)
}