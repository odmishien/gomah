package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRoot(t *testing.T) {
	tcs := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{"้ทใใใค", "gomah -m 123123123123 -s 123 -p 123 -w 111", "๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ \n", false},
		{"ๆฎ้ใฎใใค", "gomah -m 123 -s 123 -p 123 -w 111 -d 22", "๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ ๐ \n", false},
		{"็ญใใใค", "gomah -m 123", "๐ ๐ ๐ \n", false},
		{"ๅญๅจใใชใ็ใฎ็จฎ้กใฏ่ฝใกใ", "gomah -w 5", "", true},
	}
	for _, tc := range tcs {
		buf := new(bytes.Buffer)
		cmd := NewRootCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(tc.command, " ")
		cmd.SetArgs(cmdArgs[1:])
		err := cmd.Execute()
		if tc.wantErr && err == nil {
			t.Errorf("expected error, but not returned error")
		}

		get := buf.String()
		if !tc.wantErr && tc.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", tc.want, get)
		}
	}
}
