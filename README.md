# TinyGo & WebAssembly with syscall/js
This little applicaton demonstrates how we can integrate Golang programming language to work along with WebAssembly. Two example were made, the first one adds two numbers from user input, the second one capitalize a user input string. 

## Installation
In order to execute this app, you have to have some prerequisites:
* Go and install the GO programming language at https://go.dev
* Install TinyGO: follow this [installation guide](https://tinygo.org/getting-started/install/) to get tinygo installed into your machine
* Make sure to add the `C:\Program Files\Go\bin` and the `C:\tinygo\bin` to your PATH environment variable

## Step 1: Compile your GO file to WebAssembly
The bellowed command is executed in the root folder in order to compile the `main.go` to get the wasm equivalent `main.wasm` in the `assets` folder.
Note that tinygo is used to get ***tiny*** wasm file size. In my case it tooks only **181KB** for the `main.wasm` file size wherease if we use the built in go way it generates wasm file size with almost **2MB** of size.

```shell
tinygo build -o ./assets/main.wasm -target wasm ./main.go
```

## Step 2: Launch the server
In order to launch the server to serve the `assets` folder (`index.html`), run the bellow command in the `./server` folder:
```shell
go run server.go
```

## Step 3: Launch the app
Open your browser and head to: `http://localhost:9090`
