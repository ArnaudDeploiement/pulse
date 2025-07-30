package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ReceiveCMD() *cobra.Command{

	var repoName string;

		cmd := &cobra.Command{
		Use:   "recevoir",
		Short: "Recevoir des données",
		Run: func(cmd *cobra.Command, args []string) {
			if repoName == "" {
				fmt.Printf("Vous devez préciser un chemin de dépot --dépot")
			}
			fmt.Printf("\n✅ Votre dépot a bien été enregistré à ce chemin : %s.\n\n", repoName)
		},
	}

	cmd.Flags().StringVarP(&repoName, "dépot", "d", "", "Chemin du dépot")
	cmd.MarkFlagRequired("dépot")

	return cmd

}