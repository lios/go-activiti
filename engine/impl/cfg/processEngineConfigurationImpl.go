package cfg

import (
	. "github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/engine/event"
	. "github.com/lios/go-activiti/engine/impl/bpmn/parse"
	. "github.com/lios/go-activiti/engine/impl/bpmn/parse/deployer"
	. "github.com/lios/go-activiti/engine/impl/bpmn/parse/factory"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	. "github.com/lios/go-activiti/engine/impl/persistence/deploy"
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
	PostBpmnParseHandlers []BpmnParseHandler

	RuntimeService    RuntimeService
	RepositoryService RepositoryService
	TaskService       TaskServiceImpl

	Deployers                      []Deployer
	BpmnDeployer                   BpmnDeployer
	BpmnParser                     BpmnParser
	BpmnParseFactory               BpmnParseFactory
	ParsedDeploymentBuilderFactory ParsedDeploymentBuilderFactory
	//BpmnDeploymentHelper BpmnDeploymentHelper
	DeploymentManager       DeploymentManager
	ActivityBehaviorFactory ActivityBehaviorFactory
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
	initBpmnParser()
	initDeployers()
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
	commandInvoker := interceptor.GetCommandInvoker()
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
		commandExecutor := interceptor.CommandExecutorImpl{First: first}
		processEngineConfiguration.CommandExecutor = commandExecutor
		interceptor.SetCommandExecutorImpl(commandExecutor)
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

	processEngineConfiguration.RepositoryService = RepositoryServiceImpl{}

	processEngineConfiguration.TaskService = TaskServiceImpl{}
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
	//interceptor.SetProcessEngineConfiguration(processEngineConfiguration)
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
	SetAtivitiEventDispatcher(eventDispatcher)
	SetEventDispatcher(eventDispatcher)
}

func (processEngineConfiguration ProcessEngineConfigurationImpl) GetRuntimeService() RuntimeService {
	return processEngineConfiguration.RuntimeService
}
func initBpmnParser() {
	bpmnParser := BpmnParser{}
	bpmnParseFactory := DefaultBpmnParseFactory{}
	bpmnParser.BpmnParseFactory = bpmnParseFactory
	parseHandlers := BpmnParseHandlers{ParseHandlers: make(map[string][]BpmnParseHandler, 0)}
	parseHandlers.AddHandlers(getDefaultBpmnParseHandlers())
	bpmnParser.BpmnParserHandlers = parseHandlers

	activityBehaviorFactory := DefaultActivityBehaviorFactory{}
	bpmnParser.ActivityBehaviorFactory = activityBehaviorFactory
	processEngineConfiguration.ActivityBehaviorFactory = activityBehaviorFactory
	processEngineConfiguration.BpmnParser = bpmnParser
	processEngineConfiguration.BpmnParseFactory = bpmnParseFactory
}

func getDefaultBpmnParseHandlers() []BpmnParseHandler {
	handlers := make([]BpmnParseHandler, 0)
	handlers = append(handlers, ProcessParseHandler{AbstractActivityBpmnParseHandler: AbstractActivityBpmnParseHandler{AbstractBpmnParseHandler: AbstractBpmnParseHandler{ParseHandler: ParseHandler(ProcessParseHandler{})}}})
	handlers = append(handlers, StartEventParseHandler{AbstractActivityBpmnParseHandler: AbstractActivityBpmnParseHandler{AbstractBpmnParseHandler: AbstractBpmnParseHandler{ParseHandler: ParseHandler(StartEventParseHandler{})}}})
	handlers = append(handlers, UserTaskParseHandler{AbstractActivityBpmnParseHandler: AbstractActivityBpmnParseHandler{AbstractBpmnParseHandler: AbstractBpmnParseHandler{ParseHandler: ParseHandler(UserTaskParseHandler{})}}})
	handlers = append(handlers, SequenceFlowParseHandler{AbstractActivityBpmnParseHandler: AbstractActivityBpmnParseHandler{AbstractBpmnParseHandler: AbstractBpmnParseHandler{ParseHandler: ParseHandler(SequenceFlowParseHandler{})}}})

	return handlers
}
func initDeployers() {
	for _, deployer := range getDefaultDeployers() {
		processEngineConfiguration.Deployers = append(processEngineConfiguration.Deployers, deployer)
	}
	deploymentManager := DeploymentManager{}
	deploymentManager.Deployers = processEngineConfiguration.Deployers
	SetDeploymentManager(deploymentManager)
	processEngineConfiguration.DeploymentManager = deploymentManager
}

func getDefaultDeployers() []Deployer {
	defaultDeployers := make([]Deployer, 0)
	bpmnDeployer := BpmnDeployer{}
	initBpmnDeployerDependencies()
	bpmnDeployer.ParsedDeploymentBuilderFactory = processEngineConfiguration.ParsedDeploymentBuilderFactory
	processEngineConfiguration.BpmnDeployer = bpmnDeployer
	defaultDeployers = append(defaultDeployers, bpmnDeployer)
	return defaultDeployers
}
func initBpmnDeployerDependencies() {
	parsedDeploymentBuilderFactory := ParsedDeploymentBuilderFactory{}
	parsedDeploymentBuilderFactory.BpmnParser = &processEngineConfiguration.BpmnParser
	processEngineConfiguration.ParsedDeploymentBuilderFactory = parsedDeploymentBuilderFactory
}
