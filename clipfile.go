package clipfile

import (
	"bytes"
	"log"
	"os/exec"
)

const coreScripts = `tell application "System Events" to set filePath to (the clipboard as «class furl») as alias`

func GetPath() string {
	cmd := exec.Command("osascript", "-e", coreScripts, "-e", `POSIX path of filePath`)

	// Capture the output of the command
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Trim any newline or whitespace from the path
	filePath := out.String()
	filePath = filePath[:len(filePath)-1] // Remove trailing newline character

	return filePath
}
