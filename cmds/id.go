package cmds

import (
	"github.com/spf13/cobra"
)

func IdCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "ID",
		Short: "Get PeerID",
		Run: func(cmd *cobra.Command, args []string) {
			
			//Obtenir son id ou créer un ID File des autres membres
		},
	}


	return cmd
}
