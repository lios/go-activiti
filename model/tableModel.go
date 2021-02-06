package model

type TableModel struct {
	AbstractModel
}

func (tableModel TableModel) GetTableName() string {
	return tableModel.getTableName()
}
