package helpers

import "sync"

//паттер fan-in (концентратор) вокрер
// добавил сюда, чтобы не потерять хороший паттерн, может пригодиться

func Worker[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	wg.Add(len(channels))

	outChan := make(chan T)
	
	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for value := range ch {
				outChan <- value
			}
		}()
	}

	return outChan
}