package test

import (
	"basic/model"
	"testing"
)

func TestImportModel(t *testing.T) {
	xiaoming := model.Student{1, "xiaoming", 1}
	t.Log(xiaoming)
}