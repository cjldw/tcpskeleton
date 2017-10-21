package tcpskeleton

import "sync"

// ThrowErr if err not equal nil then panic it.
func ThrowErr(err error)  {
	if err != nil {
		panic(err)
	}
}

// asyncDo asynchronous anonymous function
func asyncDo(fn func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		fn()
		wg.Done()
	}()
}
