package papilo

import "fmt"

// Pipe implements a buffered pipe.
// A buffered pipe connects two filter components.
//
// A pipe buffers incoming data,
// pushes the next data to the requiring component
// and receives from another pipe.
// A pipe can also propagate a close request which closes all pipes in front of it.
type Pipe struct {
	// bufSize is the max data the pipe can hold at a time
	bufSize int
	// isClosed is true only when a pipe has been closed
	// and is no more interested in data.
	isClosed bool
	// count is the number of data in the pipe at any time
	count int
	// buffer holds the pipe's data
	buffer []interface{}
}

func newPipe(bufSize int) Pipe {
	p := Pipe{
		bufSize: bufSize,
		buffer:  make([]interface{}, bufSize),
	}
	return p
}

// Next returns the next data in the buffer.
// An error is returned if there is no data in the buffer.
func (p *Pipe) Next() (interface{}, error) {
	if p.count == 0 {
		return nil, fmt.Errorf("No data in buffer")
	}
	p.count--
	return p.buffer[p.count], nil
}
