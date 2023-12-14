package verbal

import (
	"fmt"
	"strings"
)

func Task1() {
	/*
		самым эффективным является использование strings.Join, оно основано на strings.Builder, который представляет из
		себя буфер для построения строк, то при использовании Join мы обращаемся на прямую к буферу и не создаем множество
		временных строк
	*/
	str := []string{"Hello", " ", "world", "!"}
	strForConcat := strings.Join(str, "")
	fmt.Println(strForConcat)
}
