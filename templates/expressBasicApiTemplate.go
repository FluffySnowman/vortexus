package templates

func GetExpressBasicApiTemplate() string {
	return ExpressBasicApi
}

func GetJsonExpressApiTemplate() string {
	return JsonExpressApi
}

func GetComplexExpressApiTemplate() string {
	return ComplexExpressApi
}

func GetMongodbExpressApiTemplate() string {
	return MongodbExpressApi
}

const ExpressBasicApi = `
const express = require("express");
const app = express();

app.get("/", (req, res) => {
	res.send("hello there");
});

app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`

const JsonExpressApi = `
const express = require("express");
const app = express();

app.use(express.json());

const users = [
	{
		id: 1,
		name: "John",
	},
	{
		id: 2,
		name: "Jane",
	},
];

app.get("/", (req, res) => {
	res.send("hello there");
});


app.get("/users", (req, res) => {
	res.status(200).json(users);
});

app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`

const ComplexExpressApi = `
const express = require("express");
const app = express();
const httpLogger = require("morgan");
app.use(httpLogger("combined"));

app.use(express.json());

const users = [
	{
		id: 1,
		name: "John",
	},
	{
		id: 2,
		name: "Jane",
	},
];

app.get("/", (req, res) => {
	res.send("hello there");
});


app.get("/users", (req, res) => {
	res.status(200).json(users);
});

app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`

const MongodbExpressApi = `
const express = require("express");
const MongoClient = require("mongodb").MongoClient;
const httpLogger = require("morgan");

const app = express();

app.use(httpLogger("combined"));
app.use(express.json());

// REPLACE WITH YOUR MONGO CONNECTION STRING
const mongoConnection = ""; 

const mongodbClient = new MongoClient(mongoConnection);

async function getUserList() {
	const database = mongodbClient.db("users"); // Replace with your database name
	const collection = database.collection("users"); // Replace with your collection name
	
	const users = await collection.find({}).toArray();

	mongodbClient.close();
	return users;
}

app.get("/", (req, res) => {
	res.send("hello there");
});

app.get("/users", (req, res) => {
	const users = getUserList();
	res.status(200).json(users);
});

app.listen({{.Port}}, () => {
	console.log("App should be working at http://localhost:{{.Port}}");
});
`

const ExpressPackageJson = `
{
	"name": "express_api",
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
		"mongodb": "latest"
	}
}
`

const ExpressDockerfileBasic = `
FROM node:18

WORKDIR /express_api

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["node", "index.js"]
`

const ExpressDockerIgnore = `
node_modules
`

const ExpressGitIgnore = `
node_modules
`
