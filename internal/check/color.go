package check

import (
	"crypto/md5"
	"fmt"
)

// Color - returns color hash for any string
func Color(str string) string {

	sum := md5.Sum([]byte(str))
	color := fmt.Sprintf("%x", sum)[0:6]

	return color
}
