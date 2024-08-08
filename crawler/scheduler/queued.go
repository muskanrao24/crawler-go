package scheduler

import "distributed-web-crawler/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {

	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeWorker = workerQueue[0]
				activeRequest = requestQueue[0]
			}
			select {
			case r := <-s.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-s.workerChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
