package all_generations

import (
	"fmt"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
	"github.com/gookit/color"
)



func GenerateGolangApi() {
	
	options := []gen.Option{
		{},
	}
/* 
	gen.SelectOptions(options) // run the select options thing to select the options 
 */

	mainApiFile := templates.GetGolangApiTemplate()

	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "golang_api/main.go",
			TemplateString: mainApiFile,
		},
	}


	gen.ReadInput()
	fmt.Printf("Enter " + color.Green.Sprintf("Port") + " to run the Web Server on" + color.Yellow.Sprintf(" [INT]") + color.Blue.Sprintf(" ➽  "))
	goApiPort := gen.ReadInput()

	gen.GenerateFiles(goApiPort, options, fileTemplates)
}



