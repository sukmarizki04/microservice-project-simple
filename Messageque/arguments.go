package messageque

import (
	"fmt"
	"os"
)

func MainArgument() {
	var args = os.Args
	fmt.Printf("-> %#v\n", args)
	var args1 = args[1:]
	fmt.Printf("-> %#w\n", args1)

}
