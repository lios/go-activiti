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

func (bpmnModel *BpmnModel) GetProcessById(id string) *Process {
	for _, process := range bpmnModel.Processes {
		if process.Id == id {
			return process
		}
	}
	return nil
}
