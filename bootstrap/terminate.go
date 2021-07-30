package bootstrap

import "os"

func Cleanup() error {
	if FileExists(TempFolder()) {
		err := os.Remove(TempFolder())
		if err != nil {
			return err
		}
	}

	return nil
}

func Exit(exitCode int) {
	// Cleanup()
	os.Exit(exitCode)
}
