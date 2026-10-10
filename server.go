/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/microsoft/vscode-remote-try-go/hello"
)

func handle(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintln(w, hello.Hello()); err != nil {
		log.Printf("Failed to write greeting: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handle)
	log.Print("Server listening on http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
