package db

import (
	"fmt"

	"github.com/jsnfwlr/vexil/internal/api/oapi"
)

func (t FlagType) ToAPIEnum() (flagType oapi.FlagType, fault error) {
	switch t {
	case FlagTypeBoolean:
		return oapi.Boolean, nil
	case FlagTypeString:
		return oapi.String, nil
	case FlagTypeInteger:
		return oapi.Integer, nil
	case FlagTypeJson:
		return oapi.Json, nil
	case FlagTypeStringArray:
		return oapi.StringSlice, nil
	case FlagTypeIntegerArray:
		return oapi.IntegerSlice, nil
	default:
		return oapi.FlagType("invalid"), fmt.Errorf("invalid flag type %q", t)
	}
}
