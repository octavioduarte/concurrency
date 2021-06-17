package fakedownloadfile

import (
	"fmt"
	"sync"

	"time"

	"github.com/briandowns/spinner"
)

type filesData struct {
	name string
	size int
}

func PrintProgress(wg *sync.WaitGroup) {
	defer wg.Done()
	files := []filesData{
		{
			name: "report_sales_2015",
			size: 49,
		},
		{
			name: "report_sales_2016",
			size: 74,
		},
		{
			name: "report_sales 2017",
			size: 82,
		},
		{
			name: "report_sales 2017",
			size: 56,
		},
		{
			name: "report_sales 2017",
			size: 41,
		},
	}

	for _, f := range files {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		timeDownload := int((100 - f.size) / 10)
		s.Start()
		fmt.Println("[PROCCESS 2 - DOWNLOAD FILE] Making download of", f.name, "time expected: ", timeDownload, "seg")
		fmt.Printf("------------------------------------------------------------------- \n")
		time.Sleep(time.Duration(timeDownload) * time.Second)
		s.Stop()
	}
}
