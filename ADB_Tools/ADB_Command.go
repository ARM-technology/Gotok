package adbtools

import (
	_ "bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
	_ "time"
)

func ADB(args ...string) string {

	cmd := exec.Command("adb", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(output)

}

func OpenTikTok() {

	fmt.Println("📱 Open TikTok app")

	ADB(
		"shell",
		"monkey",
		"-p",
		"com.zhiliaoapp.musically",
		"-c",
		"android.intent.category.LAUNCHER",
		"1",
	)

	fmt.Println("Done send commend 👍...")
}

func SizeScreen() {
	size := ADB(
		"shell",
		"wm",
		"size",
	)

	fmt.Println(size)
}

func GetScreenSize() (int, int) {

	out := ADB("shell", "wm", "size")

	var w, h int
	fmt.Sscanf(out, "Physical size: %dx%d", &w, &h)

	return w, h
}

func Tap(ratioX, ratioY float64) {

	w, h := GetScreenSize()

	x := int(float64(w) * ratioX)
	y := int(float64(h) * ratioY)

	fmt.Println("Tapping at:", x, y)

	ADB("shell", "input", "tap",
		fmt.Sprint(x),
		fmt.Sprint(y),
	)
}

func PasteText(text string) {

	fmt.Println("⌨️ Writing text safely:", text)

	words := strings.Split(text, " ")
	for i, word := range words {
		if word != "" {
			ADB("shell", "input", "text", word)
			time.Sleep(50 * time.Millisecond)
		}
		if i < len(words)-1 {
			ADB("shell", "input", "keyevent", "62")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
