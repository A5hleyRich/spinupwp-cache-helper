/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// warmCmd represents the warm command
var warmCmd = &cobra.Command{
	Use:   "warm",
	Short: "Warm a site's page cache",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()

		if err != nil {
			color.Error.Println(err.Error())
			os.Exit(1)
		}

		parts := strings.Split(dir, string(filepath.Separator))

		if !strings.HasPrefix(dir, "/sites") || len(parts) < 2 {
			color.Warn.Tips("This does not seem to be a SpinupWP site")
			os.Exit(1)
		}

		domain := parts[2]

		c := colly.NewCollector(
			colly.AllowedDomains(domain),
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"),
		)

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			r.Headers.Add("accept-encoding", "gzip, deflate, br")
		})

		c.OnXML("//loc", func(e *colly.XMLElement) {
			c.Visit(e.Text)
		})

		c.OnResponse(func(r *colly.Response) {
			if !strings.HasSuffix(r.Request.URL.RequestURI(), ".xml") {
				fmt.Println("Caching " + r.Request.URL.String())
			}
		})

		c.Visit("https://" + domain + "/wp-sitemap.xml")

		color.Info.Tips("Cache warmed for " + domain)
	},
}

func init() {
	rootCmd.AddCommand(warmCmd)
}
