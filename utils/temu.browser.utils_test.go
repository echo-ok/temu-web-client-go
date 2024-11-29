package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)
	
// GetAntiContent 获取反爬虫内容
func Test_GetAntiContent(t *testing.T) {
	result, err := GetAntiContent()
	fmt.Println(result)
	assert.Equal(t, nil, err, "Services.Mall.IsSemiManaged(ctx")
	assert.NotEqual(t, "", result, "result is empty")
}
