<p align="center">
  <img src="https://github.com/De-ma/genie/blob/master/genie-header.png">
</p>

# Genie
A smol &amp; simple Go web app project boiler plate generator that takes care of registering CRUD handlers, CORS, safe headers, logging, multi-threading, mongodb connection.

## How to Use
1. go get -u -v github.com/DesmondANIMUS/genie
2. Navigate to your WORKDIR, defaults: <br/>
  a. Windows: `%userprofile%/go/src/github.com/<github_username>/` <br/>
  b. Linux & MacOS: `$HOME/go/src/github.com/<github_username>/` <br/>
3. Make sure that **"go/bin"** folder has been added to environment variable **Path**
4. genie <project_name>
5. Enjoy!

## Mission
1. Minimalistic: Minimum, yet useful amount of boiler plate code generation.
2. Productivity: Let developers spend more time writing the code that matters &amp; less time on project setup.
3. Speed: Code generation should be as fast as possible.

## Features:
1. Generates minimal boiler plate code to get you started.
2. Registers handlers for CRUD operations.
3. Creates global connection with mongodb that can be extended by cloning the global session.
4. Registers CORS on the api.
5. Adds "safe" headers in the application like content time, x-frame-options etc. to try and avoid security issues.
6. Adds gorilla mux for all handlers.
7. Blazing Fast! Works almost instantly!

### Contribution Guidelines (subject to change):
1. Star & Fork the repo.
2. Propose an issue.
3. Wait until that issue is assigned to you.
4. Make pull request on **development** branch!
5. Enjoy!

Take a look at project's wiki to understand the different parts of this little project.

Please follow contribution guidelines to make things easy for everyone. Thanks ^.^
