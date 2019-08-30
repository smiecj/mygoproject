package model

import "fmt"

type Student struct {
	Id int
	Name string
	Grade int
}

func (s Student) String() string {
	return fmt.Sprintf("学生基本信息: id: %d; 姓名: %s; 年级: %d\n", s.Id, s.Name, s.Grade)
}