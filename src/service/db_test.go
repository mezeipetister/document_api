package service

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "hello", args: args{a: 5, b: 5}, want: 10},
		{name: "hello", args: args{a: 5, b: 5}, want: 10},
		{name: "hello", args: args{a: 5, b: 5}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
