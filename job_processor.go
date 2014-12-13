package main

import (
	"encoding/json"
	"io/ioutil"
)

type Job interface {
	ExitChan() chan error
	Run(todos map[string]Todo) (map[string]Todo, error)
}

func ProcessJobs(jobs chan Job, db string) {
	for {
		j := <-jobs

		todos := make(map[string]Todo, 0)
		content, err := ioutil.ReadFile(db)
		if err == nil {
			if err = json.Unmarshal(content, &todos); err == nil {
				todosMod, err := j.Run(todos)

				if err == nil && todosMod != nil {
					b, err := json.Marshal(todosMod)
					if err == nil {
						err = ioutil.WriteFile(db, b, 0644)
					}
				}
			}
		}

		j.ExitChan() <- err
	}
}
