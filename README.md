# :beer: Go Language :beer:
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
    + Renaming Imports
    + Exporting
    + Data Types
        - Basic Data Types
    + Variables
        - Zero Values
        - Unused variables
        - Multiple Declarations
        - Type Inference
        - Short Declaration
    + Blank Identifier
    ```
- [Naming Conventions In Go](07-Naming-Conventions/README.md#naming-rules-in-go-language)  :arrow_upper_right:

----------------------------------------

## Basics
- Go is a `strongly typed` language. Because of that, it helps Go compiler to identify many types of errors at `compile time` even before our program is run.

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

```diff
+ Renaming Imports
```
- We can rename an imported package name with following syntax:
    ```go
    package main

    import "fmt"
    import adi "fmt"    // Imported "fmt" package and renamed it to "adi"

    func main() {
        adi.Println("नमस्ते आदित्य")    // This will print "नमस्ते आदित्य"
    }
    ```
- We can import packages with the same name into same file by giving one of them imports a new name.

```diff
+ Exporting
```
- To export a name in Go, just make it's **first letter** an **uppercase letter**.
- For e.g.
    ```go
    package aditest

    func Adi() { // 'Adi()' will be exported and will be available throughout 'aditest' package
        // Code..
    }

    func adiNew() { // 'adiNew()' will not be exported since it's name doesn't start with uppercase letter.
        // Code
    }
    ```

```diff
+ Data Types
```
- `literal` means the `value` itself. Unline `variable`, a `literal` doesn't have a name.
- There are following data types in Go:
    * **Basic type**: Numbers, strings, and booleans come under this category.
    * **Aggregate type**: Array and structs come under this category.
    * **Reference type**: Pointers, slices, maps, functions, and channels come under this * category.
    * **Interface type**

```diff
- Basic Data Types
```
- Following are the basic data types in Go:
    * **Numeric**:
        ```go
        // Integer Types
        uint8   // Unsigned 8-bit integers (0 to 255)
        uint16  // Unsigned 16-bit integers (0 to 65535)
        uint32  // Unsigned 32-bit integers (0 to 4294967295)
        uint64  // Unsigned 64-bit integers (0 to 18446744073709551615)
        int8    // Signed 8-bit integers (-128 to 127)
        int16   // Signed 16-bit integers (-32768 to 32767)
        int32   // Signed 32-bit integers (-2147483648 to 2147483647)
        int64   // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

        // +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

        // Floating Types
        float32     // IEEE-754 32-bit floating-point numbers
        float64     // IEEE-754 64-bit floating-point numbers
        complex64   // Complex numbers with float32 real and imaginary parts
        complex128  // Complex numbers with float64 real and imaginary parts

        // +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

        // Other Numeric Types
        byte     // same as uint8
        rune     // same as int32
        uint     // 32 or 64 bits
        int      // same size as uint
        uintptr  // an unsigned integer to store the uninterpreted bits of a pointer value
        ```
    * **Boolean**:
        ```go
        bool    // Represents 'true' or 'false'
        ```
    * **String**:
        - In Go language, strings are different from other languages like Java, C++, Python, etc.
        - Strings can't be `null` in Go.
        - It is a sequence of variable-width characters where each and every character is represented by one or more bytes using UTF-8 Encoding.
        - In Go, a `string` is in effect is a read-only slice of bytes (immutable).
        - Or in other words, strings are the immutable chain of arbitrary bytes (including bytes with zero value) and the bytes of the strings can be represented in the Unicode text using UTF-8 encoding.
        - String literals can be created in 2 ways:
            * **Using double quotes**
            * **Using backticks**

```diff
+ Variables
```
- [Variables in Go Lang](https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe)
- In Go, we have to **declare** a **variable before** we can **use it**. This is required and necessary for the `compile time safety`.
- Variables are not created at `compile time`. They are created at `run time`.
- The unnamed variables are `pointers` (like in C).
- Once we declare a type for a variable, it cannot be changed later. It is static.

```diff
- Zero Values
```
- When a variable is declared and it isn't assigned any value at the time of declaration, Go will assign a `zero value` to it based on it's variable type.
- Type of a variable decides what `zero value` it will take initially when declared (and if it isn't assigned any value at the time of declaration).
    ```go
    // Zero Values assigned to variables by Go when they are declared and not assigned any values at the time of declaration.
    var adiBool bool          // false
    var adiInt int            // 0
    var adiFloat float64      // 0
    var adiStr string         // ""
    var adiPointer *string    // nil | 'nil' means it doesn't point to any memory location
    ```

```diff
- Unused variables
```
- **Unused variables in `blocked scope` are not allowed in Go since they cause `maintenance nightmares`.** If we declare a variable in `blocked scope` then we must use it or else completely remove it from the block. We cannot have unused variables declared in `blocked scope` dangling in our source codes. Go throws unused variable errors at `compile time` only.
- **We should avoid using `package level` variables.** Go doesn't throw `unused variable errors` at `compile time` for variables declared at `package level`.

```diff
- Multiple Declarations
```
- Sometimes it is also called as parallel variable declarations.
- Declaring multiple variables with `different types` in a single statement:
    ```go
    package main

    func main() {
        var (
            adiBool bool
            adiInt int
            adiFloat float64
            adiStr string
            adiPointer *string
        )
    }
    ```
- Declaring multiple variables with `same type` in a single statement:
    ```go
    package main

    func main() {
        var foo, bar, baz int
    }
    ```

```diff
- Type Inference
```
- `Type Inference` means Go can figure out the type of a variable automatically from it's assigned value.
- If we are assigning value to a variable at the time of it's declaration, we can ommit it's `type` specification.
- For e.g.
    ```go
    package main

    main() {
        var someFlag = true // We are not specifying type of 'someFlag' as bool here.
    }
    ```

```diff
- Short Declaration
```
- With `Type Inference`, Go can figure out variable type based off it's assigned value.
- In `Short Declaration`, we can declare variable by completely ommitting `var` keyword along with it's variable type.
- It `declares` and `initializes` the variable.
- **We cannot use `Short Declaration` syntax to declare variables in `Package Scope`.**
- At `Package Scope`, all declarations should start with a `keyword`. Since `Short Declaration` syntax doesn't have any `keyword` in it, it doesn't work at `Package Scope`.
- For e.g.
    ```go
    package main

    main() {
        someFlag := true // 'var' keyword and 'variable type' is not specified. It works!
    }
    ```

```diff
+ Blank Identifier
```
> “There are only two hard things in Computer Science: cache invalidation and naming things”. Tim Bray quoting Phil Karlton
- Go doesn't allow `unused variables` in `blocked scope`.
- To ignore a variable, `Blank Identifier (_)` is used as a variable name in Go.
- Go compiler will not throw unsed variable error if a blocked scope variable is named `_`.
- We cannot use value assigned to `_`.
- It is like a black hole that swallows variable.
- [Detailed information and usage of Blank Identifier](https://golang.org/doc/effective_go.html#blank)