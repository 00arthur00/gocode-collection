import "sync"

//https://github.com/nsqio/nsq/blob/ef47a839594118beb40e9d2fb7956371a7768caa/internal/util/wait_group_wrapper.go
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
