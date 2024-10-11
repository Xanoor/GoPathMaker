package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// VARIABLES TO MODIFY //
var filePathName = "../your_app.go"
var outputExe = "GoPathMakerExample"

// ///////////////////////

var specialChar = map[byte]byte{
	130: byte('e'), // é
	138: byte('e'), // è
	144: byte('E'), // É
	235: byte('e'), // ë
	195: byte('?'), // others
	255: byte(' '),
}

func verifyExtension() {
	if !strings.HasSuffix(filePathName, ".go") {
		fmt.Printf("\n[INFO] Extension .go appended to %s", filePathName)
		filePathName += ".go"
	}
	if !strings.HasSuffix(outputExe, ".exe") {
		fmt.Printf("\n[INFO] Extension .exe appended to %s", outputExe)
		outputExe += ".exe"
	}
}

func compileGoScript() bool {
	_, err := os.Stat(filePathName)
	if err != nil {
		fmt.Println("\n[FAIL] Error retrieving the path to the file.")
		return false
	}
	cmd := exec.Command("go", "build", "-o", outputExe, filePathName)
	err = cmd.Run()

	if err != nil {
		fmt.Printf("\n[FAIL] Error compiling %s! (Build command)", filePathName)
		return false
	}
	fmt.Printf("\n[DONE] File %s has been compiled!", filePathName)
	return true
}

func removeSpecialChar(text string) string {
	var outputText = ""
	for i := range text {
		if f, found := specialChar[text[i]]; found {
			outputText += string(f)
		} else {
			outputText += string(text[i])
		}
	}

	return outputText
}

func addDirToPathEnv() {
	folderPath, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("\n[FAIL] Problem creating the local path for the file: %s!", filePathName)
		return
	}

	if !isEnvPathExist(folderPath) {
		pathList := os.Getenv("PATH")
		newEnvPath := pathList + ";" + folderPath

		cmd := exec.Command("setx", "Path", newEnvPath, "/M")
		output, err := cmd.CombinedOutput()
		file, _ := os.Stat("GoPathMaker.go")

		if err != nil {
			fmt.Printf("\n\n[FAIL] Error when adding folder to the PATH environment variable!")
			fmt.Printf("\n[INFO] File authorization: %s", file.Mode())
			fmt.Printf("\n[INFO] Error returned: %s", err)
			fmt.Printf("\n[INFO] Output message: %s", removeSpecialChar(string(output)))
			fmt.Printf("[NOTE] You must execute `go run GoPathMaker.go` with administrator privileges!")
			return
		}
		fmt.Printf("\n[OUTPUT] %s", removeSpecialChar(strings.TrimSpace(string(output))))
		fmt.Printf("\n[DONE] Folder successfully added to the PATH environment variables!")
	} else {
		fmt.Println("\n[WARN] PATH environment variable already exists!")
	}
}

func isEnvPathExist(path string) bool {
	pathList := os.Getenv("PATH")
	pathListArr := strings.Split(pathList, ";")

	for i := range pathListArr {
		if pathListArr[i] == path {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(`
   ____       ____       _   _     __  __       _             
  / ___| ___ |  _ \ __ _| |_| |__ |  \/  | __ _| | _____ _ __ 
 | |  _ / _ \| |_) / _  | __| _  \| |\/| |/ _  | |/ / _ \ '__|
 | |_| | (_) |  __/ (_| | |_| | | | |  | | (_| |   <  __/ |   
  \____|\___/|_|   \__,_|\__|_| |_|_|  |_|\__,_|_|\_\___|_|   
                      Created by Xanoor https://github.com/Xanoor    
					  
   [WARN] Do not modify this script if you do not know what you are doing!
   [WARN] Variables to modify: 
   [WARN] var filePathName (set this to the path of your Go file; see example)
   [WARN] var outputExe (set this to the future name of your executable file)

		`)

	verifyExtension()
	if compileGoScript() {
		addDirToPathEnv()
	}
}
