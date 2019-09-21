package lookup

import (
	"sync"
)

type lookupTask struct {
	imgBin      *imageBinary
	templateBin *imageBinary
	x, y        int
	m           float64
}

type jobControl struct {
	wg        sync.WaitGroup
	taskQueue chan *lookupTask
	results   chan *GPoint
}

func startJob() *jobControl {
	jc := &jobControl{
		taskQueue: make(chan *lookupTask, 100),
		results:   make(chan *GPoint, 100),
	}
	for i := 0; i < 10; i++ {
		go lookupWorker(i+1, jc)
		jc.wg.Add(1)
	}
	return jc
}

func lookupWorker(id int, jc *jobControl) {
	for task := range jc.taskQueue {
		g, _ := lookup(task.imgBin, task.templateBin, task.x, task.y, task.m)
		if g != nil {
			jc.results <- g
		}
	}
	jc.wg.Done()
}

func lookupParallel(jc *jobControl, img *imageBinary, template *imageBinary, x int, y int, m float64) {
	task := &lookupTask{
		imgBin:      img,
		templateBin: template,
		x:           x,
		y:           y,
		m:           m,
	}
	jc.taskQueue <- task
}

func collectResults(jc *jobControl) []GPoint {
	close(jc.taskQueue)

	done := make(chan bool)
	var results []GPoint
	go func() {
		for r := range jc.results {
			results = append(results, *r)
		}
		done <- true
	}()

	jc.wg.Wait()
	close(jc.results)
	<-done
	return results
}
