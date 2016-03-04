// Every Go source file starts with a package clause
// Executable commands must always use package main.
package main
// To use a package, the importing source file identifies it by its package path in the import clause
import (
	"fmt"
	"go_poc/stringutil"
)

func main() {
	// Then the package name (as distinct from path) is used to qualify items from the package in the importing source file
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG , olleH"))
}
