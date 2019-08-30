package test

import (
	"basic/model"
	"testing"
)

func TestImportModel(t *testing.T) {
	xiaoming := model.Student{1, "xiaoming", 1}
	//t.Log(xiaoming)

	// 测试从其他包中引入通过指针调用的方法
	// 实现了String方法的对象，可以自定义%v 格式的输出
	t.Logf("当前学生的信息为: %v\n", xiaoming)
}