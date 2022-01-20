package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	codePath    string = os.Getenv("CODE_PATH")
	toolboxPath string = os.Getenv("TOOLBOX_PATH")
	issues      []Issue
)

func main() {
	analysisFiles, err := getAllFiles()
	if err != nil {
		log.Fatalln("Failed to read files  to analyze. Exiting...")
	}

	for _, path := range analysisFiles {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), `\n`)
		for lineNumber, line := range lines {
			if strings.Contains(line, "TODO") {
				createIssue(path, lineNumber, 0)
			}
		}
	}
	macroAnalysisResult := prepareResult()

	if writeError := writeMacroResult(macroAnalysisResult); writeError != nil {
		log.Fatalln("Error occured while writing  results :", writeError)
	}
}
