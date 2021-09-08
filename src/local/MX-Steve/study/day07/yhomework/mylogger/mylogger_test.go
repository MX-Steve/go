package mylogger

import "testing"

// 单元测试
func TestConstLevel(t *testing.T){
	t.Logf("%v %T\n",InfoLevel, InfoLevel)
	t.Logf("%v %T\n", FatalLevel,FatalLevel)
}
