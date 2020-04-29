package hw06_pipeline_execution //nolint:golint,stylecheck
import (
	"sync"
)

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var index int
	var mutex sync.Mutex

	outChan := make(Bi)
	tmpMap := make(map[int]I)
	wg := &sync.WaitGroup{}

	for v := range in {
		wg.Add(1)
		index++
		go func(inputValue I, inputItemIndex int) {
			defer wg.Done()
			for _, stage := range stages {
				select {
				case <-done:
					return
				default:
					tmpChan := make(Bi, 1)
					tmpChan <- inputValue
					inputValue = <-stage(tmpChan)
				}
			}
			mutex.Lock()
			tmpMap[inputItemIndex] = inputValue
			mutex.Unlock()
		}(v, index)
	}
	go func() {
		mapIndex := 1
		for {
			select {
			case <-done:
				close(outChan)
				return
			default:
				if mapIndex > index {
					close(outChan)
					return
				}
				mutex.Lock()
				value, ok := tmpMap[mapIndex]
				mutex.Unlock()
				if ok {
					outChan <- value
					mapIndex++
				}
			}
		}
	}()
	wg.Wait()
	return outChan
}
