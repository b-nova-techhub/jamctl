package commands

import (
	"b-nova-techub/cobra-demo/pkg/gen"
	"b-nova-techub/cobra-demo/pkg/repo"
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update git repository containing markdown content files",
		Long:  ``,

		Run: update,
	}
)

func init() {
	includeUpdateFlags(getCmd)
}

func includeUpdateFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&absolutePath, "absolutePath", "", "/tmp", "The absolute path where the git repository is going to cloned to.")
	cmd.Flags().StringVarP(&relativePath, "relativePath", "", "/content", "The directory within the git repository which contains the markdown files.")
	cmd.Flags().StringVarP(&branch, "branch", "b", "main", "The branch of the git repository that is to be used.")
	cmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "The tag that is being used as the front matter delimiter.")
	cmd.Flags().BoolVarP(&overwrite, "overwrite", "o", true, "Whether or not to overwrite an existing git repository.")
}

func update(ccmd *cobra.Command, args []string) {
	gen.Generate(repo.RepoContents())
}
