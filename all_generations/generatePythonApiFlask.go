package all_generations

import (
	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func GeneratePythonApiFlask() {
	
	options := []gen.Option{
		{},
	}
/* 
	gen.SelectOptions(options) // run the select options thing to select the options 
 */
	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "python_api/app.py",
			TemplateString: templates.PythonFlaskApi,
		},
		{
			FileName: 			"python_api/requirements.txt",
			TemplateString: templates.PythonRequirementsFile,
		},
	}

	noport := "0"

	gen.GenerateFiles(noport, options, fileTemplates)
}


