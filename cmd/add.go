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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "增加监听网址",
	Long: `增加监听网址，方便以后随时检查:

webhash add https://pinyin.sogou.com/linux/changelog.php`,
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs").([]interface{}))
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
