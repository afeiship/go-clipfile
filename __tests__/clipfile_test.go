package clipfile

import (
	"fmt"
	"testing"

	"github.com/afeiship/go-clipfile"
)

func TestGetPath(f *testing.T) {
	path := clipfile.GetPath()

	if path == "" {
		f.Error("Failed to get file path")
	}

	fmt.Println(path)
}
