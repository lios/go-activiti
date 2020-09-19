package test

import (
	"fmt"
	"github.com/lios/go-activiti/engine/behavior"
	peocess "github.com/lios/go-activiti/engine/impl"
	"github.com/lios/go-activiti/event"
	"github.com/lios/go-activiti/runtime"
	"io/ioutil"
	"os"
	"testing"
)

const (
	key       = "process_demo"
	file_path = "F:\\Work\\go-activiti\\resources\\process_demo.bpmn20.xml"
)

type ActivitiListener struct {
	name string
}

func (act ActivitiListener) OnEvent(event event.ActivitiEvent) error {
	fmt.Println(event)
	return nil
}

//测试部署流程
func TestDeployMentProcss(t *testing.T) {
	f, err := os.Open(file_path)
	if err == nil {
		bytes, err := ioutil.ReadAll(f)
		if err == nil {
			repository := peocess.RepositoryService{}
			repository.Deploy(key, key, bytes)
		}
	}

}

//测试发起流程
func TestStartProcss(t *testing.T) {
	variables := make(map[string]interface{}, 0)
	variables["name"] = "lisi"
	variables["age"] = 18
	variables["isOld"] = true
	runtime := peocess.RuntimeService{}
	runtime.StartProcessInstanceByKey(key, variables, "", "")
}

//测试完成任务
func TestComplete(t *testing.T) {
	taskService := peocess.TaskService{}
	variables := make(map[string]interface{}, 0)
	variables["code"] = "0001"
	taskService.Complete(35, variables, true)
}

//测试获取协程ID
func TestRuntime(t *testing.T) {
	id := runtime.GoroutineId()
	fmt.Println(id)
}

//测试监听器
func TestListener(t *testing.T) {
	configuration := behavior.GetProcessEngineConfiguration()
	eventListeners := make([]event.ActivitiEventListener, 0)
	eventListeners = append(eventListeners, ActivitiListener{})
	configuration.AddEventListeners(eventListeners)
	taskService := peocess.TaskService{}
	taskService.Complete(7, nil, true)
}
