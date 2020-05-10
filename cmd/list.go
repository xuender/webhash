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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "显示网页摘要",
	Long: `显示记录的网页摘要

 webhash list`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		for _, h := range hashs {
			fmt.Println(h)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
