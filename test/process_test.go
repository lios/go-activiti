package test

import (
	"fmt"
	"github.com/lios/go-activiti/engine/event"
	"github.com/lios/go-activiti/engine/impl/cfg"
	"github.com/lios/go-activiti/engine/impl/converter"
	_ "github.com/lios/go-activiti/engine/impl/handler"
	"github.com/lios/go-activiti/runtime"
	"io/ioutil"
	"os"
	"testing"
)

const (
	key         = "process_demo"
	userKey     = "userTest"
	userAutoKey = "userAuto"
	file_path   = "F:\\Work\\go-activiti\\resources\\userAuto.bpmn20.xml"
)

var processEngineConfiguration *cfg.ProcessEngineConfigurationImpl

type ActivitiListener struct {
	name string
}

func init() {
	processEngineConfiguration = cfg.GetProcessEngineConfiguration()
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
			repository := processEngineConfiguration.RepositoryService
			repository.Deploy(userAutoKey, userAutoKey, bytes)
		}
	}

}

//测试解析BPMN
func TestConverterBpmn(t *testing.T) {
	f, err := os.Open(file_path)
	if err == nil {
		bytes, err := ioutil.ReadAll(f)
		if err == nil {
			bpmnXMLConverter := converter.BpmnXMLConverter{}
			process := bpmnXMLConverter.ConvertToBpmnModel(bytes)
			fmt.Println(process)
		}
	}

}

//测试发起流程
func TestStartProcss(t *testing.T) {
	variables := make(map[string]interface{}, 0)
	variables["name"] = "lisi"
	variables["age"] = 18
	variables["isOld"] = true
	runtime := processEngineConfiguration.RuntimeService
	runtime.StartProcessInstanceByKey(userAutoKey, variables, "", "")
}

//测试发起流程
func TestStartAutoProcss(t *testing.T) {
	variables := make(map[string]interface{}, 0)
	variables["name"] = "lisi"
	variables["age"] = 18
	variables["isOld"] = true
	runtime := processEngineConfiguration.RuntimeService
	runtime.StartProcessInstanceByKey(userAutoKey, variables, "", "")
}

//测试查询代办
func TestQueryUndoTask(t *testing.T) {
	taskService := processEngineConfiguration.TaskService
	variables := make(map[string]interface{}, 0)
	variables["code"] = "0001"
	taskEntities, _ := taskService.QueryUndoTask("lisi", "")
	fmt.Println(taskEntities)
}

//测试完成任务
func TestComplete(t *testing.T) {
	taskService := processEngineConfiguration.TaskService
	variables := make(map[string]interface{}, 0)
	variables["code"] = "0001"
	taskService.Complete(132, variables, true)
}

//测试驳回
func TestBackTask(t *testing.T) {
	taskService := processEngineConfiguration.TaskService
	variables := make(map[string]interface{}, 0)
	variables["code"] = "0001"
	taskService.BackTask(26, "sid-4D05C44F-097D-4182-AD76-F4CC40F0F5F5")
}

//测试获取协程ID
func TestRuntime(t *testing.T) {
	id := runtime.GoroutineId()
	fmt.Println(id)
}

//测试监听器
func TestListener(t *testing.T) {
	eventListeners := make([]event.ActivitiEventListener, 0)
	eventListeners = append(eventListeners, ActivitiListener{})
	processEngineConfiguration.AddEventListeners(eventListeners)
	taskService := processEngineConfiguration.TaskService
	taskService.Complete(7, nil, false)
}
