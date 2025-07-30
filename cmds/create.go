package cmds

import (
	"fmt"
	"os"

	"pulse/fn"

	"github.com/spf13/cobra"
)



func CreateCmd() *cobra.Command{

var groupName string
	
	cmd := &cobra.Command{
		Use:   "créer",
		Short: "Créer un groupe privé",
		Run: func(cmd *cobra.Command, args []string) {
			if groupName == "" {
				fmt.Println("Vous devez préciser le nom avec --name")
				os.Exit(1)
			}
		  fmt.Println(fn.FnCreate(groupName))	
		},
	}

	cmd.Flags().StringVarP(&groupName, "nom", "n", "", "Nom du groupe")


	return cmd
}

