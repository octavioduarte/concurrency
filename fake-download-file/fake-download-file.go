package fakedownloadfile

import (
	"fmt"

	"time"

	"github.com/briandowns/spinner"
)

type filesData struct {
	name string
	size int
}

func PrintProgress() {
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
	}

	for _, f := range files {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Start()
		fmt.Println("[PROCCESS 2 - DOWNLOAD FILE] Making download of", f.name)
		time.Sleep(2 * time.Second)
		s.Stop()
	}
}
