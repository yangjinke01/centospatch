/*
Copyright © 2022 Jack <yangjinke80@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"centospatch/constants"
	"centospatch/utils/archive"
	"centospatch/utils/command"
	"centospatch/utils/path"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// opensslCmd represents the openssl command
var opensslCmd = &cobra.Command{
	Use:   "openssl",
	Short: "compile and install openssl 3.0.2",
	Long:  `compile and install openssl 3.0.2`,
	Run: func(cmd *cobra.Command, args []string) {
		archive.UnpackTarball(constants.PatchesDir+constants.OpensslSourcePackage, constants.UnpackingDir)
		compileSsl(constants.UnpackingDir + constants.OpensslSourceUnpackName)
	},
}

func init() {
	rootCmd.AddCommand(opensslCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opensslCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opensslCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func compileSsl(srcDir string) {
	defer command.Execute("rm", "-rf", srcDir)
	_ = os.Chdir(srcDir)
	pwd, _ := os.Getwd()
	fmt.Printf("Current Working Direcotry: %s\n", pwd)
	command.Execute("./config")
	command.Execute("make")
	command.Execute("make", "install")
	command.Execute("ln", "-sf", "/usr/local/lib64/libssl.so.3", "/usr/lib64/libssl.so.3")
	command.Execute("ln", "-sf", "/usr/local/lib64/libcrypto.so.3", "/usr/lib64/libcrypto.so.3")
	if path.Exists("/usr/bin/openssl") {
		command.Execute("mv", "-f", "/usr/bin/openssl", "/usr/bin/openssl_old")
		command.Execute("ln", "-sf", "/usr/local/bin/openssl", "/usr/bin/openssl")
	}
	command.Execute("openssl", "version")
}
