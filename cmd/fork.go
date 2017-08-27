package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tcnksm/go-gitconfig"
	"github.com/zaquestion/lab/internal/git"
	lab "github.com/zaquestion/lab/internal/gitlab"
)

// forkCmd represents the fork command
var forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Fork a remote repository on GitLab and add as remote",
	Long:  ``,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := gitconfig.Local("remote." + lab.User + ".url"); err == nil {
			log.Println("remote:", lab.User, "already exists")
			return
		}

		remoteURL, err := gitconfig.Local("remote.origin.url")
		if err != nil {
			log.Fatal(err)
		}
		if git.IsHub && strings.Contains(remoteURL, "github.com") {
			git := git.New("fork")
			git.Run()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		project, err := git.PathWithNameSpace("origin")
		if err != nil {
			log.Fatal(err)
		}
		remote, err := lab.Fork(project)
		if err != nil {
			log.Fatal(err)
		}

		git := git.New("remote", "add", lab.User, remote)
		git.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(forkCmd)
}