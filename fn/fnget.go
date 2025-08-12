package fn

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
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

func FnGet(protocolPath, storeDir string)  {
	ctx := context.Background()


	cfg:=unmarshal(protocolPath)
	keep:=keep(cfg)

	os.MkdirAll(storeDir, 0o755)


	h, _ := libp2p.New()
	fmt.Println("📡 PeerID :", h.ID().String())

	maddr, _ := ma.NewMultiaddr(cfg.RelayAddr)
	ri, _ := peer.AddrInfoFromP2pAddr(maddr)
	h.Connect(ctx, *ri)
	fmt.Println("✅ Connecté au relay")

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


func unmarshal(protocolPath string) Protocol {
	data, _ := os.ReadFile(protocolPath)
	var cfg Protocol
	json.Unmarshal(data, &cfg)
	return cfg
}

func keep(cfg Protocol) string{
		base := `C:\pulse\receivers`
		os.MkdirAll(base, 0o755)
		sum:=sha256.Sum256([]byte(cfg.Protocol))
		name:=hex.EncodeToString(sum[:8])
		keep := filepath.Join(base, name+".keep")
		os.WriteFile(keep, []byte("1"), 0o600)
		return keep
	}