package requester
//package main

import (
        "os"
        "bufio"
        "strings"
        "sync"
        //"io"
        //"log"
)

type FileReader struct {
    Filename string
    Fd *os.File
    Reader *bufio.Reader
    lock sync.Mutex
}

//func ReadLine() {
//
//    fd.Seek(offset, whence)
//    Reader.Reset(fd)
//}

func NewFileReader(fname string) *FileReader {
    return &FileReader {
        Filename : fname,
    }
}

func (fr *FileReader) Open() (err error) {
    fr.Fd, err = os.Open(fr.Filename)
    fr.Reader = bufio.NewReader(fr.Fd)
    return
}

func (fr *FileReader) ReadLine() (string, error) {
    fr.lock.Lock()
    line, err := fr.Reader.ReadString('\n')
    fr.lock.Unlock()

    line = strings.TrimSpace(line)
    if err != nil {
        fr.Reset(0, os.SEEK_SET)
        //if err == io.EOF {
        //    return line, nil
        //}
        return fr.ReadLine()
    }

    //log.Println(line)
    return line, nil
}

func (fr *FileReader) Reset(offset int64, whence int) (error) {
    fr.lock.Lock()
    fr.Fd.Seek(offset, whence)
    fr.Reader.Reset(fr.Fd)
    fr.lock.Unlock()
    return nil
}

//func main() {
//    file_reader := NewFileReader("test_data")
//    file_reader.Open()
//    for {
//        line, _ := file_reader.ReadLine()
//        log.Println(line)
//    }
//    return
//}
