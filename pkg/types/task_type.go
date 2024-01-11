package types

import "fmt"

type TaskType int //dao_sector sealing phase type
const (
	TaskType_Unknown TaskType = 0
	TaskType_AP      TaskType = 1
	TaskType_PC1     TaskType = 2
	TaskType_PC2     TaskType = 3
	TaskType_C1      TaskType = 4
	TaskType_C2      TaskType = 5
	TaskType_GET     TaskType = 6
	TaskType_FIN     TaskType = 7
)

func (t TaskType) GoString() string {
	return t.String()
}

func (t TaskType) String() string {
	switch t {
	case TaskType_AP:
		return "AP"
	case TaskType_PC1:
		return "PC1"
	case TaskType_PC2:
		return "PC2"
	case TaskType_C1:
		return "C1"
	case TaskType_C2:
		return "C2"
	case TaskType_GET:
		return "GET"
	case TaskType_FIN:
		return "FIN"
	}
	return fmt.Sprintf("Unknown<%d>", t)
}

func (t *TaskType) FromString(strTaskName string) {
	*t = MakeTaskTypeByName(strTaskName)
}

func (t TaskType) MarshalText() (value []byte, err error) {
	value = []byte(t.String())
	return value, nil
}

func (t *TaskType) UnmarshalText(value []byte) (err error) {
	t.FromString(string(value))
	return nil
}

func MakeTaskTypeByName(strTaskName string) TaskType {
	switch strTaskName {
	case "AP":
		return TaskType_AP
	case "PC1":
		return TaskType_PC1
	case "PC2":
		return TaskType_PC2
	case "C1":
		return TaskType_C1
	case "C2":
		return TaskType_C2
	case "GET":
		return TaskType_GET
	case "FIN":
		return TaskType_FIN
	}
	return TaskType_Unknown
}
