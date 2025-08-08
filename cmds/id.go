package cmds

import (
	"fmt"
	"pulse/fn"

	"github.com/spf13/cobra"
)

func IdCmd() *cobra.Command{


	cmd:=&cobra.Command{
		Use:   "ID",
		Short: "Obtenir son ID à partager",
		Run: func(cmd *cobra.Command, args []string) {
			
		 id :=	fn.GetIdPerson()

		 fmt.Println(id)
		},
	}

	
	return cmd

}