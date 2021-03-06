package bundle

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getantibody/antibody/project"
)

type zshBundle struct {
	Project project.Project
}

var zshGlobs = []string{"*.plugin.zsh", "*.zsh", "*.sh", "*.zsh-theme"}

func (bundle zshBundle) Get() (result string, err error) {
	if err = bundle.Project.Download(); err != nil {
		return result, err
	}
	for _, glob := range zshGlobs {
		files, err := filepath.Glob(filepath.Join(bundle.Project.Folder(), glob))
		if err != nil {
			return result, err
		}
		if files == nil {
			continue
		}
		var lines []string
		for _, file := range files {
			lines = append(lines, "source "+file)
		}
		lines = append(lines, fmt.Sprintf("fpath+=( %s )", bundle.Project.Folder()))
		return strings.Join(lines, "\n"), err
	}

	return result, nil
}
