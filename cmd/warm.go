/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var purgeFlag bool

var warmCmd = &cobra.Command{
	Use:   "warm",
	Short: "Warm the site's page cache",
	Long:  "Warm the site's page cache",
	Run: func(cmd *cobra.Command, args []string) {
		domain := CurrentDomain()

		if purgeFlag {
			PurgeCache(domain)
		}

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
				color.Println("<note>✓</> " + r.Request.URL.String())
			}
		})

		c.Visit("https://" + domain + "/wp-sitemap.xml")

		color.Info.Tips("Cache warmed for " + domain)
	},
}

func init() {
	rootCmd.AddCommand(warmCmd)

	warmCmd.Flags().BoolVar(&purgeFlag, "purge", false, "purge the cache before warming")
}
