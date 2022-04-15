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
	"github.com/spf13/cobra"
	"log"
	"os"
)

// opensshCmd represents the openssh command
var opensshCmd = &cobra.Command{
	Use:   "openssh",
	Short: "compile and install openssh 9.0p1",
	Long:  `compile and install openssh 9.0p1`,
	Run: func(cmd *cobra.Command, args []string) {
		archive.UnpackTarball(constants.PatchesDir+constants.OpensshSourcePackage, constants.UnpackingDir)
		compileSsh(constants.UnpackingDir + constants.OpensshSourceUnpackName)
	},
}

func rmOldService() {
	if path.Exists("/usr/lib/systemd/system/sshd.service") {
		command.Execute("mv", "-f", "/usr/lib/systemd/system/sshd.service", "/etc/ssh_old/sshd.service")
	}
	if path.Exists("/usr/lib/systemd/system/sshd.socket") {
		command.Execute("mv", "-f", "/usr/lib/systemd/system/sshd.socket", "/etc/ssh_old/sshd.socket")
	}
}

func addNewService(srcDir string) {
	command.Execute("cp", "-af", srcDir+"/contrib/redhat/sshd.init", "/etc/init.d/sshd")
	command.Execute("/etc/init.d/sshd", "restart")
	command.Execute("systemctl", "daemon-reload")
	command.Execute("chkconfig", "--add", "sshd")
	command.Execute("chkconfig", "sshd", "on")
	command.Execute("ssh", "-V")
}

func preCheck() {
	if !path.Exists("/etc/ssh_old") {
		command.Execute("rm -rf", "/etc/ssh_old")
		command.Execute("mv", "/etc/ssh", "/etc/ssh_old/")
	}
}

func compile(srcDir string) {
	_ = os.Chdir(srcDir)
	pwd, _ := os.Getwd()
	log.Printf("Current Working Direcotry: %s\n", pwd)
	command.Execute("./configure", "--prefix=/usr/", "--sysconfdir=/etc/ssh",
		"--with-ssl-dir=/usr/local/lib64/", "--with-ssl-engine")
	command.Execute("/bin/bash", "-c", "make && make install")
}

func compileSsh(srcDir string) {
	preCheck()
	compile(srcDir)
	rootLogin()
	rmOldService()
	addNewService(srcDir)
}

func rootLogin() {
	command.Execute("sed", "-i", "/^PermitRootLogin/d", "/etc/ssh/sshd_config")
	command.Execute("echo", "PermitRootLogin yes", ">>", "/etc/ssh/sshd_config")
}

func init() {
	rootCmd.AddCommand(opensshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// opensshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// opensshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
