package handler

import (
	"fmt"
	"github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

func init() {
	RegisterConstructor("userAuto", NewTestIActiviti)
}

func NewTestIActiviti(entity entity.ExecutionEntity) IActiviti {
	return &TestIActiviti{
		Entity: entity,
	}
}

type TestIActiviti struct {
	Entity entity.ExecutionEntity
	InPut  string
	OutPut string
}

func (test *TestIActiviti) GetInPut() interface{} {
	return test.InPut
}

func (test *TestIActiviti) GetOutPut() interface{} {
	return test.OutPut
}

func (test *TestIActiviti) User001() (code interface{}, err error) {
	variable := test.Entity.GetVariable()
	fmt.Println(variable)
	return contanst.ACTIVITI_HANDLER_CODE, nil
}

func (test *TestIActiviti) User002() (code interface{}, err error) {
	return contanst.ACTIVITI_HANDLER_CODE, nil
}
