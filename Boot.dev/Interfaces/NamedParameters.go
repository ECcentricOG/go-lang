package main

type Copier interface {
    copy(sourceFile string, destinationFile string) (bytesCopied int)
}

type file struct {
    name string
    path string
    size int
}

func(f file) copy(sourceFile string, destinationFile string) (bytesCopied int) {
    return f.size
}

func main() {

}
