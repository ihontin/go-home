package main

import (
	"net/http"
	"studentgit.kata.academy/Alkolex/go-kata/rewiew/rew2.1/mocks"
	"testing"
)

func TestWeatherServece_GetWather(t *testing.T) {
	type args struct {
		city string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "some test",
			args: args{city: "Moscow"},
			want: "Temperature...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			watherGet := mocks.NewWatherGetterer(t)
			watherGet.On("GetWather", tt.args.city).
				Return("Temperature...")
			c := &http.Client{}
			ws := &MyClient{
				client:        c,
				watherService: watherGet,
			}

			if got := ws.GetWather(tt.args.city); got != tt.want {
				t.Errorf("GetWather() = %v, want %v", got, tt.want)
			}
		})
	}
}
