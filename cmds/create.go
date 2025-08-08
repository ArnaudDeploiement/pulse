package cmds

import (
	"fmt"
	"os"

	"pulse/fn"

	"github.com/spf13/cobra"
)



func CreateCmd() *cobra.Command{

var groupName string
var relayAddr string

	
	cmd := &cobra.Command{
		Use:   "créer",
		Short: "Créer un groupe privé",
		Run: func(cmd *cobra.Command, args []string) {
			if groupName == "" {
				fmt.Println("Vous devez préciser le nom avec -n")
				os.Exit(1)
			}
			if relayAddr == "" {
				fmt.Println("Vous devez préciser le relai avec -r")
				os.Exit(1)
			}
		
		  fmt.Println(fn.FnCreate(groupName, relayAddr))	
		},
	}

	cmd.Flags().StringVarP(&groupName, "nom", "n", "", "Nom du groupe")
	cmd.Flags().StringVarP(&relayAddr, "relai", "r", "", "Adresse du relai")


	return cmd
}

