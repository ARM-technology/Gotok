package main

import (
	"fmt"
	"time"
	tools "toolsalbassam/ADB_Tools"
)

func main() {
	fmt.Println("Welcome for AlbassamTiktok Tools...👨‍💻")

	tools.SizeScreen()
	tools.OpenTikTok()
	time.Sleep(4 * time.Second)
	tools.Tap(0.07, 0.70)
	time.Sleep(2 * time.Second)
	tools.Tap(0.6465, 0.9375)
	time.Sleep(3 * time.Second)
	tools.PasteText("Hello From Go 👨‍💻...")
	time.Sleep(2 * time.Second)
	tools.Tap(0.1020, 0.5317)
	time.Sleep(2 * time.Second)
	tools.Tap(0.2354, 0.5048)
	time.Sleep(2 * time.Second)
	tools.Tap(0.5187, 0.1538)

}
