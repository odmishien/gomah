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
		{"長いやつ", "gomah -m 123123123123 -s 123 -p 123 -w 111", "🀇 🀈 🀉 🀇 🀈 🀉 🀇 🀈 🀉 🀇 🀈 🀉 🀐 🀑 🀒 🀙 🀚 🀛 🀀 🀀 🀀 \n", false},
		{"普通のやつ", "gomah -m 123 -s 123 -p 123 -w 111 -d 22", "🀇 🀈 🀉 🀐 🀑 🀒 🀙 🀚 🀛 🀀 🀀 🀀 🀅 🀅 \n", false},
		{"短いやつ", "gomah -m 123", "🀇 🀈 🀉 \n", false},
		{"存在しない牌の種類は落ちる", "gomah -w 5", "", true},
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
