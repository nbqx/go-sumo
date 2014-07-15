package main

import (
	"os"
	"time"
	"strings"
	"unicode/utf8"

	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

const interval = 10

func draw(l [][]string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// head
	now := "LastUpdate: "+time.Now().Local().Format("2006/01/02 15:04:05")
	for x, v := range now {
		termbox.SetCell(x, 0, v, termbox.ColorDefault, termbox.ColorDefault)
	}

	// hoshitori
	for y, v := range l {
		x := 0
		str := strings.Join(v," ")
		for len(str)>0 {
			c,w := utf8.DecodeRuneInString(str)
			str = str[w:]
			termbox.SetCell(x, y+1, c, termbox.ColorDefault, termbox.ColorDefault)
			x += runewidth.RuneWidth(c)
		}
	}

	termbox.Flush()
}

func mainHandler(c *cli.Context) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	ticker := time.NewTicker(interval * time.Second)
	dat := make(chan [][]string)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <- ticker.C:
				d,err := GetHoshitori()
				if err == nil {
					dat <- d
				}
			case <- stop:
				ticker.Stop()
				break
			}
		}
	}()

	go func() {
		for r := range dat {
			draw(r)
		}
	}()

	// init
	d,err := GetHoshitori()
	if err == nil {
		dat <- d
	}

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				close(stop)
				break loop
			case termbox.KeyCtrlG:
				close(stop)
				break loop
			case termbox.KeyCtrlC:
				close(stop)
				break loop
			}
		}
	}

}

func main() {
	app := cli.NewApp()
	app.Name = "go-sumo"
	app.Action = mainHandler
	app.Run(os.Args)
}
