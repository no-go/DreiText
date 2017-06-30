package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/andlabs/ui"
)

func main() {
	arg := os.Args[1]
	filebytes, err := ioutil.ReadFile(arg)
	if err != nil {
		log.Fatal(err)
	}

	blockWidth := 40
	blockHeight := 1400
	relPos := 0

	s := string(filebytes)
	content := ""
	for i := 0; i < len(s); i++ {
		if i%blockWidth == 0 {
			content += "\n"
		}
		if string(s[i]) == "\n" {
			content += strings.Repeat(" ", blockWidth)
			content += "\n"
		} else {
			content += string(s[i])
		}
	}

	err = ui.Main(func() {
		t1 := ui.NewLabel("")
		t2 := ui.NewLabel("")
		t3 := ui.NewLabel("")

		if len(content) <= blockHeight+relPos {
			t1.SetText(content[relPos:])
		} else {
			t1.SetText(content[relPos : blockHeight+relPos])
			if len(content) <= 2*blockHeight+relPos {
				t2.SetText(content[blockHeight+relPos : len(content)])
			} else {
				t2.SetText(content[blockHeight+relPos : 2*blockHeight+relPos])
				if len(content) <= 3*blockHeight+relPos {
					t3.SetText(content[2*blockHeight+relPos : len(content)])
				} else {
					t3.SetText(content[2*blockHeight+relPos : 3*blockHeight+relPos])
				}
			}
		}

		sl := ui.NewSlider(0, 1000)

		sl.OnChanged(func(s *ui.Slider) {
			relPos = len(content) * s.Value() / 1000.0

			if len(content) <= blockHeight+relPos {
				t1.SetText(content[relPos:])
			} else {
				t1.SetText(content[relPos : blockHeight+relPos])
				if len(content) <= 2*blockHeight+relPos {
					t2.SetText(content[blockHeight+relPos : len(content)])
				} else {
					t2.SetText(content[blockHeight+relPos : 2*blockHeight+relPos])
					if len(content) <= 3*blockHeight+relPos {
						t3.SetText(content[2*blockHeight+relPos : len(content)])
					} else {
						t3.SetText(content[2*blockHeight+relPos : 3*blockHeight+relPos])
					}
				}
			}
		})

		box := ui.NewHorizontalBox()
		box.Append(t1, true)
		box.Append(t2, true)
		box.Append(t3, true)

		box2 := ui.NewVerticalBox()
		box2.Append(box, true)
		box2.Append(sl, false)

		window := ui.NewWindow("DreiText: "+arg, 800, 600, false)
		window.SetChild(box2)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}
