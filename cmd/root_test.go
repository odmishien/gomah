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
	}{
		{"長いやつ", "gomah -m 123123123123 -s 123 -p 123 -w 111", "🀇 🀈 🀉 🀇 🀈 🀉 🀇 🀈 🀉 🀇 🀈 🀉 🀐 🀑 🀒 🀙 🀚 🀛 🀀 🀀 🀀 \n"},
		{"普通のやつ", "gomah -m 123 -s 123 -p 123 -w 111 -d 22", "🀇 🀈 🀉 🀐 🀑 🀒 🀙 🀚 🀛 🀀 🀀 🀀 🀅 🀅 \n"},
		{"短いやつ", "gomah -m 123", "🀇 🀈 🀉 \n"},
	}
	for _, tc := range tcs {
		buf := new(bytes.Buffer)
		cmd := NewRootCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(tc.command, " ")
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if tc.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", tc.want, get)
		}
	}
}
