package main

import "fmt"

// describe behaviors (reading, writing, running)
type reader interface {
	read(b []byte) (int, error)
}

func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title>Going GO Programming</title></channel></rss>"
	copy(b, s)
	return ln(s), nil
}

type file struct {
	name string
}

type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {
	//s := '{name: "bill", tittle: "developer"}'
	copy(b, s)
	return len(s), nil
}

// retrieve can read any device and process the data.
func retrieve(r reader) error {
	data := make([]byte, 100)
	len, err := r.reader(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data[:len]))
	return nil
}
func main() {
}
