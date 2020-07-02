//
// $ go run progress_print.go
//

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const col = 30
	// Clear the screen by printing \x0c.
	// bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	// Position the Cursor:
    //    \033[<L>;<C>H
    //       or
    //    \033[<L>;<C>f
    //   puts the cursor at line L and column C.
	//bar := fmt.Sprintf("\033[6;3H")
	//bar := fmt.Sprintf("\033[0;0H[%%-%vs]", col)
	//bar := fmt.Sprintf("\033[F[%%-%vs]", col)
	bar := fmt.Sprintf("\r[%%-%vs]", col)
	
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!\n", strings.Repeat("=", col))
}
