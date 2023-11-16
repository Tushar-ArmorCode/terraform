// Code generated by "stringer -type=ComponentInstanceStatus component_instance.go"; DO NOT EDIT.

package hooks

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ComponentInstanceStatusInvalid-0]
	_ = x[ComponentInstancePending-46]
	_ = x[ComponentInstancePlanning-112]
	_ = x[ComponentInstancePlanned-80]
	_ = x[ComponentInstanceApplying-97]
	_ = x[ComponentInstanceApplied-65]
	_ = x[ComponentInstanceErrored-69]
}

const (
	_ComponentInstanceStatus_name_0 = "ComponentInstanceStatusInvalid"
	_ComponentInstanceStatus_name_1 = "ComponentInstancePending"
	_ComponentInstanceStatus_name_2 = "ComponentInstanceApplied"
	_ComponentInstanceStatus_name_3 = "ComponentInstanceErrored"
	_ComponentInstanceStatus_name_4 = "ComponentInstancePlanned"
	_ComponentInstanceStatus_name_5 = "ComponentInstanceApplying"
	_ComponentInstanceStatus_name_6 = "ComponentInstancePlanning"
)

func (i ComponentInstanceStatus) String() string {
	switch {
	case i == 0:
		return _ComponentInstanceStatus_name_0
	case i == 46:
		return _ComponentInstanceStatus_name_1
	case i == 65:
		return _ComponentInstanceStatus_name_2
	case i == 69:
		return _ComponentInstanceStatus_name_3
	case i == 80:
		return _ComponentInstanceStatus_name_4
	case i == 97:
		return _ComponentInstanceStatus_name_5
	case i == 112:
		return _ComponentInstanceStatus_name_6
	default:
		return "ComponentInstanceStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
