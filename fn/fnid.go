package fn

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

type NewID struct{
	PeerID string
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
	privBytes,_ :=crypto.MarshalPrivateKey(priv)
	encoded:=base64.StdEncoding.EncodeToString(privBytes)
	id, _ := peer.IDFromPrivateKey(priv)



	newId:=NewID{
		PeerID: id.String(),
		priv : encoded,
	}

	file,_:=filepath.Glob(filepath.Join(basedir,"id_*.json"))
	increment:=len(file)+1

	data,_:=json.MarshalIndent(newId,"","  ")
	os.WriteFile(filepath.Join(basedir,"id_"+strconv.Itoa(increment)+".json"),data,0600)

	
	return newId.PeerID 
}