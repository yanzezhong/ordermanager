package errorcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCheckCode 此 ut 仅为了保证 errorcode 不重复，如果发现异常，请检查 errorcode
func TestCheckCode(t *testing.T) {
	file, err := os.Open("errorcode.go")
	assert.Nil(t, err)

	defer file.Close()

	codeMap := make(map[int]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, "=") {
			t1 := strings.Split(strings.Split(text, "=")[1], "//")[0]

			i, err := strconv.Atoi(strings.TrimSpace(t1))
			assert.Nil(t, err)
			if err != nil {
				fmt.Printf("strconv.Atoi :%v\n", text)
			}
			assert.False(t, codeMap[i])
			if codeMap[i] {
				fmt.Printf("errorcode is repeated:%v\n", i)
			}
			codeMap[i] = true
		}
	}
}
