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

		c := colly.NewCollector()

		c.OnXML("//loc", func(e *colly.XMLElement) {
			c.Visit(e.Text)
		})

		c.OnResponse(func(r *colly.Response) {
			if !strings.HasSuffix(r.Request.URL.RequestURI(), ".xml") {
				fmt.Println("Caching https://" + domain + r.Request.URL.Path)
			}
		})

		c.Visit("https://" + domain + "/wp-sitemap.xml")

		color.Info.Tips("Cache warmed for " + domain)
	},
}

func init() {
	rootCmd.AddCommand(warmCmd)
}
