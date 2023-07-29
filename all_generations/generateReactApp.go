package all_generations

import (
	"fmt"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func GenerateReactApp() {
	
	options := []gen.Option{
		{"Development codebase", "3 page app with a navbar", true},
	}

	gen.SelectOptions(options) // run the select options thing
/* 
	var indexJsTemplate string

	indexJsTemplate = templates.GetExpressBasicApiTemplate()
*/	
	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "react_app/src/App.js",
			TemplateString: templates.ReactSrcAppTemplate,
		},
		{
			FileName:       "react_app/package.json",
			TemplateString: templates.ExpressPackageJson,
		},
		{
			FileName:       "react_app/.gitignore",
			TemplateString: templates.ReactAppGitingoreTemplate,
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

	webServerPort := "0"

	gen.ReadInput()
	fmt.Println("\nCreating React App...\n")
	gen.GenerateFiles(webServerPort, options, fileTemplates)
}


