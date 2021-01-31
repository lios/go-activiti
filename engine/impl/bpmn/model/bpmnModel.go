package model

type BpmnModel struct {
	processes []Process
}

func (bpmnModel BpmnModel) GetMainProcess() []Process {
	return bpmnModel.processes
}
