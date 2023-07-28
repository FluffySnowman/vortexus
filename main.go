// sorry for anyone reading this mess

package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/gookit/color"

	// importing local packages from this project
	"github.com/fluffysnowman/vortexus/all_generations"
	"github.com/fluffysnowman/vortexus/gen"
)


var selectedIndex = 0

func main() {	
	mainMenu()

	color.Green.Println("\nOperation complete")
	fmt.Print("PRESS ")
	color.Yellow.Print("<ENTER>")
	fmt.Print(" to exit")
	gen.ReadInput()
}


// THE MAIN MENU FUNCTION
func mainMenu() {

options := []string{"Express.js API", "Simple Web Server (express.js)", "Python API (FastAPI)", "Basic GO API" }

	optionDescriptions := []string{
		"A basic and flexible REST API using Express.js",
		"Simple web server that can be used to host websites",
		"Simple API in python using FastAPI (this is the best option [SOURCE: trust me bro])",
		"A basic API in GO",	
	}

	fmt.Println("Use the arrow keys to select an API to create, then press enter to select")
	
	renderOptions(options, selectedIndex, optionDescriptions)

	err := keyboard.Open()
	if err != nil {
		fmt.Printf("Error opening keyboard: %v", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	for {
		
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Printf("Error reading input: %v", err)
			os.Exit(1)
		}

		// clears screen so the terminal isn't full of shit
		fmt.Print("\033[H\033[2J")

		// some fancy meth which handles the selection
		switch key {
		case keyboard.KeyCtrlC:
			fmt.Println("Exiting...")
			os.Exit(0)
		case keyboard.KeyArrowUp:
			selectedIndex = (selectedIndex - 1 + len(options)) % len(options)
		case keyboard.KeyArrowDown:
			selectedIndex = (selectedIndex + 1) % len(options)
		case keyboard.KeyEnter:

			switch selectedIndex {
			case 0:
				fmt.Printf("You have chosen: %s\n", color.Green.Sprint(options[selectedIndex]))
				// generate.GenerateExpressBasicApi()
				// generate.SelectExpressjsOptions()
				// fmt.Printf("\033[0m")
				// generate.GenerateExpressBasicApi()
				all_generations.GenerateExpressApi()
				return
			case 1:
				fmt.Printf("You have chosen: %s\n", color.Green.Sprint(options[selectedIndex]))
				all_generations.GenerateWebServerExpress()
				return
			case 2:
				fmt.Printf("You have chosen: %s (death) \n", color.Green.Sprint(options[selectedIndex]))
				all_generations.GeneratePythonApiFlask()
				return	
			case 3:	
				fmt.Printf("You have chosen: %s\n", color.Green.Sprint(options[selectedIndex]))
				fmt.Println("Press <enter> to continue")
				all_generations.GenerateGolangApi()
				return
			}

		}
		renderOptions(options, selectedIndex, optionDescriptions)
		// Check for character input to prevent arrow key artifacts on some systems (this shit shouldn't break or I'll get fucked in the ass)
		if char != 0 {
			continue
		}

		// printing ascii escape sequence for clearing the screen
		fmt.Printf("\033[0m")

	}
}

// thing to render the main menu options etc
func renderOptions(options []string, selectedIndex int, optionDescriptions []string) {
	for i, opt := range options {
		if i == selectedIndex {
			fmt.Println(color.Green.Sprint("➤ " + opt))
			fmt.Println(color.Yellow.Sprintf("  Description: " + optionDescriptions[i]) )
		} else {
			fmt.Println("  " + opt)
		}
	}
}

