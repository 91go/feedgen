/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "feedgen",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.feedgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringP("filename", "f", "", "HTML file ready to display")
	err := rootCmd.MarkPersistentFlagRequired("filename")
	if err != nil {
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringP("title", "n", "interview-questions", "custom feed title")
	rootCmd.PersistentFlags().StringP("description", "d", "feedgen is a cli tool that help you quickly build RSS feed", "custom feed description")
	rootCmd.PersistentFlags().StringP("link", "l", "https://github.com/dashboard", "custom feed link")
	rootCmd.PersistentFlags().StringP("author", "a", "nobody", "custom feed author")
	rootCmd.PersistentFlags().StringP("mail", "m", "nobody@gmail.com", "custom author's mail")
}
