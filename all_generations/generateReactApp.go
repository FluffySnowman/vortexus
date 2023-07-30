package all_generations

import (
	"fmt"
	"os/exec"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
	"github.com/gookit/color"
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

	fmt.Printf("\n\nReact App created\n")

	color.Yellow.Printf("\nDo you want to install the dependencies for the react app ? (y/n) ")
	installDependencies := gen.ReadInput()

	if installDependencies == "y" {
		color.Green.Printf("\nInstalling dependencies...\n\n")

		installDependenciesCommandWindows := exec.Command("cmd", "/C", "cd", "react_app", "&&", "npm", "install")
		out, err := installDependenciesCommandWindows.Output()

		if err != nil {

			installDependenciesCommandLinux := exec.Command("sh", "-c", "cd react_app && npm install")
			out, err = installDependenciesCommandLinux.Output()

			if err != nil {
				fmt.Println("Error installing dependencies: ", err)
			}
		}

		fmt.Println(string(out))
		color.Green.Printf("\nDependencies installed\n\n")
	} else {
		fmt.Printf("\nSkipping dependency installation (you can do this manually by running \"npm install\" in the react_app directory\n")
	}

/* 
	color.Yellow.Printf("\nDo you want to start the react app (y/n) ")
	startReactApp := gen.ReadInput()

	if startReactApp == "y" {
    color.Green.Printf("\nStarting react app...\n\n")

    var startReactAppCommand *exec.Cmd
    if runtime.GOOS == "windows" {
        startReactAppCommand = exec.Command("cmd", "/C", "cd", "react_app", "&&", "npm", "start")
    } else {
        startReactAppCommand = exec.Command("sh", "-c", "cd react_app && npm start")
    }

    startReactAppCommand.Stdout = os.Stdout
    startReactAppCommand.Stderr = os.Stderr

    startReactAppCommand.Stdin = os.Stdin

    // start the app async
    err := startReactAppCommand.Start()
    if err != nil {
        fmt.Println("Error starting react app: ", err)
        return
    }

    //  wait for the cmd to fininsh
    err = startReactAppCommand.Wait()
    if err != nil {
        fmt.Println("Error while running react app: ", err)
    }

    color.Green.Printf("\nReact app finished\n\n")
	} else {
		fmt.Printf("\nSkipping react app start (you can do this manually by running \"npm start\" in the react_app directory\n")
	}
 */


	color.Green.Printf("To run the react app, run \"cd react_app && npm start\"\n\n")


}


