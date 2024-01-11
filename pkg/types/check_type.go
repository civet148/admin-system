package types

import "fmt"

type CheckType int

const (
	CheckType_UserName        CheckType = 0
	CheckType_UserPhoneNumber CheckType = 1
	CheckType_RoleName        CheckType = 2
	CheckType_PoolName        CheckType = 3
	CheckType_ClusterName     CheckType = 4
	CheckType_UserEmail       CheckType = 5
)

var checkTypes = map[CheckType]string{
	CheckType_UserName:        "CheckType_UserName",
	CheckType_UserPhoneNumber: "CheckType_UserPhoneNumber",
	CheckType_RoleName:        "CheckType_RoleName",
	CheckType_PoolName:        "CheckType_PoolName",
	CheckType_ClusterName:     "CheckType_ClusterName",
	CheckType_UserEmail:       "CheckType_UserEmail",
}

func (t CheckType) String() string {
	if strType, ok := checkTypes[t]; ok {
		return strType
	}
	return fmt.Sprintf("CheckType_Unknown<%d>", t)
}

func (t CheckType) GoString() string {
	return t.String()
}
