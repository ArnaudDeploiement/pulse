package cmds

import (
	"fmt"
	"os"
	"pulse/fn"

	"github.com/spf13/cobra"
)

func PostCmd() *cobra.Command{

	var file string
	var protocol string
	
	cmd := &cobra.Command{
		Use:   "post",
		Short: "Publish data to a protocol",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		if protocol=="" || file == ""   {
				fmt.Printf("You have to specify Protocol path & File path")
				os.Exit(1)
			}
			
			fn.FnPost(protocol,file)
			
			
		},
	}

	cmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Protocol Path")
	cmd.MarkFlagRequired("protocol")
	cmd.Flags().StringVarP(&file, "file", "f", "", "File Path")
	cmd.MarkFlagRequired("file")

	return cmd

}