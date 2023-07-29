package templates

func GetReactAppSrcTemplate() string {
	return ReactSrcAppTemplate
}

const ReactSrcAppTemplate = `

import React from "react";
import { BrowserRouter as Router, Route, Link, Routes } from "react-router-dom";
import "./App.css";

function Home() {
	return (
		<div>
			<h2>Home Page</h2>
			<h3>This is the home page for this website</h3>
			<h1>
				feel free to edit this website template in the <code>./src/App.js</code> file in the root directory of this project
			</h1>
		</div>
	);
}

function About() {
	return (
		<div>
			<h2>About Page</h2>
			<h3>This is the about page</h3>
		</div>
	);
}

function Contact() {
	return (
		<div>
			<h2>Contact Page</h2>
			<h3>This is the contact page</h3>
		</div>
	);
}

function App() {
	return (
		<Router>
			<div className="app">
				<nav>
					<ul className="nav-links">
						<li>
							<Link to="/">Home</Link>
						</li>
						<li>
							<Link to="/about">About</Link>
						</li>
						<li>
							<Link to="/contact">Contact</Link>
						</li>
					</ul>
				</nav>

				<hr />

				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/about" element={<About />} />
					<Route path="/contact" element={<Contact />} />
				</Routes>
			</div>
		</Router>
	);
}

export default App;

`

const ReactSrcAppCssTemplate = `

:root {
	background-color: black;
	color: #20c20e;
	font-family: "Courier New", Courier, monospace;
}

.app {
	max-width: 600px;
	margin: 0 auto;
	padding: 20px;
}

.nav-links {
	list-style-type: none;
	padding: 0;
	margin: 0;
	display: flex;
	justify-content: center;
}

.nav-links li {
	margin: 0 10px;
	padding: 5px;
}

.nav-links li a {
	text-decoration: none;
	color: #ffffff;
	font-weight: bold;
}

.nav-links li a:hover {
	color: #007bff;
	background-color: #ffffff4f;
	padding: 10px;
}

hr {
	margin: 20px 0;
	border: 0;
	border-top: 1px solid #ccc;
}



`

const ReactAppPackageJsonTemplate = `

{
  "name": "react_app",
  "version": "0.1.0",
  "private": true,
  "dependencies": {
    "@testing-library/jest-dom": "^5.17.0",
    "@testing-library/react": "^13.4.0",
    "@testing-library/user-event": "^13.5.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-scripts": "5.0.1",
    "web-vitals": "^2.1.4",
    "react-router-dom": "latest"
  },
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test",
    "eject": "react-scripts eject"
  },
  "eslintConfig": {
    "extends": [
      "react-app",
      "react-app/jest"
    ]
  },
  "browserslist": {
    "production": [
      ">0.2%",
      "not dead",
      "not op_mini all"
    ],
    "development": [
      "last 1 chrome version",
      "last 1 firefox version",
      "last 1 safari version"
    ]
  }
}
`

const ReactAppGitingoreTemplate = `
/node_modules
node_modules
node_modules/
/.pnp
.pnp.js

# testing
/coverage

# production
/build

# misc
.DS_Store
.env.local
.env.development.local
.env.test.local
.env.production.local

npm-debug.log*
yarn-debug.log*
yarn-error.log*

`

const ReactAppPublicIndexHtmlTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <link rel="icon" href="%PUBLIC_URL%/favicon.ico" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    <title>React App</title>
  </head>
  <body>
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root"></div>

  </body>
</html>
`

const ReactAppIndexJsTemplate = `

import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

`
