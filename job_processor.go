package main

import ()

type Job interface {
	ExitChan() chan error
	Run(db string) error
}

func ProcessJobs(jobs chan Job, db string) {
	for {
		j := <-jobs
		err := j.Run(db)

		j.ExitChan() <- err
	}
}
