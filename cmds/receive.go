package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ReceiveCMD() *cobra.Command{

	var repoName string;
	var invite string;

		cmd := &cobra.Command{
		Use:   "recevoir",
		Short: "Recevoir des données",
		Run: func(cmd *cobra.Command, args []string) {
			if repoName == "" {
				fmt.Printf("Vous devez préciser un chemin de dépot --d")
			}
			fmt.Printf("\n✅ Votre dépot a bien été enregistré à ce chemin : %s.", repoName)

			if invite == "" {
				fmt.Printf("Vous devez préciser le chemin de l'invite.json --i")
			}
			fmt.Printf("\n✅ Votre invite a bien été enregistrée %s\n\n", invite)
		},
	}



	cmd.Flags().StringVarP(&repoName, "dépot", "d", "", "Chemin du dépot")
	cmd.Flags().StringVarP(&invite, "invite", "i", "", "Chemin de l'invite")
	cmd.MarkFlagRequired("dépot")
	cmd.MarkFlagRequired("invite")

	return cmd

}