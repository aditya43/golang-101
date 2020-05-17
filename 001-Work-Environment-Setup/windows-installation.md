# WINDOWS INSTALLATION CHEATSHEET

## NOTE

If you have [chocolatey.org](https://chocolatey.org/) package manager, you can easily install Go like so:

```
choco install golang
```

## 1- Install Visual Studio Code Editor

1. Install it but don't open it yet.
2. Go to: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Select Windows and start downloading
4. Run the installer

## 2- Install Git

1. Grab and run the installer. Go to: [https://gitforwindows.org](https://gitforwindows.org)
2. Select VSCode as the default editor
3. Enable all the checkboxes
4. Select: _"Use Git from the Windows Command Prompt"_
5. Encodings: Select: _"Checkout as is..." option._

## 3- Install Go

1. Go to [https://golang.org/dl](https://golang.org/dl)
2. Select Windows and download
3. Start the installer

## 4- Configure VS Code

1. Open VS Code; from the extensions tab at the left, search for "go" and install it
2. Close VS Code completely and open it up again

3. Go to View menu; select **Command Palette**
    1. Or just press `ctrl+shift+p`
    2. Type: `go install`
    3. Select _"Go: Install/Update Tools"_
    4. Check all the checkboxes

## 5- Install Git-Bash

* [Linux Subsystem for Windows](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **So, to use git bash, follow these steps:**
    1. Setup VS Code to use git-bash by default:
        1. Open VS Code
        2. Go to Command Palette
            1. Type: `terminal`
            2. Select: _"Terminal: Select Default Shell"_
            3. And, Select: _"Git Bash"_

    4. **NOTE:** Normally, you can find your files under `c:\`, however, when you're using git bash, you'll find them under `/c/` directory. It's actually the very same directory, it's just a shortcut.

## That's all! 🤩
