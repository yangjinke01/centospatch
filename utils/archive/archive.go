package archive

import (
	"centospatch/utils/command"
	"centospatch/utils/path"
	"log"
)

func UnpackTarball(packagePath string, UnpackDir string) {
	if path.Exists(packagePath) {
		command.Execute("tar", "-xf", packagePath, "-C", UnpackDir)
	} else {
		log.Println(packagePath + " not exist")
	}
}
