package memscat

// #cgo CFLAGS: -g -Wall -Wno-implicit-function-declaration
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdint.h>
// #include <errno.h>
// #include <assert.h>
// #include <sys/uio.h>
//
// ssize_t
// memscat_read(pid_t pid, void * addr, void * out, size_t outsz)
// {
//     struct iovec local = {
//         .iov_base = out,
//         .iov_len  = outsz,
//     };
//
//     struct iovec remote = {
//         .iov_base = addr,
//         .iov_len  = outsz,
//     };
//
//     return process_vm_readv(pid, &local, 1, &remote, 1, 0);
// }
//
// ssize_t
// memscat_write(pid_t pid, void * addr, void * buf, size_t len)
// {
//     struct iovec local = {
//         .iov_base = buf,
//         .iov_len  = len,
//     };
//
//     struct iovec remote = {
//         .iov_base = addr,
//         .iov_len  = len,
//     };
//
//     return process_vm_writev(pid, &local, 1, &remote, 1, 0);
// }
import "C"
import "unsafe"

// Read will read `size` bytes from the process `pid` at address `addr`, and
// returns `size` bytes of that memory location data.
func Read(pid int, addr unsafe.Pointer, size int) []byte {
	buf := C.malloc(C.sizeof_char * C.size_t(size))
	defer C.free(unsafe.Pointer(buf))

	i := C.memscat_read(C.pid_t(pid),
		unsafe.Pointer(addr),
		unsafe.Pointer(buf),
		C.size_t(size))

	return C.GoBytes(buf, C.int(i))
}

// Write will write `size` bytes TO the process `pid` at the memory address of
// `addr`
func Write(pid int, addr unsafe.Pointer, buf unsafe.Pointer, size int) int {
	ret := C.memscat_write(C.pid_t(pid), addr, buf, C.size_t(size))

	return int(ret)
}
