package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func scrapeXPath() error {
	html, err := html.Parse(conf.in)

	if err != nil {
		return err
	}

	htmlquery.FindEach(html, conf.xpath, processNode)

	return nil
}

func processNode(_ int, node *html.Node) {
	switch conf.attr {
	case "html":
		fmt.Println(htmlquery.OutputHTML(node, true))
	case "text":
		fmt.Println(htmlquery.InnerText(node))
	default:
		for _, n := range node.Attr {
			if n.Key == conf.attr {
				fmt.Println(n.Val)
			}
		}
	}

}
