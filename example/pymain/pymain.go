package main

import (
	"os"

	python "github.com/dzonerzy/cpy310"
)

func main() {
	python.Py_Initialize()
	python.Py_Main(os.Args)
	python.Py_Finalize()
}
