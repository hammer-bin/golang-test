package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

//func (f *File) Close() {
//
//}

func ReadFile(reader Reader) {
	c, ok := reader.(Closer)
	if ok {
		c.Close()
	}

	if c, ok := reader.(Closer); ok { //자주사용하는 구문
		c.Close()
	}

}

func main() {
	file := &File{}
	ReadFile(file)
}
