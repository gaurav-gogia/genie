<p align="center">
  <img src="./genie-header.png">
</p>

# Genie
A smol &amp; simple Go web app project boiler plate generator that takes care of registering CRUD handlers, CORS, safe headers, logging, multi-threading, mongodb connection.

## How to Use
1. git clone <any_dir>
2. cd <any_dir/genie>
3. go install
4. Navigate to your any dir
5. Make sure that **"$HOME/go/bin"** folder has been added to environment variable **Path**
6. genie <project_name>
7. Enjoy!

### Caveats
Make sure you have MongoDB Server installed and running. Generated code makes connections with MongoDB by default. It'll panic if it cannot connect to a running instance. You may switch to a different database post code generation if it doesn't fit your requirements.

## Mission
- **Minimalistic**: Minimum, yet useful amount of boiler plate code generation.
- **Productivity**: Let developers spend more time writing the code that matters &amp; less time on project setup.
- **Speed**: Code generation should be as fast as possible.

## Features:
1. Generates minimal boiler plate code to get you started.
2. Registers handlers for CRUD operations.
3. Creates global connection with mongodb that can be extended by cloning the global session.
4. Registers CORS on the api.
5. Adds "safe" headers in the application like content time, x-frame-options etc. to try and avoid security issues.
6. Adds gorilla mux for all handlers.
7. Blazing Fast! Works almost instantly!

### Contribution Guidelines :
1. Fork the repo and create your branch from master.
2. Add the changes/fixes! 
3. Make sure your code lints.
4. Issue that pull request (against `development` branch)!

Please follow contribution guidelines to make things easy for everyone. Thanks ^.^
