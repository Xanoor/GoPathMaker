# 🛠️GoPathMaker.go

**GoPathMaker.go** is a Go utility that compiles a Go script and adds its directory to the system's PATH environment variables, allowing you to run your compiled program from any terminal.

## Technologies Used

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) 

## Features

- Compiles a specified Go file into an executable.
- Automatically adds the directory of the executable to the system's PATH environment variables.
- Supports handling of certain special characters in error messages.

## Prerequisites

Before using this program, ensure the following are installed on your machine:

- [Go](https://golang.org/dl/) (Go programming language)
- Administrator privileges (to modify PATH environment variables)

## Usage

### Steps to use GoPathMaker.go

1. **Clone this repository** or download the files into a local directory:

   ```
   git clone https://github.com/Xanoor/GoPathMaker
   ```


2. **Edit the variables in the Go script**:

   Open the `GoPathMaker.go` file and modify the variables according to your needs:

   - `filePathName`: Path to your Go file that you want to compile. (e.g., `"your_app.go"`).
   - `outputExe`: Name of the generated executable (e.g., `"MyApp.exe"`).

   > **⚠️ Note:** If you clone this repository directly inside your project folder, the Go file you want to compile will be located in the parent directory relative to `GoPathMaker.go`. In that case, you will need to add `../` before the file name. For example, if the Go file to compile is named `your_app.go`, you should set `filePathName` to `"../your_app.go"`.
3. **Run the script** with administrative privileges:

   - **Windows**:
   
     Open a terminal **as Administrator** and run the following command:
     
     ```
     go run GoPathMaker.go
     ```
4. If everything works fine, you will see messages indicating that the Go file has been compiled and the path has been added to the PATH.

5. **Verification**: After execution, ***open a new terminal*** and type the name of your executable (`outputExe`) to check if it is now accessible globally through PATH.



### Example Usage

Suppose you want to compile a Go file named `your_app.go` and name the executable `MyGoApp.exe`:

1. Modify the script as follows:

   ```go
   var filePathName = "../your_app.go"
   var outputExe = "MyGoApp"
   ```

2. Run the program as administrator.

   ```
   go run GoPathMaker.go
   ```

3. Once complete, you can launch `MyGoApp` from any terminal.

## How It Works
   
1. **Compilation**: The specified Go file is compiled into an executable using the `go build` command.

2. **Add to PATH**: The directory where the compiled executable resides is added to the system's `PATH` environment variable, allowing the program to be executed from any terminal.

## Required Permissions

- **Administrator Permissions**: You must run this program with administrator privileges because it modifies the system’s PATH environment variables.

### Potential Issues

- **Error when adding to PATH**: If you encounter permission issues, ensure you run the script with **administrator privileges**. You will also receive a message if a permission is denied.
  
- **Path Already Exists**: If the directory is already in the `PATH` environment variable, the program will display a warning indicating that the path already exists.

## Warnings

- **Use at Your Own Risk**: This script modifies system environment variables, so be cautious when using it, especially on production systems.
- **Windows Only**: This script uses the `setx` command, which is specific to Windows to modify environment variables. If you're working on Unix or Linux, you'll need to adjust the logic to use `export` or another suitable method.

## Contributions

Contributions are welcome. If you have improvements to offer, please submit a pull request or open an issue for discussion.

## Author

- **Xanoor** - Find me on [GitHub](https://github.com/Xanoor) or xanoor1 on discord.
