package templates

func GetSimpleWebserverExpress() string {
	return SimpleWebserverExpress
}

func GetSimpleWebserverLoggingExpress() string {
	return SimpleWebserverLoggingExpress
}

const SimpleWebserverExpress = `
const express = require("express");
const app = express();
const path = require("path");

app.use(express.static("./website"));


app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`
const SimpleWebserverLoggingExpress = `
const express = require("express");
const app = express();
const path = require("path");
const httpLogger = require("morgan");

app.use(httpLogger("combined"));
app.use(express.static("./website"));

app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`

const SimpleWebserverExpressPackageJson = `
{
	"name": "website",
	"version": "1.0.0",
	"main": "index.js",
	"type": "commonjs",
	"license": "MIT",
	"scripts": {
		"dev": "node index.js",
		"prod": "pm2 start index.js"
	},
	"dependencies": {
		"express": "latest",
		"morgan": "latest",
		"nodemon": "latest",
		"pm2": "latest",
		"vhost": "latest"
	}
}
`

const SimpleWebserverExpressDockerfile = `
FROM node:18

WORKDIR /simple_webserver

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["node", "index.js"]
`

const SimpleWebserverExpressDockerIgnore = `
node_modules
`

const SimpleWebserverExpressGitignore = `
node_modules
`
