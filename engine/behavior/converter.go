package behavior

import (
	"encoding/xml"
	. "github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/model"
)

func FindCurrentTask(bytearries model.Bytearry, taskDefineKey string) FlowElement {
	process := ConverterBpmn(bytearries)
	flowElement := process.FlowMap[taskDefineKey]
	return flowElement
}
func GetBpmn(bytes model.Bytearry) Process {
	return ConverterBpmn(bytes)
}
func ConverterBpmn(bytes model.Bytearry) Process {
	process, err := GetProcess(bytes.Id)
	if err == nil {
		return process
	}
	process = Converter([]byte(bytes.Bytes))
	SetProcess(bytes.Id, process)
	return process
}

func Converter(bytes []byte) Process {
	define := new(Definitions)
	xml.Unmarshal(bytes, &define)
	processes := define.Process
	if processes != nil {
		for j, p := range processes {
			flowMap := make(map[string]FlowElement, 0)
			processes[j].FlowMap = flowMap
			processes[j].Name = p.Documentation
			start := p.StartEvent
			if start != nil {
				for i, sta := range start {
					//flowMap[sta.Id] = start[i]
					processes[j].FlowMap[sta.Id] = start[i]
					processes[j].InitialFlowElement = start[i]
				}
			}
			se := p.SequenceFlow
			if se != nil {
				for i, s := range se {
					processes[j].FlowMap[s.Id] = se[i]
				}
			}
			user := p.UserTask
			if user != nil {
				for i, u := range user {
					behavior := UserTaskActivityBehavior{UserTask: user[i]}
					user[i].SetBehavior(behavior)
					processes[j].FlowMap[u.Id] = user[i]
				}
			}
			end := p.EndEvent
			if end != nil {
				for i, e := range end {
					processes[j].FlowMap[e.Id] = end[i]
				}
			}
			exclusiveGateways := p.ExclusiveGateway
			if exclusiveGateways != nil {
				for i, exclusiveGateway := range exclusiveGateways {
					behavior := ExclusiveGatewayActivityBehavior{}
					exclusiveGateways[i].SetBehavior(behavior)
					processes[j].FlowMap[exclusiveGateway.Id] = exclusiveGateways[i]
				}
			}
			inclusiveGateways := p.InclusiveGateway
			if inclusiveGateways != nil {
				for i, inclusiveGateway := range inclusiveGateways {
					behavior := InclusiveGatewayActivityBehavior{}
					inclusiveGateways[i].SetBehavior(behavior)
					processes[j].FlowMap[inclusiveGateway.Id] = inclusiveGateways[i]
				}
			}
			flows := make([]FlowElement, len(flowMap))
			count := 0
			for _, v := range flowMap {
				flows[count] = v
				count++
			}
			processes[j].Flow = flows
		}
	}
	ConvertXMLToElement(define)
	return define.Process[0]
}

//设置元素的出入口
func ConvertXMLToElement(model *Definitions) {
	processes := model.Process
	if processes != nil {
		for j, p := range processes {
			flows := p.Flow
			for i := range flows {
				value, ok := flows[i].(SequenceFlow)
				if ok {
					SourceRef := value.SourceRef
					//上一个节点
					lastFlow := processes[j].FlowMap[SourceRef]
					if lastFlow != nil {
						var outgoing = lastFlow.GetOutgoing()
						if outgoing == nil {
							outgoing = make([]*FlowElement, 0)
						}
						newOut := append(outgoing, &flows[i])
						//设置上一个节点出口
						lastFlow.SetOutgoing(newOut)
						//设置当前连线入口
						flows[i].SetSourceFlowElement(&lastFlow)

					}
					//下一个节点
					TargetRef := value.TargetRef
					nextFlow := processes[j].FlowMap[TargetRef]
					if nextFlow != nil {
						incoming := nextFlow.GetIncoming()
						if incoming == nil {
							incoming = make([]*FlowElement, 0)
						}
						newIn := append(incoming, &flows[i])
						m := make([]*FlowElement, 1)
						m[0] = &nextFlow
						//设置当前连线出口
						flows[i].SetTargetFlowElement(&nextFlow)
						//设置下一个节点入口
						nextFlow.SetIncoming(newIn)
					}
				}
			}
		}
	}
}
