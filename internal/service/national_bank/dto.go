package national_bank

import "encoding/xml"

type (
	Rss struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Channel Channel  `xml:"channel"`
	}

	Channel struct {
		Generator   string `xml:"generator"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Copyright   string `xml:"copyright"`
		Item        []Item `xml:"item"`
	}

	Item struct {
		Title       string `xml:"title"`
		PubDate     string `xml:"pubDate"`
		Description string `xml:"description"`
		Quant       string `xml:"quant"`
		Index       string `xml:"index"`
		Change      string `xml:"change"`
		Link        string `xml:"link"`
	}
)
