package download

import (
	"os/exec"
	"os"
)

// Download
func Run(options Options) {

	// Set the commands.
	download := "curl -OL http://wordpress.org/latest.tar.gz && "
	unzip := "tar xzf latest.tar.gz && "
	relocate := "mv wordpress/* "+ options.Destination +" && "
	cleanup := "rm -rf latest.tar.gz wordpress;"

	// Concatenate the command.
	command := download + unzip + relocate + cleanup

	// Prepare the commands.
	cmd := exec.Command("/bin/sh", "-c", command)

	// Show the command output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the commands.
	err := cmd.Run()

	// Check for errors.
	if err != nil {
		panic("Something went wrong.");
	}
}