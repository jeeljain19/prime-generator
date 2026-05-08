package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"prime-generator/internal/service"
)

func main() {
	fmt.Println("Prime Generator Starting...")

	//step 1
	start := flag.Int("start", 1, "start of range")   //stores the input
	end := flag.Int("end", 100, "end of range")
	algo := flag.String("algo", "brute", "algorithm to use")
	flag.Parse()
    

	//step 2
	//basic validation
	if *start > *end {
		fmt.Println("Invalid range: start must be less than or equal to end.")
		return
	}

	svc := service.NewPrimeService()
	result, err := svc.Generate(*start, *end, *algo)
	if err != nil {
		log.Fatal(err)
	}
    
	result2:=svc.GetStats()
	// Step 3: Print result
	// fmt.Println(result)
	// fmt.Println(result2)	


b, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(b))

b2, _ := json.MarshalIndent(result2, "", "  ")
fmt.Println(string(b2))




}