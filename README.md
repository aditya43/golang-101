# The Go Programming Language :beer:
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
    + Packages
    - Executable Packages
    - Library Packages

    + Scopes
    ```

----------------------------------------

## Basics
```diff
+ Packages
```
- All package files, should be in the same (single) directory. i.e. all package source code files should be located in a one single directory.
- All files in a specific folder should belong to a one single package. It's a convention, not a rule.
> There are 2 kinds of packages in Go: `Executable Packages` and `Library Packages`.
- To make a package executable, name that package `main`. It's a special package.
- `package` clause can be used only once per file and it should be the first line in `.go` source file.
- Package contains multiple `Go` files belonging to same folder.

```diff
- Executable Packages
```
- It's name should always be `package main`.
- `Executable Package` should also contain `main()` function and that too only once.
- These are created only for `running` it as a Go program.
- These cannot be imported into a Go program.
- Package name should be `main`.

```diff
- Library Packages
```
- Almost all `Go Standard Library Packages` are of type `Library Packages`.
- They are reusable packages.
- They are not executable packages. So we can't run them.
- We can only `import` them.
- These are created only for `reusability` purposes.
- Package name can have any name.
- Doesn't need to have function named `main()`. To avoid confusion, it's better not to have function named `main()` in a reusable package.

```diff
+ Scopes
```
- Same name cannot be declared again inside a same scope.
- There are following types of scopes in Go:
    1. `package`: Each Go package has it's own `scope`. For e.g. declared `funcs` are only `visible` to the files belonging to same `package`.
    2. `file`: Imported packages are only visible to the importing file. Each file has to import external packages on it's own.
    3. `func`.
    4. `block`.