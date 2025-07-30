package fn

import (
	"fmt"
)

type GroupKey struct{
	GroupName string
	PeerId string
	PrivKey string
}


func FnCreate(flag string) string {
	
	//Créer le paire priv/pub


	//priv, pub, err:= crypto.GenerateEd25519Key(crypto.Ed25519, -1)
	
	//privBytes,_ :=crypto.MarshalPrivateKey(priv)
	//encoded:=base64.StdEncoding.EncodeToString(privBytes)

	//tmpl:=


	return fmt.Sprintf("Le groupe %s a été crée avec succès", flag)
}