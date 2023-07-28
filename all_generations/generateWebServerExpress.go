package all_generations

import (
	"fmt"

	"github.com/gookit/color"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func GenerateWebServerExpress() {
	
	options := []gen.Option{
		{"Basic Web Server", "web server + logging", true},
	}

	gen.SelectOptions(options) // run the select options thing to select the options 

	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "simple_webserver/index.js",
			TemplateString: templates.SimpleWebserverLoggingExpress,
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
	fmt.Printf("Enter " + color.Green.Sprintf("Port") + " to run the Web Server on" + color.Yellow.Sprintf(" [INT]") + color.Blue.Sprintf(" ➽  "))
	webServerPort := gen.ReadInput()

	gen.GenerateFiles(webServerPort, options, fileTemplates)
}


