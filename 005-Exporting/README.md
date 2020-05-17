## NOTE

- About `printer` library package:
    * `printer.go` file must reside under `/printer` directory. This is because the name of that `library package` is `printer`.
- About `goversion` library package:
    * `goversion.go` file must reside under `/goversion` directory. This is because the name of that `library package` is `goversion`.
- Only `Func` names starting with `upper case` letter will get exported.

## To Execute

- **NOTE:** `CMD+e+e` (Run Code extension in VS Code) won't work directly on `main.go`. We must have to execute `go run main.go` command.
- `printer` library package:
    * Go to:
        ```sh
        cd src/github.com/aditya43/golang/005-Exporting/printer/cmd/
        ```
    * Run command:
        ```sh
        go run main.go
        ```
- `goversion` library package:
    * Go to:
        ```sh
        cd src/github.com/aditya43/golang/005-Exporting/goversion/cmd/
        ```
    * Run command:
        ```sh
        go run main.go
        ```
