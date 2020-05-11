package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/webhash"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "检查网址修改",
	Long: `检查配置文件中保存的网址:

webhash check`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		for _, h := range hashs {
			if hash, err := webhash.Parse(h); err == nil {
				if !hash.Get() {
					fmt.Println(hash)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
