package fn

import (
	"crypto/rand"
	"crypto/sha256"
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

type Protocol struct{
	Protocol string `json:"Protocol"`
	RelayAddr string`json:"relay_addr"`
	As string `json:"as"`
}

func FnCreate(groupname string, relayAddr string) string {
	basedir:=`C:/pulse_test`
	os.MkdirAll(basedir, 0o755); 
	
	priv,_, _:= crypto.GenerateKeyPair(crypto.Ed25519, -1)
	privBytes,_ :=crypto.MarshalPrivateKey(priv)
	encoded:=base64.StdEncoding.EncodeToString(privBytes)
	id,_:=peer.IDFromPrivateKey(priv)

	groupKey:=GroupKey{
	GroupName: groupname,
	PeerId: id.String(),
	PrivKey: encoded,
	RelayAddr: relayAddr,
}
	dataGK,_:=json.MarshalIndent(groupKey,"","  ")
	os.WriteFile(filepath.Join(basedir,"GroupKey_"+groupname+".json") ,dataGK, 0600);

	protocol:=Protocol{
		Protocol: deriveProtocolName(groupname,id.String()) ,
		RelayAddr: relayAddr,
		As:as(32),
	
	}

	dataIV,_:=json.MarshalIndent(protocol,"","  ")
	os.WriteFile(filepath.Join(basedir,"Protocol_"+groupname+".json"),dataIV,0600)
	
	return fmt.Sprintf("Le groupe %s a été crée avec succès.\nLe fichier Invite_%s.json a été créé avec succès dans %s", groupname,groupname,basedir)
}

func deriveProtocolName(groupname, peerid string) string{
	deriv:=sha256.Sum256([]byte(groupname+ ";"+ peerid))
	return "/pulse/"+base64.RawURLEncoding.EncodeToString(deriv[:16])+"/1.0"
}

func as(n int) string {
	b:=make([]byte,n)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}