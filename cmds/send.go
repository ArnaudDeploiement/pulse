package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SendCmd() *cobra.Command{

	var group string
	
	cmd := &cobra.Command{
		Use:   "envoyer",
		Short: "Envoyer des données privées",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file:=args[0]
			if file==""{
				fmt.Printf("Vous devez précier le chemin du fichier")
			}
			
			if group == ""  {
				fmt.Printf("Vous devez préciser le nom du fichier")
			}
			fmt.Printf("\n✅ %s a bien été envoyé au groupe %s.\n\n", file, group)
		},
	}

	cmd.Flags().StringVarP(&group, "groupe", "g", "", "Nom du groupe")
	cmd.MarkFlagRequired("groupe")

	return cmd

}