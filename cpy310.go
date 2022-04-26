package python3

// #cgo CFLAGS: -I${SRCDIR}/Include/
// #cgo linux CFLAGS: -I${SRCDIR}/config/linux -pthread
// #cgo windows CFLAGS: -I${SRCDIR}/include/windows/
// #cgo LDFLAGS: -L${SRCDIR}/libs/linux/ -l:libpython3.10.a -lcrypt -lpthread -ldl  -lutil -lm
/*
#include <Python.h>
*/
import "C"

func init() {
	Py_Initialize()
}
