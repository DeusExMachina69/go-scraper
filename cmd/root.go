/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scr",
	Short: "a web scraper",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
        for _, arg := range args {
            if !hasTLD(arg) {
                arg = appendTLD(arg)
            }
            if !hasSchema(arg) {
                arg = prependSchema(arg)
            }
            res, err := Get(arg)
            if err != nil {
                panic(err)
            }
            jsOutput, err := cmd.Flags().GetBool("json")
            if err != nil {
                panic(err)
            }
            if jsOutput == true {
                res.Json()
                return
            }
            res.Print()
        }
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scr.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("json", "j", false, "Format output to JSON")
}
