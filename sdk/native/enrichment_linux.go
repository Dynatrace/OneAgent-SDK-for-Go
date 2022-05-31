package native

/*
#include <fcntl.h>

const char* const MAGIC_FILE_NAME = "dt_metadata_e617c525669e072eebe3d0f08212e8f2.properties";

// Wrap `open` because cgo doesn't understand variadic functions
static int myopen(const char *pathname, int flags) {
	return open(pathname, flags);
}
*/
import "C"

import (
	"bytes"
	"fmt"
	"os"
	"unsafe"
)

// On Linux, os.Open uses the `open` syscall, which OneAgent can't hook directly. Hence, we use a cgo
// implementation that calls the libc `open` function.
func openEnrichmentMetadataFile() (*os.File, error) {
	var f *os.File
	if fd, errno := C.myopen(C.MAGIC_FILE_NAME, C.O_RDONLY|C.O_CLOEXEC); fd < 0 {
		return nil, fmt.Errorf("failed to open magic file: %v", errno)
	} else {
		f = os.NewFile(uintptr(fd), C.GoString(C.MAGIC_FILE_NAME))
		defer func() { _ = f.Close() }()
	}

	buf := new(bytes.Buffer)
	if strlen, err := buf.ReadFrom(f); err != nil {
		return nil, err
	} else if b := buf.Bytes(); b[strlen-1] == '\n' {
		// Make sure the buffer is null-terminated for use as a C string
		b[strlen-1] = 0
	} else {
		buf.WriteByte(0)
	}

	cbuf := (*C.char)(unsafe.Pointer(&buf.Bytes()[0]))
	if fd, errno := C.myopen(cbuf, C.O_RDONLY|C.O_CLOEXEC); fd < 0 {
		return nil, fmt.Errorf("failed to open magic file: %v", errno)
	} else {
		buf.Truncate(buf.Len() - 1)
		return os.NewFile(uintptr(fd), buf.String()), nil
	}
}
