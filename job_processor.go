package main

import ()

func ProcessJobs(jobs chan Job, db string) {
	for {
		j := <-jobs
		err := j.Run(db)

		j.ExitChan() <- err
	}
}
