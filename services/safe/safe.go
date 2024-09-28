package safe

import "sync"

var Safety ThreadSafety

type ThreadSafety struct {
	mutex sync.Mutex
}

func (receiver *ThreadSafety) Lock(x func() interface{}) interface{} {
	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()
	return x()
}
