package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"
	"strings"

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
		report, _ := cmd.Flags().GetBool("report")
		hashs := webhash.NewHashs(viper.Get("hashs"))
		urls := []string{}
		for _, h := range hashs {
			if hash, err := webhash.Parse(h); err == nil {
				if !hash.Get() {
					if report {
						urls = append(urls, hash.URL)
					} else {
						fmt.Println(hash)
					}
				}
			} else {
				if !report {
					fmt.Println(err)
				}
			}
		}
		if report {
			if len(urls) > 0 {
				fmt.Printf("%d个网址发生修改, [%s]\n", len(urls), strings.Join(urls, ","))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("report", "r", false, "报告输入")
}
