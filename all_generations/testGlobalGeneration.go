package all_generations

import (
	"fmt"

	"github.com/gookit/color"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func TestGlobalGenertaion() {
	
	options := []gen.Option{
		{"localhost", "basic web server", false},
		{"logging enabled", "web server + logging", false},
	}

	gen.SelectOptions(options) // run the select options thing to select the options 

	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "simple_webserver/index.js",
			TemplateString: templates.SimpleWebserverExpress,
		},
		{
			FileName:       "simple_webserver/package.json",
			TemplateString: templates.SimpleWebserverExpressPackageJson,
		},
		{
			FileName:       "simple_webserver/Dockerfile",
			TemplateString: templates.SimpleWebserverExpressDockerfile,
		},
		{
			FileName:       "simple_webserver/.dockerignore",
			TemplateString: templates.SimpleWebserverExpressDockerIgnore,
		},
		{
			FileName:       "simple_webserver/.gitignore",
			TemplateString: templates.SimpleWebserverExpressGitignore,
		},
	}

	gen.ReadInput()
	fmt.Printf("Enter " + color.Green.Sprintf("Port") + " to run the Express API on" + color.Yellow.Sprintf(" [INT]") + color.Blue.Sprintf(" ➽  "))
	expressApiPort := gen.ReadInput()

	gen.GenerateFiles(expressApiPort, options, fileTemplates)
}

