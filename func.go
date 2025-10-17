package go_mod_test_b

import "fmt"
import "github.com/moto1o/go_mod_test_a"


func PrintDepVersion() {
	fmt.Println("version: v1.0.1", "dep go_mod_test_a => ",go_mod_test_a.NameAndVersion())
}