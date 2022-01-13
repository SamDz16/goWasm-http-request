package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"syscall/js"
)

// Copyright (C) 2020 Alessandro Segala (ItalyPaleAle)
// License: MIT

// MyGoFunc fetches an external resource by making a HTTP request from Go
// The JavaScript method accepts one argument, which is the URL to request
func MyGoFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get the URL as argument
		// args[0] is a js.Value, so we need to get a string out of it
		requestUrl := args[0].String()
		sparqlQuery := args[1].String()

		// Handler for the Promise
		// We need to return a Promise because HTTP requests are blocking in Go
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			// Run this code asynchronously
			go func() {
				// Make the HTTP request
				// res, err := http.DefaultClient.Get(requestUrl)
				data := url.Values{}
				data.Set("query", sparqlQuery)

				u, _ := url.ParseRequestURI(requestUrl)
				urlStr := u.String()

				client := &http.Client{}
				r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				res, err := client.Do(r)
				if err != nil {
					// Handle errors: reject the Promise if we have an error
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}
				defer res.Body.Close()

				// Read the response body
				dataBody, err := ioutil.ReadAll(res.Body)
				if err != nil {
					// Handle errors here too
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
				arrayConstructor := js.Global().Get("Uint8Array")
				dataJS := arrayConstructor.New(len(dataBody))
				js.CopyBytesToJS(dataJS, dataBody)

				// Create a Response object and pass the data
				responseConstructor := js.Global().Get("Response")
				response := responseConstructor.New(dataJS)

				// Resolve the Promise
				resolve.Invoke(response)
			}()

			// The handler of a Promise doesn't return any value
			return nil
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {
	c := make(chan int)

	js.Global().Set("MyGoFunc", MyGoFunc())

	<-c
}
