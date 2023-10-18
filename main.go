package main

import (
	"encoding/json"
	"log"
	"sync"
)

type jsonRow struct {
	Dummy1   string `json:"dummy1"`
	Dummy2   string `json:"dummy2"`
	Valuable string `json:"valuable"`
}

var (
	file1Contents = []byte(`[{"dummy1":"1_dummy1","dummy2":"1_dummy2","valuable":"1_valuable"},{"dummy1":"2_dummy1","dummy2":"2_dummy2","valuable":"2_valuable"}]`)
	file2Contents = []byte(`[{"dummy1":"3_dummy1","dummy2":"3_dummy2","valuable":"3_valuable"},{"dummy1":"4_dummy1","dummy2":"4_dummy2","valuable":"4_valuable"}]`)
)

type processor func(row jsonRow)

func main() {
	// Some processor function, for the sake of the test task, we will just print the result
	printer := processor(func(row jsonRow) {
		// Print just the field we are interested in, omit the rest.
		log.Print(row.Valuable)
	})

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		processFile(file1Contents, printer)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		processFile(file2Contents, printer)
		wg.Done()
	}()

	wg.Wait()

	log.Print("all done")
}

func processFile(file []byte, fn processor) {
	var contents []jsonRow

	// TODO if files can get big - use json.Decoder. I will use json.Unmarshal for simplicity
	if err := json.Unmarshal(file, &contents); err != nil {
		// Handle the error in a simple way - just log and terminate
		log.Fatal(err)
	}

	for i := range contents {
		fn(contents[i])
	}
}
