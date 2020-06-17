package cmd

/*
Copyright © 2020 妙音 <xuender@139.com>
*/

import (
	"fmt"
	"os/exec"
	"runtime"
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
				if keep, err2 := hash.Get(); !keep && err2 == nil {
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
		if len(urls) > 0 {
			title := "'Webhash 提示:'"
			message := fmt.Sprintf("'%d个网页发生修改，[%s]'", len(urls), strings.Join(urls, ", "))
			switch runtime.GOOS {
			case "linux":
				if _, err := exec.LookPath("notify-send"); err != nil {
					return
				}
				comm := fmt.Sprintf(`XDG_RUNTIME_DIR=/run/user/$(id -u) notify-send -u critical -i applications-internet %s %s`, title, message)
				// fmt.Println(comm)
				exec.Command("sh", "-c", comm).Run()
				// fmt.Println(err)
			case "darwin":
				if _, err := exec.LookPath("terminal-notifier"); err != nil {
					return
				}
				args := []string{
					"-e",
					"display",
					"notification",
					message,
					"with",
					"title",
					title,
				}
				exec.Command("osascript", args...).Run()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
