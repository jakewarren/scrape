package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

// Exit codes
const (
	exitOK = iota
	exitScraping
	exitNoInput
	exitOpenFile
	exitFetchURL
)

// struct to hold config / app data
type config struct {
	xpath     string
	css       string
	userAgent string
	attr      string
	insecure  bool
	in        io.Reader
}

var conf config

func main() {

	helpFlag := pflag.BoolP("help", "h", false, "usage information")
	pflag.StringVarP(&conf.css, "css", "c", "", "css selector")
	pflag.StringVarP(&conf.xpath, "xpath", "x", "", "xpath query")
	pflag.StringVarP(&conf.userAgent, "agent", "A", "Mozilla/4.0 (Mozilla/4.0; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 3.0.04506.30)", "user agent string")
	pflag.StringVarP(&conf.attr, "attr", "a", "html", "attribute to scrape")
	pflag.BoolVarP(&conf.insecure, "insecure", "k", false, "skip SSL verification")
	pflag.Parse()

	if *helpFlag {
		pflag.Usage()
		os.Exit(exitOK)
	}

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Determine whether the program's input should be:
	// file, HTTP URL or stdin
	filename := pflag.Arg(0)
	if filename == "" || filename == "-" {

		// no file provided so attempt to read piped data from stdin
		info, err := os.Stdin.Stat()
		if err != nil {
			fatal(exitNoInput, err)
		}
		// ensure stdin is a pipe, bail if not
		if info.Mode()&os.ModeNamedPipe == 0 {
			fatal(exitNoInput, fmt.Errorf("%s", "No input provided to STDIN"))
		}
		conf.in = os.Stdin
	} else {
		if !validURL(filename) { // read from a file
			r, err := os.Open(filename)
			if err != nil {
				fatal(exitOpenFile, err)
			}
			conf.in = r
		} else { // read from URL
			r, err := getURL(filename)
			if err != nil {
				fatal(exitFetchURL, err)
			}
			conf.in = r
		}
	}

	if len(conf.xpath) > 0 { // XPath provided
		err := scrapeXPath()
		if err != nil {
			log.Error().Msgf("%s", err)
			os.Exit(exitScraping)
		}
	} else if len(conf.css) > 0 { // CSS selector provided
		err := scrapeCSS()
		if err != nil {
			log.Error().Msgf("%s", err)
			os.Exit(exitScraping)
		}
	}

}

func fatal(i int, e error) {
	log.Error().Msgf("%s", e)
	os.Exit(i)
}

// validURL performs extremely basic validation of a URL
func validURL(url string) bool {
	r := regexp.MustCompile("(?i)^http(?:s)?://")
	return r.MatchString(url)
}

// getURL takes in a URL and returns an io.Reader
func getURL(url string) (io.Reader, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.insecure},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   20 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", conf.userAgent)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return bufio.NewReader(resp.Body), err
}
