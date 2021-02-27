## Go语言流程引擎go-activiti
项目传送门[go-activiti](https://github.com/lios/go-activiti)

参考[Activiti](https://github.com/Activiti/Activiti)实现，满足部分功能。项目还在完善中，欢迎activiti爱好者加入。
	

现有能力

 -  节点类型，支持用户审批节点、排他网关、包容网关、并行网关
 - 流程功能，支持流程部署、流程发起、流程审批，流程驳回
 - 支持历史数据回溯
 - 全局事务
 - 全局事件监听，现支持节点事件处理
 - 节点事件回调



BPMN文件解析，先使用的是JSON库，使用方便，但存在缺陷，不支持扩展后续完善件[process](https://github.com/lios/go-activiti/blob/master/engine/process.go)。



## 全局事务

参考activiti设计模式，依赖命令模式和责任链模式，使用gorm的事务能力，不需要关注事务。

```
defer db.ClearTXDB()
db.GORM_DB.Transaction(func(tx *gorm.DB) error {
	db.InitTXDB(tx)
	value, err = transactionContextInterceptor.Next.Execute(command)
	return err
 })
```

## 流程事件监听
需实现ActivitiEventListener，将实现类加入配置文件

```java
type ActivitiListener struct {
	name string
}

func (act ActivitiListener) OnEvent(event event.ActivitiEvent) error {
	fmt.Println(event)
	return nil
}
```

```go
	configuration := behavior.GetProcessEngineConfiguration()
	eventListeners := make([]event.ActivitiEventListener, 0)
	eventListeners = append(eventListeners, ActivitiListener{})
	configuration.AddEventListeners(eventListeners)
```

## 节点事件回调
需实现IActiviti，注册构造器，参考iActivitiDemo.go文件

```java
func init() {
	RegisterConstructor("userAuto", NewTestIActiviti)
}

func NewTestIActiviti(entity ExecutionEntity) IActiviti {
	return &TestIActiviti{
		Entity: entity,
	}
}
```

后续计划：
	

 -  支持更多节点类型
 - bpmn解析完善，可扩展
 - 流程能力支持：流程跳转，撤销等等
 - 日志打印（优先）
 - 数据库默认值处理
 - 项目结构调整

期待您的加入。

