package all_generations

import (
	"fmt"

	"github.com/gookit/color"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func GenerateExpressApi() {
	
	options := []gen.Option{
		{"Basic (/) endpoint", "Responds with hello world", true},
		{"JSON response endpoint", "Simple API with JSON data", false},
		{"MongoDB Connected API", "MongoDB connected 'users' endpoint", false},
		{"Full Featured API", "All of the above with alot more + logging", false},
	}

	expressApiSelectedOption := gen.SelectOptions(options) // run the select options thing to select the options 

	var indexJsTemplate string

	indexJsTemplate = templates.GetExpressBasicApiTemplate()

	if expressApiSelectedOption == 0 {
		indexJsTemplate = templates.GetExpressBasicApiTemplate()
	} else if expressApiSelectedOption == 1 {
		indexJsTemplate = templates.GetJsonExpressApiTemplate()
	} else if expressApiSelectedOption == 2 {
		indexJsTemplate = templates.GetMongodbExpressApiTemplate()
	} else if expressApiSelectedOption == 3 {
		indexJsTemplate = templates.GetComplexExpressApiTemplate()
	}

	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "express_api/index.js",
			TemplateString: indexJsTemplate,
		},
		{
			FileName:       "express_api/package.json",
			TemplateString: templates.ExpressPackageJson,
		},
		{
			FileName:       "express_api/Dockerfile",
			TemplateString: templates.ExpressDockerfileBasic,
		},
		{
			FileName:       "express_api/.dockerignore",
			TemplateString: templates.ExpressDockerIgnore,
		},
		{
			FileName:       "express_api/.gitignore",
			TemplateString: templates.ExpressGitIgnore,
		},
	}

	gen.ReadInput()
	fmt.Printf("Enter " + color.Green.Sprintf("Port") + " to run the Web Server on" + color.Yellow.Sprintf(" [INT]") + color.Blue.Sprintf(" ➽  "))
	webServerPort := gen.ReadInput()

	gen.GenerateFiles(webServerPort, options, fileTemplates)

	if expressApiSelectedOption == 2 {
		color.Green.Println("\n\n===========================================")
		fmt.Println("Replace the 'mongoConnection' variable in 'index.js' with your MongoDB connection string")
		fmt.Println("Also make sure to change the database and collection names based on your requirements")
		color.Green.Println("===========================================")
	}
}


