package dir

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andodeki/gen_project/aOK_GenerateCode/project"
	"github.com/andodeki/gen_project/aOK_GenerateCode/util"
)

const (
	Wdir = "../"
)

func Generate_Dir() {
	GenerateProjectName_Dir(Wdir + project.ProjectName + `/`)
	GenerateProjects_SubDir(Wdir + project.ProjectName + `/`)
}
func GenerateProjects_SubDir(name string) {
	Directories := PopulatedDirectoriesSlice()
	for _, dir := range Directories {
		if err := util.EnsureDir(name + dir); err != nil {
			log.Println("Directory creation failed with error: " + err.Error())
		}
		tdir := strings.Split(dir, "/")[0]
		if cherr := os.Chmod(Wdir+project.ProjectName+`/`+tdir, 0700); cherr != nil {
			log.Printf("cannot chmod to file 1:%v", cherr)
		}
		if cherr := os.Chmod(Wdir+project.ProjectName+`/`+dir, 0700); cherr != nil {
			log.Printf("cannot chmod to file 2:%v", cherr)
		}
	}
}

func GenerateProjectName_Dir(name string) {
	if err := util.EnsureDir2(name); err != nil {
		log.Println("Directory creation failed with error: " + err.Error())
	}
}

func Create_Files(filename, generating_file string) {
	Files := PopulatedFileSlice()
	for _, file := range Files {
		if strings.Contains(file, generating_file) {
			// dirName := Wdir + project.ProjectName + `/src/` + generating_file
			// if cherr := os.Chmod(dirName, 0700); cherr != nil {
			// 	log.Printf("cannot chmod to file 1c:%v", cherr)
			// }
			util.C_WFile(filename, Wdir+project.ProjectName+"/src/"+file)
		}

	}
}

func FormStringDir(projectName string, directories []string) (sb strings.Builder) {
	for _, dir := range directories {
		path := fmt.Sprintf("github.com/andodeki/%s/%s", projectName, dir)
		sb.WriteString(`"` + path + `"` + "\n\t")
	}
	return sb
}
