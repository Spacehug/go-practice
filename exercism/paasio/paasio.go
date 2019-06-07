package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader io.Reader
	counter
}

type writeCounter struct {
	writer io.Writer
	counter
}

type rwCounter struct {
	WriteCounter
	ReadCounter
}

type counter struct {
	bytes int64
	ops   int
	sync.RWMutex
}

func newCounter() counter {
	return counter{}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		writer:  w,
		counter: newCounter(),
	}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		reader:  r,
		counter: newCounter(),
	}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.reader.Read(p)
	rc.addBytes(m)
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, ops int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.writer.Write(p)
	wc.addBytes(m)
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, ops int) {
	return wc.count()
}

func (c *counter) addBytes(n int) {
	c.Lock()
	defer c.Unlock()
	c.bytes += int64(n)
	c.ops++
}

func (c *counter) count() (int64, int) {
	c.RLock()
	defer c.RUnlock()
	return c.bytes, c.ops
}
