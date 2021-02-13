package model

type BpmnModel struct {
	Processes []*Process
}

func (bpmnModel BpmnModel) GetMainProcess() []*Process {
	return bpmnModel.Processes
}

func (bpmnModel *BpmnModel) AddProcess(process *Process) {
	bpmnModel.Processes = append(bpmnModel.Processes, process)
}
