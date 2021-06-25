package hello

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tom", args{"Tom"}, "Hello, hello, Tom"},
		{"jerry", args{"Jerry"}, "Hello, hello, Jerry"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Sayhi() = %v, want %v", got, tt.want)
			}
		})
	}
}
