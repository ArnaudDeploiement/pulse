package cmds

import (
	"fmt"
	"pulse/fn"

	"github.com/spf13/cobra"
)

func ReceiveCMD() *cobra.Command{
	var repoName string;
	var protocol string;

		cmd := &cobra.Command{
		Use:   "recevoir",
		Short: "Recevoir des données",
		Run: func(cmd *cobra.Command, args []string) {
			if repoName == "" || protocol == "" {
				fmt.Printf("Vous devez préciser un chemin de dépot --d\nVous devez préciser le chemin du protocol.json --p")
			} 
			fn.FnReceive(protocol,repoName)	
		},
	}



	cmd.Flags().StringVarP(&repoName, "dépot", "d", "", "Chemin du dépot")
	cmd.Flags().StringVarP(&protocol, "protocole", "p", "", "Chemin du protocol")
	cmd.MarkFlagRequired("dépot")
	cmd.MarkFlagRequired("protocol")

	return cmd

}