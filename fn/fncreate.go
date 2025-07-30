package fn

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
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

	id,_:=peer.IDFromPrivateKey(priv)

	groupKey:=GroupKey{
	GroupName: flag,
	PeerId: id.String(),
	PrivKey: encoded,
}
	

	data,_:=json.MarshalIndent(groupKey,"","")

	err:=os.WriteFile("GroupKey.json",data, 0600);
	if err!=nil{
		fmt.Println("erreur en écriture pour le fichier json")
	}




	return fmt.Sprintf("Le groupe %s a été crée avec succès", flag)
}