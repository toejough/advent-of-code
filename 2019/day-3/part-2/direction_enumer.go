// Code generated by "enumer -type=direction"; DO NOT EDIT.

//
package main

import (
	"fmt"
)

const _directionName = "LeftRightUpDown"

var _directionIndex = [...]uint8{0, 4, 9, 11, 15}

func (i direction) String() string {
	if i < 0 || i >= direction(len(_directionIndex)-1) {
		return fmt.Sprintf("direction(%d)", i)
	}
	return _directionName[_directionIndex[i]:_directionIndex[i+1]]
}

var _directionValues = []direction{0, 1, 2, 3}

var _directionNameToValueMap = map[string]direction{
	_directionName[0:4]:   0,
	_directionName[4:9]:   1,
	_directionName[9:11]:  2,
	_directionName[11:15]: 3,
}

// directionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func directionString(s string) (direction, error) {
	if val, ok := _directionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to direction values", s)
}

// directionValues returns all values of the enum
func directionValues() []direction {
	return _directionValues
}

// IsAdirection returns "true" if the value is listed in the enum definition. "false" otherwise
func (i direction) IsAdirection() bool {
	for _, v := range _directionValues {
		if i == v {
			return true
		}
	}
	return false
}
