package behavior

import (
	. "github.com/lios/go-activiti/engine/variable"
	. "github.com/lios/go-activiti/event"
)

var processEngineConfiguration ProcessEngineConfiguration

type ProcessEngineConfiguration struct {
	CommandInvoker        CommandInterceptor
	CommandInterceptors   []CommandInterceptor
	EventListeners        []ActivitiEventListener
	Service               ServiceImpl
	CommandExecutor       CommandExecutor
	CommandContextFactory CommandContextFactory
	VariableTypes         VariableTypes
	EventDispatcher       ActivitiEventDispatcher
}

func GetProcessEngineConfiguration() *ProcessEngineConfiguration {
	return &processEngineConfiguration
}
func init() {
	processEngineConfiguration = ProcessEngineConfiguration{}
	initCommandContextFactory()
	initCommandInvoker()
	initCommandInterceptors()
	initCommandExecutor()
	initService()
	initCommandContext(processEngineConfiguration)
	initVariableTypes()
	initEventDispatcher()
}

func initCommandContext(configuration ProcessEngineConfiguration) {
	//context := CommandContext{}
}

func (processEngineConfiguration ProcessEngineConfiguration) AddEventListeners(eventListeners []ActivitiEventListener) {
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

func getDefaultCommandInterceptors() []CommandInterceptor {
	var interceptors []CommandInterceptor
	interceptors = append(interceptors, &CommandContextInterceptor{CommandContextFactory: processEngineConfiguration.CommandContextFactory})
	//interceptors = append(interceptors, CommandInvoker{})
	interceptors = append(interceptors, &TransactionContextInterceptor{})
	return interceptors
}
func initCommandInvoker() {
	commandInvoker := CommandInvoker{}
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
func initService() {
	serviceImpl := ServiceImpl{CommandExecutor: processEngineConfiguration.CommandExecutor}
	SetServiceImpl(serviceImpl)
	processEngineConfiguration.Service = serviceImpl
}

func initInterceptorChain(interceptors []CommandInterceptor) CommandInterceptor {
	if len(interceptors) > 0 {
		for i := 0; i < len(interceptors)-1; i++ {
			interceptor := interceptors[i]
			interceptor.SetNext(interceptors[i+1])
		}
	}
	return interceptors[0]
}

func initCommandContextFactory() {
	factory := CommandContextFactory{}
	processEngineConfiguration.CommandContextFactory = factory
}

func initVariableTypes() {
	defaultVariableTypes := DefaultVariableTypes{}
	defaultVariableTypes.AddType(BooleanType{})
	defaultVariableTypes.AddType(IntType{})
	defaultVariableTypes.AddType(StringType{})
	defaultVariableTypes.AddType(MapType{})
	processEngineConfiguration.VariableTypes = defaultVariableTypes
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
