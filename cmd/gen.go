package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/feeds"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("filename")
		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		author, _ := cmd.Flags().GetString("author")
		mail, _ := cmd.Flags().GetString("mail")
		link, _ := cmd.Flags().GetString("link")
		cts, _ := os.ReadFile(filename)
		now := time.Now()

		feed := &feeds.Feed{
			Title:       title,
			Description: description,
			Author:      &feeds.Author{Name: author, Email: mail},
			Link:        &feeds.Link{Href: link},
			Created:     now,
		}

		feed.Items = []*feeds.Item{
			{
				Title:   fmt.Sprintf("[%s] %s", GetToday(), title),
				Author:  &feeds.Author{Name: author, Email: mail},
				Content: string(cts),
				Id:      strconv.Itoa(int(time.Now().Unix())),
				Created: now,
			},
		}

		atom, err := feed.ToRss()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(atom)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GetToday 获取今天的零点时间
func GetToday() string {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}
