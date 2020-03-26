# The Go Programming Language
My personal notes and projects.

## Author
Aditya Hajare ([Linkedin](https://in.linkedin.com/in/aditya-hajare)).

## Current Status
WIP (Work In Progress)!

## License
Open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).

----------------------------------------

## Important Notes
- [Basics](#basics)
    ```diff
    + Basic Commands
    ```

----------------------------------------

## Basics
```diff
+ Basic Commands
```
- Print `Environment Variables`:
    ```sh
    go env
    ```
- To format all codes in a directory recursively:
    ```sh
    # 'cd' into project root directory and execute:
    go fmt ./...
    ```
- To run/execute a `.go` file:
    ```sh
    go run FILE_NAME
    ```
- To build an executable:
    ```sh
    go build FILE_NAME
    ```
    * For a file:
        - Builds the file.
        - Report errors, if any.
        - If there are no errors, it puts an executable into the current directory.
    * For a package:
        - Builds the file.
        - Report errors, if any.
        - Throws away binary.
- To compile and install:
    ```sh
    go install FILE_NAME
    ```
    * For an executable:
        - Compiles the program (builds it).
        - Names the executable
            * `Mac`: The foldername holding the code.
            * `Windows`: File name
        - Puts the executable in `workspace/bin`.
            * `$GOPATH/bin`
    * For a package:
        - Compiles the package (builds it).
        - Puts the executable in `workspace/pkg`.
            * `$GOPATH/pkg`
        - Makes it an archive file.