package fn

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

type GroupKey struct{
	GroupName string `json:"group_name"`
	PeerId string `json:"peer_id"`
	PrivKey string`json:"priv_key"`
	RelayAddr string `json:"relay_addr"`
}

type Invite struct{
	GroupName string `json:"group_name"`
	PeerId string `json:"peer_id"`
	Members []string`json:"members,omitempty"`
	RelayAddr string`json:"relay_addr"`
}

func FnCreate(groupname string, relayAddr string) string {


	basedir:=`C:/pulse_test`
	err := os.MkdirAll(basedir, 0o755); 
	if err != nil {
		return "erreur création dossier"
	}

	priv,_, _:= crypto.GenerateKeyPair(crypto.Ed25519, -1)
	privBytes,_ :=crypto.MarshalPrivateKey(priv)
	encoded:=base64.StdEncoding.EncodeToString(privBytes)
	id,_:=peer.IDFromPrivateKey(priv)

	groupKey:=GroupKey{
	GroupName: groupname,
	PeerId: id.String(),
	PrivKey: encoded,
	RelayAddr: "voici le relay_addr",
}
	
	
	dataGK,_:=json.MarshalIndent(groupKey,"","")
	
	err=os.WriteFile(filepath.Join(basedir,"GroupKey_"+groupname+".json") ,dataGK, 0600);
	if err!=nil{
		fmt.Println("erreur en écriture pour le fichier json")
	}


	invite:=Invite{
		GroupName: groupname,
		PeerId: id.String(),
		Members: []string{id.String(),"ajouter_un_membre_ici"},
		RelayAddr: "voici le relay",
	
	}

	dataIV,_:=json.MarshalIndent(invite,"","")

	err=os.WriteFile(filepath.Join(basedir,"Invite_"+groupname+".json"),dataIV,0600)
	if err!=nil{
		fmt.Println("Erreur en écriture pour le fichier json")
	}
	

	return fmt.Sprintf("Le groupe %s a été crée avec succès", groupname)
}