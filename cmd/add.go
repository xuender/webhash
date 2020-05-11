package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/webhash"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "增加监听网址",
	Long: `增加监听网址，方便以后随时检查:

webhash add https://api.github.com/repos/golang/go/milestones/72`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		for _, hash := range webhash.Batch(args) {
			if hash.Error == nil {
				hashs.Add(hash)
			}
			fmt.Println(hash)
		}
		viper.Set("hashs", hashs)
		err := viper.WriteConfig()
		if err != nil {
			if home, err := homedir.Dir(); err == nil {
				viper.WriteConfigAs(home + "/.webhash.yaml")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
