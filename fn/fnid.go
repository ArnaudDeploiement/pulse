package fn

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

type NewID struct{
	PeerID string
	pub string
	priv string
}

func GetIdPerson() string {

	basedir:=`C:/pulse`
	err := os.MkdirAll(basedir, 0o755); 
	if err != nil {
		return "erreur création dossier"
	}
	

	//récupérer la pub et la priv, encoder pour json
	priv, _, _ := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	id, _ := peer.IDFromPrivateKey(priv)



	newId:=NewID{
		PeerID: id.String(),
	}

	data,_:=json.MarshalIndent(newId,"","  ")

	err=os.WriteFile(filepath.Join(basedir,"id.json"),data,0600)
	if err!=nil{
		fmt.Println("erreur d'écriture")
	}


	return id.String()
}