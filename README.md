# Try Out Development Containers: Go

This is a sample project that lets you try out the **[VS Code Remote - Containers](https://aka.ms/vscode-remote/containers)** extension in a few easy steps.

> **Note:** If you're following the quick start, you can jump to the [Things to try](#things-to-try) section. 

## Setting up the development container

Follow these steps to open this sample in a container:

1. If this is your first time using a development container, please follow the [getting started steps](https://aka.ms/vscode-remote/containers/getting-started).

2. To use this repository, you can either open the repository in an isolated Docker volume:

    - Press <kbd>F1</kbd> and select the **Remote-Containers: Try a Sample...** command.
    - Choose the "Go" sample, wait for the container to start and try things out!
        > **Note:** Under the hood, this will use **Remote-Containers: Open Repository in Container...** command to clone the source code in a Docker volume instead of the local filesystem.

   Or open a locally cloned copy of the code:

   - Clone this repository to your local filesystem.
   - Press <kbd>F1</kbd> and select the **Remote-Containers: Open Folder in Container...** command.
   - Select the cloned copy of this folder, wait for the container to start, and try things out!
   
## Things to try

Once you have this sample opened in a container, you'll be able to work with it like you would locally.

> **Note:** This container runs as a non-root user with sudo access by default. Comment out `"remoteUser": "vscode"` in `.devcontainer/devcontainer.json` if you'd prefer to run as root.

Some things to try:

1. **Edit:**
   - Open `server.go`
   - Try adding some code and check out the language features.
2. **Terminal:** Press <kbd>ctrl</kbd>+<kbd>shift</kbd>+<kbd>\`</kbd> and type `uname` and other Linux commands from the terminal window.
3. **Build, Run, and Debug:**
   - Open `server.go`
   - Add a breakpoint (e.g. on line 22).
   - Press <kbd>F5</kbd> to launch the app in the container.
   - Once the breakpoint is hit, try hovering over variables, examining locals, and more.
   - Continue, then open a local browser and go to `http://localhost:9000` and note you can connect to the server in the container.
4. **Forward another port:**
   - Stop debugging and remove the breakpoint.
   - Open `server.go`
   - Change the server port to 5000. (`portNumber := "5000"`)
   - Press <kbd>F5</kbd> to launch the app in the container.
   - Press <kbd>F1</kbd> and run the **Forward a Port** command.
   - Select port 5000.
   - Click "Open Browser" in the notification that appears to access the web app on this new port.
5. **Refactoring - rename:**
    - Open `hello.go`, select method name `Hello` press <kbd>F1</kbd> and run the **Rename Symbol** command.
6. **Refactoring - extract:**
   - Open `hello.go` and select string, press <kbd>F1</kbd> and run the **Go: Extract to variable** command.
   - Open `hello.go` and select line with return statement, press <kbd>F1</kbd> and run the **Go: Extract to function** command.
7. **Generate tests:**
    - Open `hello.go` and press <kbd>F1</kbd> and run the **Go: Generate Unit Tests For File** command.
    - Implement a test case: Open file `hello_test.go` and edit the line with the `TODO` comment: `{"hello without name", "Hello, "},` 
    - You can toggle between implementation file and test file with press <kbd>F1</kbd> and run the **Go: Toggle Test File**
    - Tests can also run as benchmarks: Open file `hello_test.go`, press <kbd>F1</kbd> and run the **Go: Benchmark File**
8. **Stub generation:** ( [details](https://github.com/josharian/impl))
   - define a struct `type mock struct {}`, enter a new line , press <kbd>F1</kbd> and run the **Go: Generate interface stubs** command.
   - edit command `m *mock http.ResponseWriter`
9. **Fill structs:** ([details](https://github.com/davidrjenni/reftools/tree/master/cmd/fillstruct))
   - Open `hello.go` and select `User{}` of variable asignment, press <kbd>F1</kbd> and run the **Go: Fill struct** command.
10. **Add json tags to structs:** ([details](https://github.com/fatih/gomodifytags))
   - Open `hello.go` and go with cursor in to a struct, press <kbd>F1</kbd> and run the **Go: Add Tags To Struct Fields** command.



## Contributing

This project welcomes contributions and suggestions. Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., label, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## License

Copyright © Microsoft Corporation All rights reserved.<br />
Licensed under the MIT License. See LICENSE in the project root for license information.
