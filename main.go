package main

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/jordan-lumley/a1pos/cmd/monitor"
	"github.com/jordan-lumley/a1pos/cmd/periphies"
	"github.com/jordan-lumley/a1pos/internal/logger"

	"github.com/jordan-lumley/service"
)

func main() {
	logger.Logger().Info("main() initialize")

	monitor.Execute()
	periphies.Execute()

	service.Run()
}

type document interface {
	writeTo(w io.Writer) (int64, error)
	read(r io.Reader) (int64, error)
	readAll(r io.Reader) (int64, error)
}

type documentStuff struct {
	document
	contents string
}

func (d *documentStuff) readAll(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err == io.EOF {
		panic(err)
	}

	return string(b), nil
}

func (d *documentStuff) read(r io.Reader) (string, error) {
	all := ""
	p := make([]byte, 18)

	for {
		n, err := r.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		all += string(p[:n])
	}

	println("done reading")

	return all, nil
}

func (d *documentStuff) writeTo(w io.Writer) (int64, error) {
	_, err := w.Write([]byte(d.contents))
	if err != nil {
		panic(err)
	}
	return 1, nil
}
