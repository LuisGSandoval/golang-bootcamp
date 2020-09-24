package writesubdirnamestofile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	outputFile   = "output.txt"
	workingDir   = "writesubdirnamestofile"
	fullWorkPath = workingDir + "/" + outputFile // the plan, is that this could be a relative or complete path, but this works for now
	dueChecking  []string                        // por utilizar
)

// WriteSubDirNamesToFile testing
func WriteSubDirNamesToFile() {

	if len(os.Args) < 2 {
		fmt.Println("Please add 1 directory path")
		return
	}

	args := os.Args[1:]

	fileRoute := args[0]

	err := os.Remove(fullWorkPath) // this removes the file where the info is going to be written on

	if err != nil {
		fmt.Println("[Error] WriteSubDirNamesToFile():", err.Error())
		return
	}

	writeDirNames(fileRoute)

}

// This is the main function of the app
func writeDirNames(fileRoute string) {

	checkFileExistance()

	for _, data := range readDirInfo(fileRoute) {

		fileContent, err := ioutil.ReadFile(fullWorkPath)

		if err != nil {
			fmt.Println("[Error] writeDirNames: ", err.Error())
			return
		}

		if data.IsDir() {
			fullFilePath, err := filepath.Abs(data.Name())
			if err != nil {
				fmt.Println("[Error], writeDirNames: ", err.Error())
			}

			dueChecking = append(dueChecking, fullFilePath)
			fileContent = append(fileContent, data.Name()+"/ \n"...)
			ioutil.WriteFile(fullWorkPath, fileContent, 0644)
		}

		for _, filePathToCkeck := range dueChecking {
			writeDirNames(filePathToCkeck)
		}
	}

}

// this file ensures that the file where we're going to write, acually exists if not, it creates the file
func checkFileExistance() {

	var existance bool

	fileInfo, err := ioutil.ReadDir(workingDir)

	if err != nil {
		fmt.Println("[Error] checkFileExistance:", err.Error())
		return
	}

	for _, file := range fileInfo {
		if file.Name() == outputFile {
			existance = true
			break
		}
	}

	if !existance {
		ioutil.WriteFile(fullWorkPath, []byte{}, 0644)
	}
}

func readDirInfo(fileRoute string) []os.FileInfo {

	fileInfo, err := ioutil.ReadDir(fileRoute)

	if err != nil {
		fmt.Println("[Error] readDirInfo: ", err.Error())
		return nil
	}

	return fileInfo
}
