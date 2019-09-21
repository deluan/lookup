package lookup

import (
	"sync"
)

type lookupTask struct {
	img       *imageBinary
	fs        *fontSymbol
	threshold float64
}

type jobControl struct {
	wg      sync.WaitGroup
	tasks   chan *lookupTask
	results chan *fontSymbolLookup
}

func startJob() *jobControl {
	jc := &jobControl{
		tasks:   make(chan *lookupTask, 100),
		results: make(chan *fontSymbolLookup, 100),
	}
	for i := 0; i < 10; i++ {
		go jc.lookupSymbolWorker()
		jc.wg.Add(1)
	}
	return jc
}

func (jc *jobControl) lookupSymbolWorker() {
	for task := range jc.tasks {
		pp, _ := lookupAll(task.img, task.fs.image, task.threshold)
		if pp != nil {
			for _, p := range pp {
				l := newFontSymbolLookup(task.fs, p.X, p.Y, p.G)
				jc.results <- l
			}
		}
	}
	jc.wg.Done()
}

func (jc *jobControl) lookupSymbolParallel(img *imageBinary, fs *fontSymbol, threshold float64) {
	task := &lookupTask{
		img:       img,
		fs:        fs,
		threshold: threshold,
	}
	jc.tasks <- task
}

func (jc *jobControl) collectResults() ([]*fontSymbolLookup, error) {
	close(jc.tasks)

	done := make(chan bool)
	var results []*fontSymbolLookup
	go func() {
		for r := range jc.results {
			results = append(results, r)
		}
		done <- true
	}()

	jc.wg.Wait()
	close(jc.results)
	<-done
	return results, nil // TODO Implement error handling
}
