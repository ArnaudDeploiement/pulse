package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func IdCmd() *cobra.Command{


	cmd:=&cobra.Command{
		Use:   "ID",
		Short: "Obtenir son ID à partager",
		Run: func(cmd *cobra.Command, args []string) {
			
			fmt.Printf("voici votre ID")
		},
	}

	
	return cmd

}