package term

import "fmt"
import termbox "github.com/nsf/termbox-go"

func MainTerm() {
	var menu = [3]string{"1.Start", "2.Option", "3.Quit"}
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	var pointer = 1
	// termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	// termbox.SetOutputMode(termbox.Output256)
	fmt.Println("Welcome to termFarm 🏘")
	PrintMenu(1, menu)
	Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				fmt.Println("👋 Bye~ ")
				break Loop
			case termbox.KeyArrowUp:
				if pointer > 1 {
					pointer--
				}
				PrintMenu(pointer, menu)
			case termbox.KeyArrowDown:
				if pointer < len(menu) {
					pointer++
				}
				PrintMenu(pointer, menu)
			// case termbox.KeyArrowLeft:
			// 	fmt.Println(" ")
			// case termbox.KeyArrowRight:
			// 	fmt.Println(" ")
			case termbox.KeyEnter:
				fmt.Println("👋 Enter~ ")
				break Loop
			default:
			}
		}
	}
}

func PrintMenu(num int, menu[3]string) {
	var unpoint = "   "
	var point = " > "
	var i int
	for i = 0; i < len(menu); i++ {
		if num == i+1 {
			fmt.Println(point + menu[i])
		} else {
			fmt.Println(unpoint + menu[i])
		}
	}
}

// func Clean(num int) {
// 	var i int
// 	for i =0; i<num;i++ {
// 		fmt.Println("\r")
// 	}
// }