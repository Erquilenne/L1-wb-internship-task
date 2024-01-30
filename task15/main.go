package task15

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Данный фрагмент кода может привести к проблемам с памятью из-за создания огромной строки и присваивания её среза переменной justString

Использование bytes.Buffer в этом примере предпочтительно, потому что он может более эффективно управлять памятью при
выполнении множества операций записи и чтения, особенно с большими объемами данных. Это связано с тем,
что он динамически увеличивает свою емкость при необходимости, что может уменьшить издержки на реальлокацию памяти.
*/

import (
	"bytes"
	"fmt"
	"math/rand"
)

var justString string

func createHugeString(size int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, size)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(len(result))
	return string(result)
}

func someFunc() {
	fmt.Println(1 << 10)
	v := createHugeString(1 << 10)
	var buffer bytes.Buffer
	buffer.WriteString(v)
	justString = buffer.String()[:100]
}

func Run() {
	someFunc()
	fmt.Println(justString)
}
