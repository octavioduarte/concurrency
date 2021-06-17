package main

import (
	"sync"

	"github.com/octavioduarte/concurrency/calculator"
	fakedownloadfile "github.com/octavioduarte/concurrency/fake-download-file"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go calculator.LoopCalcs(wg)
	go fakedownloadfile.PrintProgress(wg)
	wg.Wait()
}
