package internal

import "sync"

type WorkerResult struct {
	IP string
	Port int
	Up bool
	Error error
}

func Worker(wg *sync.WaitGroup, input chan *DialSpec, result chan *WorkerResult) {
	defer wg.Done()
	for in := range input{
		up, err := Check(in)
		result<-&WorkerResult{
			IP:    in.Ip,
			Port:  in.Port,
			Up:    up,
			Error: err,
		}
	}
}
