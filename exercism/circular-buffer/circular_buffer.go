package circular

import "errors"

type Buffer struct {
	buffer   chan byte
	capacity int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{make(chan byte, size), size}
}

func (buffer *Buffer) ReadByte() (result byte, err error) {
	if len(buffer.buffer) == 0 {
		return result, errors.New("empty buffer")
	}
	return <-buffer.buffer, nil
}

func (buffer *Buffer) WriteByte(channel byte) error {
	if len(buffer.buffer) == cap(buffer.buffer) {
		return errors.New("buffer is full")
	}
	buffer.buffer <- channel
	return nil
}

func (buffer *Buffer) Overwrite(channel byte) {
	err := buffer.WriteByte(channel)
	if err != nil {
		<-buffer.buffer
		buffer.buffer <- channel
	}
}

func (buffer *Buffer) Reset() {
	buffer.buffer = make(chan byte, buffer.capacity)
}
