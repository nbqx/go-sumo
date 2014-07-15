package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const url = "http://www.sumo.or.jp/"

// TODO: torikumi struct

func GetHoshitori() ([][]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	ret := make([][]string,0)
	doc.Find("div[class='cover'] table tr").Each(func(i int,tr *goquery.Selection) {

		v := tr.Find("td").Map(func(j int, td *goquery.Selection) string {
			if td.Find("img").Length() == 0 {
				if str := td.Text(); str != "" {
					return strings.TrimSpace(str)
				} else {
					return "-"
				}
			} else {
				alt,exist := td.Find("img").Attr("alt")
				if exist {
					if alt == "白丸" {
						return "◯"
					} else {
						return "×"
					}
				} else {
					return ""
				}
			}
		})
		
		ret = append(ret,v)
	})

	return ret, nil
}
