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
			TemplateString: templates.ReactAppPackageJsonTemplate,
		},
		{
			FileName:       "react_app/.gitignore",
			TemplateString: templates.ReactAppGitingoreTemplate,
		},
		{
			FileName:       "react_app/public/index.html",
			TemplateString: templates.ReactAppPublicIndexHtmlTemplate,
		},
		{
			FileName: 			"react_app/src/index.js",
			TemplateString: templates.ReactAppIndexJsTemplate,
		},
		{
			FileName: 			"react_app/src/App.css",
			TemplateString: templates.ReactSrcAppCssTemplate,
		},
	}

	webServerPort := "0"

	gen.ReadInput()
	fmt.Printf("\nCreating React App...\n")
	gen.GenerateFiles(webServerPort, options, fileTemplates)

}


