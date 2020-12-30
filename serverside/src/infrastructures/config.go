package infrastructures

import (
	"fmt"
	"os"
)

// Config returns environmet variable speicfied
func Config(name string) string {
	return os.Getenv(name)
}

// BucketName returns GCS bucket name for corresponding environment
func BucketName() string {
	return fmt.Sprintf("%s-temperature-load", os.Getenv("stage"))
}
