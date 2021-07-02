package repo

import (
	"b-nova-openhub/stapagen/pkg/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
	"os"
)

func Clone(toMemory bool, path string) error {
	if toMemory {
		return cloneToMemory()
	}
	return cloneToFilesystem(path)
}

func cloneToMemory() error {
	log.Println("\nGit clone to memory.")
	_, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      config.AppConfig.SourceGitRepositoryUrl,
		Progress: os.Stdout,
	})
	return err
}

func cloneToFilesystem(path string) error {
	log.Println("Git clone to path: ", path)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      config.AppConfig.SourceGitRepositoryUrl,
		Progress: os.Stdout,
	})
	return err
}
