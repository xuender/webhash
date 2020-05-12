package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/webhash"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "监听网页变化",
	Long: `监听网页变化，并提示

webhash watch`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		urls := []string{}
		for i, h := range hashs {
			if hash, err := webhash.Parse(h); err == nil {
				if !hash.Get() {
					urls = append(urls, hash.URL)
					hashs[i].Time = time.Now().Unix()
					hashs[i].Sum = hash.Sum
				}
			}
		}
		viper.Set("hashs", hashs)
		err := viper.WriteConfig()
		if err != nil {
			if home, err := homedir.Dir(); err == nil {
				viper.WriteConfigAs(home + "/.webhash.yaml")
			}
		}
		// TODO mac
		// Linux
		if len(urls) > 0 {
			commands := []string{
				"XDG_RUNTIME_DIR=/run/user/$(id -u)",
				"notify-send",
				"\"Webhash 提示:\"",
				"",
				"-u",
				"critical",
				"-i",
				"applications-internet",
			}
			commands[3] = fmt.Sprintf("\"%d个网页发生修改，[%s]\"", len(urls), strings.Join(urls, ", "))
			c := exec.Command("sh", "-c", strings.Join(commands, " "))
			c.Output()
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
