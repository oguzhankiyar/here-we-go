package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/bamzi/jobrunner"
)

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", ReportJob{})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

type ReportJob struct {

}

func (j ReportJob) Run() {
	fmt.Printf("Every 5 sec create report\n")

	statusJson, _ := json.Marshal(jobrunner.StatusJson())

	fmt.Println("status:", string(statusJson))
}