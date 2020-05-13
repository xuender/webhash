package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/webhash"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "webhash",
	Short:   "网页摘要",
	Long:    `生成网页摘要`,
	Version: "1.1.3",
	Run: func(cmd *cobra.Command, args []string) {
		hashs := webhash.NewHashs(viper.Get("hashs"))
		if len(hashs) == 0 {
			cmd.Help()
			return
		}
		for _, c := range cmd.Commands() {
			if c.Use == "check" {
				c.Run(cmd, args)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .webhash.yaml)")

	// rootCmd.Flags().BoolP("report", "r", false, "报告输出")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".webhash")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("读取配置文件:", viper.ConfigFileUsed())
	// }
}
