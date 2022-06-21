package os

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

// CheckRootExecution checks if this app is running as root in an OS-independent way
func CheckRootExecution() bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Unable to get current user: %s", err)
	}

	switch runtime.GOOS {
	case "linux", "darwin":
		return currentUser.Uid == "0"

	case "windows":
		// see https://gist.github.com/jerblack/d0eb182cc5a1c1d92d92a4c4fcc416c6
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		if err != nil {
			fmt.Println("Not running as admin")
			return false
		}
		return true

	default:
		fmt.Println("Cannot get user info")
		return false
	}
}
