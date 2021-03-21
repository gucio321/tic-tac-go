package ttgcommon

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

// IntToCords converts intager to X-Y cords
func IntToCords(i int) (x, y int) {
	for {
		if i-BoardW >= 0 {
			y++

			i -= BoardW
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
