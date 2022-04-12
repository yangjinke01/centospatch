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
	"github.com/spf13/cobra"
)

// commonCmd represents the common command
var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "patch common rpm",
	Long:  `patch common rpm, no need to download the rpm`,
	Run: func(cmd *cobra.Command, args []string) {
		archive.UnpackTarball(constants.PatchesDir+constants.RpmPackage, constants.UnpackingDir)
		installRpm(constants.UnpackingDir + constants.RpmUnpackName + "/*.rpm")
	},
}

func init() {
	rootCmd.AddCommand(commonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func installRpm(packagePath string) {
	command.Execute("/bin/sh", "-c", "yum install -y "+packagePath)
}
