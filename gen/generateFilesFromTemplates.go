package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/eiannone/keyboard"
	"github.com/gookit/color"
)

type FileTemplatePair struct {
	FileName       string
	TemplateString string
}

type Option struct {
	Name        string
	Description string
	Selected    bool
}

func SelectOptions(options []Option) int {
	fmt.Printf("\033[0m")

	fmt.Println("Press <space> to select options and then press <enter> to proceed")

	err := keyboard.Open()
	if err != nil {
		fmt.Printf("Error opening keyboard: %v", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	selectedIndex := 0

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Printf("Error reading input: %v", err)
			os.Exit(1)
		}

		fmt.Print("\033[H\033[2J") // clears the screen

		switch key {
		case keyboard.KeyCtrlC:
			fmt.Println("Exiting...")
			os.Exit(0)
		case keyboard.KeyArrowUp:
			selectedIndex = (selectedIndex - 1 + len(options)) % len(options)
		case keyboard.KeyArrowDown:
			selectedIndex = (selectedIndex + 1) % len(options)
		case keyboard.KeySpace:
			toggleSelection(options, selectedIndex)
		case keyboard.KeyEnter:
			getSelectedOptions(options)
			fmt.Println("You have chosen:")
			for _, opt := range getSelectedOptions(options) {
				color.Green.Print(opt + "\n")
			}
			fmt.Print("press ")
			color.BgGray.Print("<enter>")
			fmt.Print(" to continue")
			return selectedIndex
		}

		renderOptions(options, selectedIndex)

		fmt.Printf("\033[0m")

		// Check for character input to prevent arrow key artifacts on some systems
		if char != 0 {
			continue
		}
	}
}

func toggleSelection(options []Option, index int) {
	// unselect all options
	for i := range options {
		options[i].Selected = false
	}
	// toggle current selected option state 
	options[index].Selected = true
}

func getSelectedOptions(options []Option) []string {
	selectedOptions := make([]string, 0)
	for _, opt := range options {
		if opt.Selected {
			selectedOptions = append(selectedOptions, opt.Name)
		}
	}
	return selectedOptions
}

func renderOptions(options []Option, selectedIndex int) {
	output := ""

	for i, opt := range options {
		selected := opt.Selected

		// terminal output
		var optionOutput string
		if selected {
			optionOutput = color.Blue.Sprintf("(*) %s", opt.Name)
		} else {
			optionOutput = fmt.Sprintf("( ) %s", opt.Name)
		}

		if i == selectedIndex {
			optionOutput = color.Green.Sprintf("%s", optionOutput)
			optionOutput += color.Yellow.Sprintf("   %s", opt.Description)
			color.BgGray.Println("<arrow up/down> + <space> to select. double <enter> to confirm")
			fmt.Printf("\033[0m")
		}

		output += "\n" + optionOutput
	}

	// prints everything
	fmt.Println(output)
	fmt.Printf("\033[0m")
}

func GenerateFiles(apiPort string, options []Option, fileTemplates []FileTemplatePair) {
	// create data struct for the template based on the port passed to the function
	data := struct {
		Port string
	}{
		Port: apiPort,
	}

	fmt.Printf("\nGenerating files...\n")
	fmt.Println("")

	// Loop through the file-template thing to generate files
	for _, pair := range fileTemplates {
		// created dir if it doesn't exist
		dir := filepath.Dir(pair.FileName)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Println("Error creating the directory:", err)
			return
		}

		// create files
		file, err := os.Create(pair.FileName)
		if err != nil {
			fmt.Println("Error creating the file:", err)
			return
		}
		defer file.Close()

		// parse everything and create the template 
		tmpl, err := template.New(pair.FileName).Parse(pair.TemplateString)
		if err != nil {
			fmt.Println("Error parsing the template:", err)
			return
		}

		// Execute write ops
		err = tmpl.Execute(file, data)
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}

		fmt.Println("File", pair.FileName, "has been successfully created.")
	}

	fmt.Print("press ")
	color.BgGray.Print("<enter>")
	fmt.Print(" to continue")

}
