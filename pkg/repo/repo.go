package repo

import (
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func RepoContents() []string {
	getGitRepository(viper.GetString("absolutePath"))
	files, mdErr := getAllMdFilesInPath()
	if mdErr != nil {
		log.Fatalln("Error during markdown files parsing.")
	}

	contentFiles := make([]string, 0)
	for _, f := range files {
		readFile, _ := ioutil.ReadFile(f)
		contentFiles = append(contentFiles, string(readFile))
	}
	return contentFiles
}

func getGitRepository(path string) {
	projectPath := viper.GetString("absolutePath")
	pathExists, pathErr := pathExists(projectPath)

	if pathErr != nil {
		log.Fatalf("Error before git clone as path already seems to exist: %+x\n", projectPath)
	}

	if !pathExists {
		cloneErr := cloneToFilesystem(projectPath)
		if cloneErr != nil {
			log.Fatalf("Error during git clone. Path: %+x\n", projectPath)
		}
	}
}

func cloneToFilesystem(path string) error {
	log.Println("Git clone to path: ", path)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      viper.GetString("repositoryUrl"),
		Progress: os.Stdout,
	})
	return err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getAllMdFilesInPath() ([]string, error) {
	// absoluteContentPath
	return walkMatch(viper.GetString("absolutePath"), "*.md")
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
