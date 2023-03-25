package loader

import (
	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"github.com/pkg/errors"
)

type Compose struct{}

func (c *Compose) LoadFile(files []string) (project *types.Project, err error) {
	workingDir, err := getComposeFileDir(files)
	if err != nil {
		return
	}
	projectOptions, err := cli.NewProjectOptions(files, cli.WithOsEnv, cli.WithWorkingDirectory(workingDir), cli.WithInterpolation(false))
	if err != nil {
		err = errors.Wrap(err, "Unable to create compose options")
		return
	}

	project, err = cli.ProjectFromOptions(projectOptions)
	if err != nil {
		err = errors.Wrap(err, "Unable to load files")
		return
	}
	return
}
