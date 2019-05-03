/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello remote world!")
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", hello)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
