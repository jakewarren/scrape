# scrape

[![CircleCI](https://circleci.com/gh/jakewarren/scrape.svg?style=shield)](https://circleci.com/gh/jakewarren/scrape)
[![GitHub release](http://img.shields.io/github/release/jakewarren/scrape.svg?style=flat-square)](https://github.com/jakewarren/scrape/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/scrape/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/scrape)](https://goreportcard.com/report/github.com/jakewarren/scrape)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)


A command line scraping utility inspired by [scrape]( https://github.com/jeroenjanssens/data-science-at-the-command-line/blob/master/tools/scrape).

## Features

* Scrape using XPath or CSS selectors
* Process HTML from a URL, STDIN, or a local file
* Extract a particular attribute

## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/scrape/releases/latest](https://github.com/jakewarren/scrape/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/scrape
```


## Usage

```
Usage of scrape:
  -A, --agent string   user agent string (default "Mozilla/4.0 (Mozilla/4.0; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 3.0.04506.30)")
  -a, --attr string    attribute to scrape (default "html")
  -c, --css string     css selector
  -h, --help           usage information
  -k, --insecure       skip SSL verification
  -x, --xpath string   xpath query
```

### Examples:

#### Read from URL:
```
❯ scrape -c "h4 a" -a href "https://www.webscraper.io/test-sites/e-commerce/allinone"
/test-sites/e-commerce/allinone/product/244
/test-sites/e-commerce/allinone/product/269
/test-sites/e-commerce/allinone/product/192
```

#### Read from STDIN:
```
❯ curl -A 'Mozilla/4.0 (Mozilla/4.0; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 3.0.04506.30)' -s "https://www.webscraper.io/test-sites/e-commerce/allinone" | scrape -x "//h4/a" -a href
/test-sites/e-commerce/allinone/product/223
/test-sites/e-commerce/allinone/product/280
/test-sites/e-commerce/allinone/product/278
```

#### Read from file:
```
❯ scrape -x "//h4/a" /tmp/webscrapetest.html
<a href="/test-sites/e-commerce/allinone/product/223" class="title" title="Aspire E1-510">Aspire E1-510</a>
<a href="/test-sites/e-commerce/allinone/product/280" class="title" title="Lenovo V510 Black">Lenovo V510 Blac...</a>
<a href="/test-sites/e-commerce/allinone/product/278" class="title" title="Lenovo V510 Black">Lenovo V510 Blac...</a>
```


