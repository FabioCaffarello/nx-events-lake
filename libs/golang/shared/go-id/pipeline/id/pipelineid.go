package pipelineid

import "fmt"

type ID = string

func NewID(context, serviceStart, source string) ID {
	return ID(fmt.Sprintf("%s-%s-%s", context, serviceStart, source))
}

