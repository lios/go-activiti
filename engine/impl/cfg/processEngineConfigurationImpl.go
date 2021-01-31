package cfg

import (
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
	"github.com/lios/go-activiti/engine/interceptor"
	. "github.com/lios/go-activiti/event"
)

var processEngineConfiguration ProcessEngineConfigurationImpl

type ProcessEngineConfigurationImpl struct {
	CommandInvoker        interceptor.CommandInterceptor
	CommandInterceptors   []interceptor.CommandInterceptor
	EventListeners        []ActivitiEventListener
	Service               ServiceImpl
	CommandExecutor       interceptor.CommandExecutor
	CommandContextFactory interceptor.CommandContextFactory
	EventDispatcher       ActivitiEventDispatcher
	PostBpmnParseHandlers []parse.BpmnParseHandler

	RuntimeService engine.RuntimeService
}

func GetProcessEngineConfiguration() *ProcessEngineConfigurationImpl {
	return &processEngineConfiguration
}
func init() {
	processEngineConfiguration = ProcessEngineConfigurationImpl{}
	initCommandContextFactory()
	initCommandInvoker()
	initCommandInterceptors()
	initCommandExecutor()
	initServices()
	initService()
	initCommandContext(processEngineConfiguration)
	initEventDispatcher()
}

func initCommandContext(configuration ProcessEngineConfigurationImpl) {
	//context := CommandContext{}
}

func (processEngineConfiguration ProcessEngineConfigurationImpl) AddEventListeners(eventListeners []ActivitiEventListener) {
	var EventListeners []ActivitiEventListener
	dispatcher := GetEventDispatcher()
	if len(eventListeners) > 0 {
		for _, listener := range eventListeners {
			EventListeners = append(EventListeners, listener)
			dispatcher.AddEventListener(listener)
		}
	}
	SetEventDispatcher(dispatcher)
	processEngineConfiguration.EventListeners = EventListeners
}

func getDefaultCommandInterceptors() []interceptor.CommandInterceptor {
	var interceptors []interceptor.CommandInterceptor
	interceptors = append(interceptors, &interceptor.CommandContextInterceptor{CommandContextFactory: processEngineConfiguration.CommandContextFactory})
	//interceptors = append(interceptors, CommandInvoker{})
	interceptors = append(interceptors, &interceptor.TransactionContextInterceptor{})
	return interceptors
}
func initCommandInvoker() {
	commandInvoker := interceptor.CommandInvoker{}
	processEngineConfiguration.CommandInvoker = &commandInvoker
}

func initCommandInterceptors() {
	interceptors := getDefaultCommandInterceptors()
	interceptors = append(interceptors, processEngineConfiguration.CommandInvoker)
	processEngineConfiguration.CommandInterceptors = interceptors
}

func initCommandExecutor() {
	if processEngineConfiguration.CommandExecutor == nil {
		first := initInterceptorChain(processEngineConfiguration.CommandInterceptors)
		commandExecutor := CommandExecutorImpl{First: first}
		processEngineConfiguration.CommandExecutor = commandExecutor
	}
}

func initServices() {
	processEngineConfiguration.RuntimeService = RuntimeServiceImpl{}
}
func initService() {
	serviceImpl := ServiceImpl{}
	serviceImpl.SetCommandExecutor(processEngineConfiguration.CommandExecutor)
	SetServiceImpl(serviceImpl)

	processEngineConfiguration.Service = serviceImpl

	processEngineConfiguration.RuntimeService = RuntimeServiceImpl{}
}

func initInterceptorChain(interceptors []interceptor.CommandInterceptor) interceptor.CommandInterceptor {
	if len(interceptors) > 0 {
		for i := 0; i < len(interceptors)-1; i++ {
			interceptor := interceptors[i]
			interceptor.SetNext(interceptors[i+1])
		}
	}
	return interceptors[0]
}

func initCommandContextFactory() {
	factory := interceptor.CommandContextFactory{}
	processEngineConfiguration.CommandContextFactory = factory
	interceptor.SetProcessEngineConfiguration(processEngineConfiguration)
}

func initEventDispatcher() {
	eventDispatcher := processEngineConfiguration.EventDispatcher
	if processEngineConfiguration.EventDispatcher == nil {
		eventDispatcher = ActivitiEventDispatcherImpl{EventSupport: &ActivitiEventSupport{}, Enabled: true}
	}
	if processEngineConfiguration.EventListeners != nil && len(processEngineConfiguration.EventListeners) > 0 {
		for _, listenerToAdd := range processEngineConfiguration.EventListeners {
			eventDispatcher.AddEventListener(listenerToAdd)
		}
	}
	processEngineConfiguration.EventDispatcher = eventDispatcher
	SetEventDispatcher(eventDispatcher)
}

func (processEngineConfiguration ProcessEngineConfigurationImpl) GetRuntimeService() engine.RuntimeService {
	return processEngineConfiguration.RuntimeService
}
