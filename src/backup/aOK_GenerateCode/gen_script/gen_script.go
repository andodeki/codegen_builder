package gen_script

import (
	"github.com/andodeki/gen_project/aOK_GenerateCode/dir"
)

const (
	header = `#!/bin/sh

go mod init && go mod tidy
`
)

func Generate_ScriptContents() {
	// sb := dir.FormStringDir(project.ProjectName, StartAppImports)

	// startup_ := fmt.Sprintf(header, sb.String())
	dir.Create_Files(header, "scripts/gomod.sh")

}
