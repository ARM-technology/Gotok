package FuncTools

import (
	"fmt"
	"time"
	tools "toolsalbassam/ADB_Tools"
)

func Comments_Tiktok(x int, textStr string) {
	fmt.Println("We Do loop: ", x)

	var counter int = 0
	for i := 0; i < x; i++ {
		var stop int = 5
		tools.SizeScreen()
		tools.OpenTikTok()
		time.Sleep(9 * time.Second)
		fmt.Println("=========>", i)

		for o := counter; o < counter+stop; o++ {
			fmt.Println("==========>", o)
			tools.Tap(0.07, 0.70)
			time.Sleep(4 * time.Second)
			tools.Tap(0.6465, 0.9219)
			time.Sleep(4 * time.Second)
			tools.PasteText(textStr)
			time.Sleep(4 * time.Second)
			tools.Tap(0.0965, 0.5317)

			time.Sleep(4 * time.Second)
			tools.Tap(0.5187, 0.1538)
			time.Sleep(4 * time.Second)
			tools.Swipe(0.5138, 0.5757, 0.4888, 0.2578)
			time.Sleep(4 * time.Second)

		}
		tools.CloseTikTok()

		time.Sleep(30 * time.Second)

	}

}
