package infrastructures

import (
	"time"

	"github.com/TechDepa/c_tool/domain/model"
)

type Timer struct {
}

func NewTimer() Timer {
	return Timer{}
}

func (t Timer) Now() model.AppTime {
	return model.AppTime(time.Now())
}
