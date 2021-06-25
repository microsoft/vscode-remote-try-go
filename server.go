/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()["name"]
	name := "<default>"
	if names != nil {
		name = names[0]
	}
	io.WriteString(w, hello.Hello(name))
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
