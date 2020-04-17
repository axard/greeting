package greet

import (
	"testing"
)

func TestGreetingFor(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Привествие без имени",
			args: args{
				name: "",
			},
			want: "Привет!",
		}, {
			name: "Приветствие с именем",
			args: args{
				name: "Василиса",
			},
			want: "Привет, Василиса!",
		}, {
			name: "Приветствие с именем и отчеством",
			args: args{
				name: "Василиса Ивановна",
			},
			want: "Привет, Василиса Ивановна!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreetingFor(tt.args.name); got != tt.want {
				t.Errorf("GreetingFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
