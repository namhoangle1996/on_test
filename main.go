package main

import (
	 "encoding/json"
	 "fmt"
	 "math"
	 "os"
)

type Event struct {
	 EventName EventName
	 Value     []byte
}

type EventName string

const (
	 EventReadFile           EventName = "EventReadFile"
	 EventProcessPrimeNumber EventName = "EventProcessPrimeNumber"
)

var eventReadFile = make(chan Event)
var eventProcessPrimeNumber = make(chan Event)

func main() {
	 foreverCh := make(chan bool)
	 go func() {
		  consumer()
	 }()

	 eventReadFile <- Event{
		  EventName: EventReadFile,
	 }

	 <-foreverCh
}

func consumer() {
	 for {
		  select {
		  case <-eventReadFile:
				ReadFileImplement()
		  case data := <-eventProcessPrimeNumber:
		  	 fmt.Println("getxxx",data)
				ProcessPrimeNumberImplement(data.Value)
		  default:
		  }
	 }

}

func ReadFileImplement() {
	 r, err := os.ReadFile("fileInput.txt")
	 if err != nil {
		  panic(err)
	 }

	 go func() {
		  eventProcessPrimeNumber <- Event{
				EventName: EventProcessPrimeNumber,
				Value:     r,
		  }
	 }()
}

func ProcessPrimeNumberImplement(data []byte) {
	 var intArr []int
	 err := json.Unmarshal(data, &intArr)
	 if err != nil {
		  panic(err)
	 }

	 var out = map[int][]int{}

	 for _, v := range intArr {
		  out[v] =  getPrimeNumber(v)
	 }


	 bOut , _ := json.Marshal(out)

	 os.WriteFile("fileOutput.txt", bOut , 0644)

}

func getPrimeNumber(number int)  (o []int ) {
	 if number > 1 {
		  for i := 2; i <= number; i++ {
				if isPrimeNumber(i) {
					 o = append(o, i)
				}
		  }
		  return o
	 }
	 return o
}

func isPrimeNumber(i int)  bool {
	 for j := 2; j <= int(math.Sqrt(float64(i))) ; j++ {
		  if i % j == 0 {
				fmt.Printf("not a prime number: %v\n", i)
				return false
		  }
	 }
	 return true
}
