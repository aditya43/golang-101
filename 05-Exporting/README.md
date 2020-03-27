## NOTE

- `printer.go` file must reside under `/printer` directory. This is because the name of this `library package` is `printer`.
- Only `Func` names starting with `upper case` letter will get exported.

## To Execute

- Go to:
    ```sh
    cd src/github.com/aditya43/golang/05-Exporting/printer/cmd/
    ```
- Run command:
    ```sh
    go run main.go
    ```
- **NOTE:** `CMD+e+e` (Run Code extension in VS Code) won't work directly on `main.go`. We must have to execute `go run main.go` command.