package basic

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_ReadFile01(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/jun/Workspace/go_repos/src/github.com/jsgygujun/go_playground/data/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
