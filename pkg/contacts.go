package pkg

import "fmt"

const Email = "support@example.com"

var support string

func SetSupport(s string) {
	support = s
}

func GetSupport() string {
	return fmt.Sprintf("%s : <%s>", support, Email)
}
