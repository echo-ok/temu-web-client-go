package normal

import (
	"fmt"
	"testing"
)

func TestTemuErrorCode(t *testing.T) {

	err := TemuErrorCode(6000001, stringPtr("需要短信验证码"))

	fmt.Println(err)

}
