package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func PostCmd() *cobra.Command{

	var group string
	
	cmd := &cobra.Command{
		Use:   "post",
		Short: "Publish data to a protocol",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file:=args[0]
			if file==""{
				fmt.Printf("You have to specify file path")
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