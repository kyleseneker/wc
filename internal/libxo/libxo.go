package libxo

// #cgo CFLAGS: -I./converter
// #cgo LDFLAGS: -L/usr/local/Cellar/libxo/1.7.5/lib -lxo
// #include "converter.h"
// #include <stdlib.h>
import "C"
import "unsafe"

// ConvertToXML converts a string to XML using libxo
func ConvertToXML(data []byte) string {
	cData := C.CString(string(data))
	defer C.free(unsafe.Pointer(cData))

	C.initialize_xo()
	defer C.finalize_xo()

	var output *C.char
	C.convert_to_xml(cData)
	defer C.free(unsafe.Pointer(output))

	return C.GoString(output)
}
