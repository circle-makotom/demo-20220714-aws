package uname

import (
	"os/exec"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetUnameString(t *testing.T) {
	actual, err := GetUnameString()
	if err != nil {
		t.Error(err)
	}

	expect, err := runUnameCmd()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, actual, expect)
}

func runUnameCmd() (string, error) {
	cmd := exec.Command("uname", "-srvm")
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(output)), nil
	}
}
