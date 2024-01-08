package pipelinemd5id

import (
	"fmt"
	"libs/golang/shared/go-id/md5"
)

type ID = string

func NewPipelineGraphID(context string, serviceStart string, source string, processingId string) ID {
	id := md5.NewMd5Hash(fmt.Sprintf("%s-%s-%s-%s", context, serviceStart, source, processingId))
	return ID(string(id))
}
