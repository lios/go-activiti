package model

type BpmnModel struct {
	processes []*Process
}

func (bpmnModel BpmnModel) GetMainProcess() []*Process {
	return bpmnModel.processes
}

func (bpmnModel BpmnModel) AddProcess(process *Process) {
	bpmnModel.processes = append(bpmnModel.processes, process)
}
