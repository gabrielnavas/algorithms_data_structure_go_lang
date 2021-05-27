package main

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func NewDataTransferObjectFactory() *DataTransferObjectFactory {
	return &DataTransferObjectFactory{
		pool: make(map[string]DataTransferObject),
	}
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	dto := factory.pool[dtoType]
	if dto == nil {
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}
