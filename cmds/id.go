package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func IdCmd() *cobra.Command {

	var flag string

	cmd := &cobra.Command{
		Use:   "ID",
		Short: "Get PeerID",
		Run: func(cmd *cobra.Command, args []string) {
			
		switch flag {
		case "ID" :
			fmt.Println("get id")
		case "Add"	:
		 fmt.Println( "add id")
		default:
			fmt.Println("erreur")
		}
	
		cmd.Flags().StringVarP(&flag, "mode", "m", "get", "Action: get or add")


		},
	}


	return cmd
}
