package writesubdirnamestofile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	outputFile   = "output.txt"
	workingDir   = "writesubdirnamestofile"
	fullWorkPath = workingDir + "/" + outputFile // the plan, is that this could be a relative or complete path, but this works for now
	dueChecking  []string                        // por utilizar
)

// WriteSubDirNamesToFile testing
func WriteSubDirNamesToFile() {

	paths := os.Args[1:]
	if len(paths) < 1 {
		log.Panic("not enough folder paths to process, \nplease add a folder to read the diretories name")

	}

	var dirs []byte

	for _, dir := range paths {

		file, err := ioutil.ReadDir(dir)

		if err != nil {
			log.Panic("Error: ", err.Error())
		}

		dirs = append(dirs, dir+"\n"...)

		for _, fileData := range file {
			if fileData.IsDir() {

				dirs = append(dirs, "\t"+fileData.Name()+"\n"...)

				secondPath := dir + fileData.Name() + "/"

				file2, err := ioutil.ReadDir(secondPath)

				if err != nil {
					log.Fatal("Error:", err.Error())
				}

				for _, fileData2 := range file2 {
					dirs = append(dirs, "\t \t"+fileData2.Name()+"\n"...)
				}

			}
		}

	}

	ioutil.WriteFile(fullWorkPath, dirs, 0644)
	fmt.Printf("\n%s \n", dirs)

}
