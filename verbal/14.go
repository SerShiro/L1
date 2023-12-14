package verbal

import "fmt"

func Task14() {
	// создаем слайс с 2 элементами
	slice := []string{"a", "a"}
	// в функции локально создаем новый слайс, аналогично предыдущему примеру, все изменения будут только внутри функции
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	// то есть на экран выведется [b b a][a a]
	fmt.Print(slice)

}
