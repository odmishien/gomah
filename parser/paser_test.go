package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	tcs := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"普通にうまく行く", "123m123s123p111w22d", "🀇 🀈 🀉 🀐 🀑 🀒 🀙 🀚 🀛 🀀 🀀 🀀 🀅 🀅 ", false},
		{"存在しないruleを指定すると落ちる", "123a123s123p111w22d", "", true},
		{"存在しない牌のindexを指定すると落ちる", "000m123s123p111w22d", "", true},
		{"順番をバラバラにしても大丈夫", "123s123m123p222d11w", "🀐 🀑 🀒 🀇 🀈 🀉 🀙 🀚 🀛 🀅 🀅 🀅 🀀 🀀 ", false},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.in)
			if tc.wantErr {
				if err == nil {
					t.Errorf("wantErr but didn't return any errors!")
				}
			}
			if got != tc.want {
				t.Errorf("got \"%s\", but want \"%s\"", got, tc.want)
			}
		})
	}
}
