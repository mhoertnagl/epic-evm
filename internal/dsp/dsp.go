package dsp

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

type Dsp struct {
	scr   tcell.Screen
	style tcell.Style
}

func NewDsp() *Dsp {

	scr, e := tcell.NewScreen()

	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if e := scr.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	scr.SetStyle(style)

	scr.Clear()

	quit := make(chan struct{})

	dsp := &Dsp{
		scr:   scr,
		style: style,
	}

	for x := 0; x < 80; x++ {
		for y := 0; y < 25; y++ {
			dsp.SetContent(x, y, 'R')
		}
	}

	go func() {
		for {
			ev := dsp.scr.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyESC:
					close(quit)
					return
				}
			}
		}
	}()

	<-quit

	return dsp
}

func (dsp *Dsp) SetContent(x int, y int, c rune) {
	dsp.scr.SetContent(x, y, c, nil, dsp.style)
	dsp.scr.Show()
}

// func (dsp *Dsp) start() {
// 	for {
// 		ev := dsp.scr.PollEvent()
// 		switch ev := ev.(type) {
// 		case *tcell.EventKey:
// 			switch ev.Key() {
// 			case tcell.KeyESC:
// 				return
// 			}
// 		}
// 	}
// }

// func (dsp *Dsp) Continue() bool {
// 	ev := dsp.scr.PollEvent()
// 	switch ev := ev.(type) {
// 	case *tcell.EventKey:
// 		switch ev.Key() {
// 		case tcell.KeyESC:
// 			return false
// 		}
// 	}
// 	return true
// }

func (dsp *Dsp) Finalize() {
	dsp.scr.Fini()
}
