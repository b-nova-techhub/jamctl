package repo

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadRepoContents(repo string) []string {
	files, mdErr := getAllMdFilesInPath(repo)
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

func validateRepoUrl(repoUrl string) string {
	if !strings.HasPrefix(repoUrl, "https://") {
		repoUrl = "https://" + repoUrl
	}
	fmt.Printf("Repository URL: ", repoUrl)
	return repoUrl
}

func GetGitRepository(repoUrl string) {
	projectPath := viper.GetString("absolutePath") + "/" + getSlug(repoUrl)
	fmt.Printf("Target repository clone path:", projectPath)
	pathExists, pathErr := pathExists(projectPath)

	if pathErr != nil {
		log.Fatalf("Error before git clone as path already seems to exist: %+x\n", projectPath)
	}

	if !pathExists {
		cloneErr := cloneToFilesystem(projectPath, repoUrl)
		if cloneErr != nil {
			log.Fatalf("Error during git clone. Path: %+x\n", projectPath)
		}
	}
}

func cloneToFilesystem(path, url string) error {
	log.Println("Git clone to path: ", path)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
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

func getAllMdFilesInPath(projectName string) ([]string, error) {
	absoluteContentPath := viper.GetString("absolutePath") + "/" + projectName + viper.GetString("relativePath")
	return walkMatch(absoluteContentPath, "*.md")
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

func getSlug(s string) string {
	slices := strings.Split(s, "/")
	return slices[len(slices)-1]
}
