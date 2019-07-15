# go-memscat

Very tiny golang wrapper around the process read/write scatter-gather system
api.


# example usage

Lets say we know the memory address of a buffer that contains the string
"XYZbar", and we want to change the first 3 bytes to be "foo" externally,
without ptrace or such, we can use this little library as follows:

```go
package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/nathanfrench/memscat"
)

func main() {
	i_str := []byte{'X', 'Y', 'Z', 'b', 'a', 'r'}
	o_str := []byte{'f', 'o', 'o'}

	fmt.Printf("The contents of i_str is \"%v\"\n", string(i_str))

	memscat.Write(syscall.Getpid(),
		unsafe.Pointer(&i_str[0]),
		unsafe.Pointer(&o_str[0]), 3)

	fmt.Printf("The contents of i_str should now be 'foobar': (i_str=\"%v\")\n", string(i_str))
}
```

```shell
➜ go run ~/test_memscat.go
The contents of i_str is "XYZbar"
The contents of i_str should now be 'foobar': (i_str="foobar")
```

