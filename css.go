package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func scrapeCSS() error {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	doc, err := goquery.NewDocumentFromReader(conf.in)
	if err != nil {
		return err
	}

	doc.Find(conf.css).Each(func(_ int, s *goquery.Selection) {

		switch conf.attr {
		case "html":
			html, htmlErr := goquery.OuterHtml(s)
			if htmlErr == nil {
				fmt.Println(html)
			}
		case "text":
			fmt.Println(s.Text())
		default:
			attr, exists := s.Attr(conf.attr)
			if exists {
				fmt.Println(attr)
			} else {
				log.Fatal().Err(fmt.Errorf("%s", "requested attribute does not exist")).Msg("error scraping the page")
			}

		}

	})

	return nil
}
