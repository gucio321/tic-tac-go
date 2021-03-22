package ttgcommon

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

// IntToCords converts intager to X-Y cords
func IntToCords(w, h, i int) (x, y int) {
	for {
		if i-w >= 0 {
			y++

			i -= w
		} else {
			x = i

			break
		}
	}

	return
}

// Clear clears console
func Clear() {
	var err error

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	}

	if err != nil {
		log.Print(err)
	}
}
