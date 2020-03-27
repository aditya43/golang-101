# First Go Program & Basics

## BUILDING & RUNNING GO PROGRAMS:

* **Build a Go program:**

    * While inside a program directory, type:
        * `go build`

* **Run a Go program:**

    * While inside a program directory, type:
        * `go run main.go`

## WHAT IS $GOPATH?

* _$GOPATH_ is an environment variable which points to a directory where the downloaded and your own Go files are stored.

    * **On Windows**, it's in: `%USERPROFILE%\go`

    * **On OS X & Linux**, it's in: `~/go`

    * **NOTE:** Never set your GOPATH manually. It's by default under your users directory.

* **GOPATH has 3 directories:**

    * **src:** Contains the source files for your own or other downloaded packages. You can build and run programs while in a program directory under this directory.

    * **pkg:** Contains compiled package files. Go uses this to make the builds (compilation & linking) faster.

    * **bin:** Contains compiled and executable Go programs. When you call go install on a program directory, Go will put the executable under this folder.

        * _You might want to add this to your `PATH` environment variable if it's not there already._

## WHERE YOU SHOULD PUT YOUR SOURCE FILES?

* `$GOPATH/src/github.com/yourUsername/programDirectory`

* **Example:**

    * My GitHub username is: aditya43

    * So, I put all my programs under: `~/go/src/github.com/aditya43/`

    * And, let's say that I've a program named hello, then I put it under this directory: `~/go/src/github.com/aditya43/hello`

## FIRST PROGRAM

* **Create directories:**
    * **OS X & Linux (or git bash):**
        * Create a new directory:
            * `mkdir -p ~/go/src/yourname/hello`
        * Go to that directory:
            * `cd ~/go/src/yourname/hello`

    * **Windows:**
        * Create a new directory:
            * `mkdir c:\Go\src\yourname\hello`
        * Go to that directory:
            * `cd c:\Go\src\yourname\hello`

* Create a new `code main.go` file under Visual Studio Code.
* And add the following code to it and save it.
* Then, return back to the command-line.
    * Run it like this: `go run main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Aditya.")
}
```

## NOTE:

* There is a new *Go Modules* support that allows you to run your programs in any directory that you want.