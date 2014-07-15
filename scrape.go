package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const url = "http://www.sumo.or.jp/"

type Rikishi struct {
	Name string
	Result string
}

type Torikumi struct {
	Id int
	Higashi Rikishi
	Nishi Rikishi
	Kimarite string
}

func (t *Torikumi) ResultString() string {
	return strings.Join([]string{
		t.Higashi.Name,
		t.Higashi.Result,
		t.Kimarite,
		t.Nishi.Result,
		t.Nishi.Name,
	}," ")
}

func GetHoshitori() ([]Torikumi, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	ret := make([]Torikumi,0)
	doc.Find("div[class='cover'] table tr").Each(func(i int,tr *goquery.Selection) {
		_id := i
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
		
		torikumi := &Torikumi{}
		torikumi.Id = _id
		torikumi.Higashi.Name = v[0]
		torikumi.Higashi.Result = v[1]
		torikumi.Kimarite = v[2]
		torikumi.Nishi.Name = v[4]
		torikumi.Nishi.Result = v[3]

		ret = append(ret,*torikumi)
	})

	return ret, nil
}
