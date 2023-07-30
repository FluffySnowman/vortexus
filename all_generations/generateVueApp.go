package all_generations

import (
	"fmt"

	"github.com/fluffysnowman/vortexus/gen"
	"github.com/fluffysnowman/vortexus/templates"
)



func GenerateVueApp() {
	
	options := []gen.Option{
		{"Basic Vue Website", "Simple Vue Website", true},
	}

	gen.SelectOptions(options) // select options

	fileTemplates := []gen.FileTemplatePair{
		{
			FileName:       "vue_app/src/App.vue",
			TemplateString: templates.VueAppMainAppDotVueTemplate,
		},
		{
			FileName:       "vue_app/package.json",
			TemplateString: templates.VueAppPackageJsonTemplate,
		},
		{
			FileName:       "vue_app/.gitignore",
			TemplateString: templates.VueAppGitIgnoreTemplate,
		},
		{
			FileName:       "vue_app/public/index.html",
			TemplateString: templates.VueAppPublicHtmlTemplace,
		},
		{
			FileName: 			"vue_app/src/components/LandingPage.vue",
			TemplateString: templates.VueAppLandingPageTemplate,
		},
		{
			FileName: 			"vue_app/src/main.js",
			TemplateString: templates.VueAppMainJsTemplate,
		},
		{
			FileName: 			"vue_app/babel.config.js",
			TemplateString: templates.VueAppBabelConfigJsTemplate,
		},
		{
			FileName: 			"vue_app/jsconfig.json",
			TemplateString: templates.VueAppJsConfigJsonTemplate,
		},
		{
			FileName: 			"vue_app/vue.config.js",
			TemplateString: templates.VueAppVueConfigJs,
		},
	}

	webServerPort := "0"

	gen.ReadInput()
	fmt.Printf("\nCreating Vue App...\n")
	gen.GenerateFiles(webServerPort, options, fileTemplates)

	fmt.Printf("\n\nVue App created\n")
}


