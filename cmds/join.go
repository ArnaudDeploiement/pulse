package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func JoinCmd() *cobra.Command{

var groupFileName string
	
	cmd := &cobra.Command{
		Use:   "Rejoindre",
		Short: "Rejoindre un groupe privé",
		Run: func(cmd *cobra.Command, args []string) {
			if groupFileName == "" {
				fmt.Println("Vous devez préciser le nom avec --name")
				os.Exit(1)
			}
			fmt.Printf("\n✅ Le groupe %s a été créé avec succès. \n\n", groupFileName)
		},
	}

	cmd.Flags().StringVarP(&groupFileName, "nom", "n", "", "Nom du groupe")


	return cmd
}

