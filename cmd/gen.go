package cmd

import (
	"fmt"
	"log"
	"math/rand"
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
		cts, _ := os.ReadFile(filename)
		now := time.Now()

		feed := &feeds.Feed{
			Title:       "docs-training",
			Link:        &feeds.Link{Href: "https://blog.wrss.top"},
			Description: "discussion about tech, footie, photos",
			Author:      &feeds.Author{Name: "hhacking", Email: "yyzw@live.com"},
			Created:     now,
		}

		feed.Items = []*feeds.Item{
			{
				Title: fmt.Sprintf("[%s] docs-training", GetToday()),
				// Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
				// Description: "A discussion on controlled parallelism in golang",
				Author:  &feeds.Author{Name: "hhacking", Email: "yyzw@live.com"},
				Content: string(cts),
				Id:      strconv.Itoa(rand.Intn(9999999)),
				Created: now,
			},
		}

		atom, err := feed.ToAtom()
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

// Feed 通用Feed
type Feed struct {
	URL, Author string
	UpdatedTime time.Time
	Title       string
}

// Item feed的item
type Item struct {
	UpdatedTime time.Time
	Enclosure   *feeds.Enclosure
	URL         string
	Title       string
	Contents    string
	ID          string
	Author      string
}

// Rss 输出rss
func Rss(fe *Feed, items []Item) string {
	if len(items) == 0 {
		feed := feeds.Feed{
			Title:   fe.Title,
			Link:    &feeds.Link{Href: fe.URL},
			Author:  &feeds.Author{Name: fe.Author},
			Updated: fe.UpdatedTime,
		}
		atom, _ := feed.ToRss()
		return atom
	}

	return rss(fe, items)
}

func rss(fe *Feed, items []Item) string {
	feed := feeds.Feed{
		Title:   fe.Title,
		Link:    &feeds.Link{Href: fe.URL},
		Author:  &feeds.Author{Name: fe.Author},
		Updated: items[0].UpdatedTime, // 直接使用最新item的时间戳
	}

	for key := range items {
		feed.Add(&feeds.Item{
			Title:       items[key].Title,
			Link:        &feeds.Link{Href: items[key].URL},
			Description: items[key].Contents,
			Author:      &feeds.Author{Name: items[key].Author},
			Id:          items[key].ID,
			Enclosure:   items[key].Enclosure,
			Updated:     items[key].UpdatedTime,
		})
	}

	// 输出atom，跟rsshub保持一致
	atom, err := feed.ToAtom()
	if err != nil {
		return ""
	}
	return atom
}

// GetToday 获取今天的零点时间
func GetToday() string {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}
