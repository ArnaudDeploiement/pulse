package main

import (
	"pulse/cmds"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func cli() {

	const banner = `

__/\\\\\\\\\\\\\____/\\\________/\\\__/\\\_________________/\\\\\\\\\\\____/\\\\\\\\\\\\\\\_        
 _\/\\\/////////\\\_\/\\\_______\/\\\_\/\\\_______________/\\\/////////\\\_\/\\\///////////__       
  _\/\\\_______\/\\\_\/\\\_______\/\\\_\/\\\______________\//\\\______\///__\/\\\_____________      
   _\/\\\\\\\\\\\\\/__\/\\\_______\/\\\_\/\\\_______________\////\\\_________\/\\\\\\\\\\\_____     
    _\/\\\/////////____\/\\\_______\/\\\_\/\\\__________________\////\\\______\/\\\///////______    
     _\/\\\_____________\/\\\_______\/\\\_\/\\\_____________________\////\\\___\/\\\_____________   
      _\/\\\_____________\//\\\______/\\\__\/\\\______________/\\\______\//\\\__\/\\\_____________  
       _\/\\\______________\///\\\\\\\\\/___\/\\\\\\\\\\\\\\\_\///\\\\\\\\\\\/___\/\\\\\\\\\\\\\\\_ 
        _\///_________________\/////////_____\///////////////____\///////////_____\///////////////__
		
`
	root := &cobra.Command{
		Use:   "Pulse",
		Short: "Pulse - Version 0.1",
		Long:  color.HiMagentaString(banner) + "\n\n" + "Share your data with total freedom: a private, secure, and decentralized protocol — no servers, just you and your peers." + "\n\n" + color.HiCyanString("Pulse reinvents the way you share: ditch the cloud, embrace fast, direct, and secure exchange."),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}


	createCmd:= cmds.CreateCmd();
	postCmd:=cmds.PostCmd();
	getCmd:= cmds.GetCMD();
	listGCmd:=cmds.ListCmd();
	stopCmd:=cmds.StopCmd()



	root.CompletionOptions.DisableDefaultCmd = true

	//créer une commande pour connaître son PeerID
	root.AddCommand(
		createCmd,
		postCmd,
		getCmd,
		listGCmd,
		stopCmd,
	)

	root.SetHelpTemplate(`

{{with .Long}}{{.}}{{end}}

Use Pulse :

  Pulse + cmd

Cmds :
{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}
`)

	root.Execute()

}