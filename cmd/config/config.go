package config

import (
	"github.com/spf13/cobra"
	"github.com/wearefair/gurl/pkg/config"
	"github.com/wearefair/gurl/pkg/log"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure gurl",
	Long: `Set configurations for gurl, which will be saved at ~/.gurl/config.
	This will prompt to set import paths and service paths. Import paths are for handling any 
	protos external to your project. Service paths are for protos internal to your project. 
	For now, the paths require absolute paths.`,
	Run: configure,
}

func configure(cmd *cobra.Command, args []string) {
	config.Read()
	config.Prompt()
	conf := config.Instance()
	err := config.Save(conf)
	if err != nil {
		log.Logger().Fatal(err.Error())
	}
}