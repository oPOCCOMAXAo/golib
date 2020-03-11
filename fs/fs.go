package fs

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func AppendFile(fname string, data []byte) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	_, err = f.Write(data)
	check(err)
	err = f.Close()
	check(err)
}

func ReadLastBytes(fname string, buffer []byte) {
	l := len(buffer)
	f, err := os.Open(fname)
	check(err)
	_, err = f.Seek(int64(-l), 2)
	check(err)
	_, err = f.Read(buffer)
	check(err)
	err = f.Close()
	check(err)
	return
}
