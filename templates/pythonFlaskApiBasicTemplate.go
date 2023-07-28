package templates

const PythonFlaskApi = `
from flask import Flask, request

app = Flask(__name__)

@app.route('/', methods=['GET'])
def get_items():
    return "well hello there"

if __name__ == '__main__':
    app.run(debug=True)
`

const PythonRequirementsFile = `
flask
`
