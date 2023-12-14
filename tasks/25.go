package tasks

import (
	"fmt"
	"time"
)

func Task25() {
	fmt.Println("Hello")
	mySweetSleep(3)
	fmt.Println("Hello again after 3 second")
}

func mySweetSleep(seconds int) {
	// замеряем время перед началом цикла и завершаем его как только время становится больше
	start := time.Now()
	for int(time.Since(start).Seconds()) < seconds {
		continue
	}
}
