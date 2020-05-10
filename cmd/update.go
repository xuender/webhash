package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/webhash"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "修改监听网址状态",
	Long: `修改配置中的摘要，方便下次监听

webhash update`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		for i, h := range hashs {
			if hash, err := webhash.Parse(h); err == nil {
				if hash.Get() {
					fmt.Println(hash)
				} else {
					fmt.Printf("发生修改 -> %s\n", hash)
					hashs[i].Time = time.Now().Unix()
					hashs[i].Sum = hash.Sum
				}
			} else {
				fmt.Println(err)
			}
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
	rootCmd.AddCommand(updateCmd)
}
