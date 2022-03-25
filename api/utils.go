package api

import "fmt"

func path(id string) string {
	return fmt.Sprintf("%s/%s", uploadDir, id)
}
