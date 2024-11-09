package clipfile

import (
	"fmt"
	"testing"

	"github.com/afeiship/go-clipfile"
)

func TestGetFilePath(f *testing.T) {
	path := clipfile.GetFilePath()

	if path == "" {
		f.Error("Failed to get file path")
	}

	fmt.Println(path)
}
