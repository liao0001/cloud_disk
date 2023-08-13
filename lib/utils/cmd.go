package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecStd(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	fmt.Println("执行脚本: ", name, strings.Join(args, " "))
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return bs, nil
}
