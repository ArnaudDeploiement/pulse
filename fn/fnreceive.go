package fn

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	ma "github.com/multiformats/go-multiaddr"
)

func FnReceive(protocolPath, storeDir string)  {
	ctx := context.Background()

	// 1) Lire protocol.json
	data, _ := os.ReadFile(protocolPath)
	var cfg Protocol
	json.Unmarshal(data, &cfg)

	// 2) Préparer dépôt + keepalive
	os.MkdirAll(storeDir, 0o755)
	base := `C:\pulse\receivers`
	os.MkdirAll(base, 0o755)
	keep := filepath.Join(base, cfg.Protocol+".keep")
	os.WriteFile(keep, []byte("1"), 0o600)

	// 3) Créer l’hôte libp2p
	h, _ := libp2p.New()
	fmt.Println("📡 PeerID :", h.ID().String())

	// 4) Connexion au relay
	maddr, _ := ma.NewMultiaddr(cfg.RelayAddr)
	ri, _ := peer.AddrInfoFromP2pAddr(maddr)
	h.Connect(ctx, *ri)
	fmt.Println("✅ Connecté au relay")


	

	// 5) Handler réception

	 handler := func (s network.Stream) {
		defer s.Close()
		
		reader:=bufio.NewReader(s)
		filename,_:=reader.ReadString('\n')
		filename=strings.TrimSpace(filename)
		filename=filepath.Base(filename)


		path:=filepath.Join(storeDir,filename)
		f, _ := os.Create(path)
		defer f.Close()
		io.Copy(f, reader)

		fmt.Println("📥 Reçu →", path)
	}

	h.SetStreamHandler(protocol.ID(cfg.Protocol), handler)


	// 6) Watcher stop
	go func() {
		t := time.NewTicker(2 * time.Second)
		for range t.C {
			if _, err := os.Stat(keep); os.IsNotExist(err) {
				fmt.Println("🛑 Stop détecté")
				h.Close()
				exec.Command("taskkill", "/PID", fmt.Sprint(os.Getpid()), "/F").Run()
				return
			}
		}
	}()
	
	fmt.Println("👂 En écoute sur", cfg.Protocol, "→ dépôt :", storeDir)
    select {}

}
