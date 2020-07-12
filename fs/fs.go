package fs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func AppendFile(fname string, data []byte) error {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return f.Close()
}

func ReadLastBytes(fname string, buffer []byte) error {
	l := len(buffer)
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	_, err = f.Seek(int64(-l), 2)
	if err != nil {
		return err
	}
	_, err = f.Read(buffer)
	if err != nil {
		return err
	}
	return f.Close()
}

func ReadJSON(fname string, resObj interface{}) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &resObj)
}
