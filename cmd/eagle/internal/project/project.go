package project

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a project template",
	Long:  "Create a project project using the repository template. Example: eagle new helloworld",
	Run:   run,
}

var repoURL string

func init() {
	if repoURL = os.Getenv("eagle_LAYOUT_REPO"); repoURL == "" {
		repoURL = "https://github.com/go-eagle/eagle-layout.git"
	}
	CmdNew.Flags().StringVarP(&repoURL, "-repo-url", "r", repoURL, "layout repo")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	name := ""
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name.",
		}
		survey.AskOne(prompt, &name)
		if name == "" {
			return
		}
	} else {
		name = args[0]
	}
	p := &Project{Name: name}
	if err := p.New(ctx, wd, repoURL); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
}
