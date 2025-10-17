package go_mod_test_b

import "fmt"
import "github.com/moto1o/go_mod_test_a"


func PrintDepVersion() {
	fmt.Println(go_mod_test_a.NameAndVersion())
}