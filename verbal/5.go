package verbal

import (
	"fmt"
	"unsafe"
)

func Task5() {
	// создаем пустую структуру, размер структуры определяется суммой размеров полей, но структура пустая и полей у нее нет
	var emptyStruct struct{}
	fmt.Println("Размер структуры:", unsafe.Sizeof(emptyStruct), "байт")
}
