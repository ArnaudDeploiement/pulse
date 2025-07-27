package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func StopCmd() *cobra.Command{

	var groupFileName string;

	
	cmd := &cobra.Command{
		Use:   "Stop",
		Short: "Stop la réception",
		Run: func(cmd *cobra.Command, args []string) {
			if groupFileName == "" {
				fmt.Printf("Vous devez préciser le nom du groupe pour l'arrêter --name")
			}
			fmt.Printf("\n✅ Votre dépot a bien été enregistré à ce chemin : %s.\n\n", groupFileName)
		},
	}

	cmd.Flags().StringVarP(&groupFileName, "name", "n", "", "Nom du groupe")

	return cmd

}