package handler

import (
	"fmt"
	. "github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/common"
	"github.com/lios/go-activiti/errs"
)

func init() {
	RegisterConstructor("userAuto", NewTestIActiviti)
}

func NewTestIActiviti(entity ExecutionEntity) IActiviti {
	return &TestIActiviti{
		Entity: entity,
	}

}

type TestIActiviti struct {
	Entity ExecutionEntity
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
	return common.ACTIVITI_HANDLER_CODE, nil
}

func (test *TestIActiviti) User002() (code interface{}, err error) {
	return "", errs.ProcessError{"1007", "err"}
}
