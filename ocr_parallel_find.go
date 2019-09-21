package lookup

import (
	"sync"
)

// Search for all symbols in the image in parallel. Uses a Fan-out/fan-in approach.
func findAllInParallel(numWorkers int, symbols []*fontSymbol, bi *imageBinary, threshold float64) ([]*fontSymbolLookup, error) {
	f := &parallelFinder{
		img:        bi,
		threshold:  threshold,
		numWorkers: numWorkers,
		symbols:    symbols,
	}
	return f.lookupAll()
}

type parallelFinder struct {
	img        *imageBinary
	threshold  float64
	numWorkers int
	symbols    []*fontSymbol
}

type lookupResult struct {
	l   *fontSymbolLookup
	err error
}

func (f *parallelFinder) prepare(done <-chan struct{}) <-chan *fontSymbol {
	out := make(chan *fontSymbol)
	go func() {
		defer close(out)
		for _, s := range f.symbols {
			select {
			case out <- s:
			case <-done:
				return
			}
		}
	}()
	return out
}

func (f *parallelFinder) addWorker(done <-chan struct{}, in <-chan *fontSymbol) <-chan lookupResult {
	out := make(chan lookupResult)
	go func() {
		defer close(out)
		for symbol := range in {
			pp, err := lookupAll(f.img, symbol.image, f.threshold)
			if err != nil {
				out <- lookupResult{nil, err}
				continue
			}
			if pp != nil {
				for _, p := range pp {
					l := newFontSymbolLookup(symbol, p.X, p.Y, p.G)
					select {
					case out <- lookupResult{l, nil}:
					case <-done:
						return
					}
				}
			}

		}
	}()
	return out
}

func (f *parallelFinder) merge(done chan struct{}, cs []<-chan lookupResult) <-chan lookupResult {
	var wg sync.WaitGroup
	out := make(chan lookupResult)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan lookupResult) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
			if n.err != nil {
				close(done)
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (f *parallelFinder) lookupAll() ([]*fontSymbolLookup, error) {
	done := make(chan struct{})
	in := f.prepare(done)

	var workerOutputs = make([]<-chan lookupResult, f.numWorkers)
	for w := 0; w < max(f.numWorkers, 1); w++ {
		workerOutputs[w] = f.addWorker(done, in)
	}

	var result []*fontSymbolLookup
	for r := range f.merge(done, workerOutputs) {
		if r.err != nil {
			return nil, r.err
		}
		result = append(result, r.l)
	}
	close(done)
	return result, nil
}
