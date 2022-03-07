package main

import (
	 "encoding/json"
	 "reflect"
	 "testing"
)

func Test_isPrimeNumber(t *testing.T) {
	 type args struct {
		  number int
	 }
	 tests := []struct {
		  name  string
		  args  args
		  wantO []int
	 }{
		  {
				name:  "",
				args:  args{
					 number: 5,
				},
				wantO: []int{2,3,5},
		  },
	 }
	 for _, tt := range tests {
		  t.Run(tt.name, func(t *testing.T) {
				if gotO := getPrimeNumber(tt.args.number); !reflect.DeepEqual(gotO, tt.wantO) {
					 t.Errorf("isPrimeNumber() = %v, want %v", gotO, tt.wantO)
				}

		  })
	 }
}

func TestProcessPrimeNumber(t *testing.T) {
	 type args struct {
		  data []byte
	 }

	 var ip = []int{7,8}
	 bIp , _ := json.Marshal(ip)

	 tests := []struct {
		  name string
		  args args
	 }{
		  {
				name: "",
				args: args{
					 data: bIp,
				},
		  },
	 }
	 for _, tt := range tests {
		  t.Run(tt.name, func(t *testing.T) {
				ProcessPrimeNumberImplement(tt.args.data)

		  })
	 }
}