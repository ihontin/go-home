package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	var mainMenu = "My menu:\n1. Submenu 1\n2. Submenu 2\nPress q to quit\n"
	var firstMenu = "Submenu 1\nContent submenu 1\n"
	var secMenu = "Submenu 2\nContent submenu 2\n"
	var menuNow = mainMenu

	for {
		fmt.Print(menuNow)
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if char == 49 {
			menuNow = firstMenu
		} else if char == 50 {
			menuNow = secMenu
		} else if char == 113 {
			break
		}
		//fmt.Printf("You pressed: rune %d, key %X\r\n", char, key)
		if key == keyboard.KeyBackspace2 || key == keyboard.KeyBackspace {
			menuNow = mainMenu
		}
		fmt.Print("\033[H\033[2J")
	}
}
