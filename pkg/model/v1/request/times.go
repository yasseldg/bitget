package request

import "github.com/yasseldg/go-simple/types/sInts"

type times struct {
	startTime *string
	endTime   *string
}

func (o *times) StartTime(startTime int64) {
	str := sInts.ToString(startTime)
	o.startTime = &str
}

func (o *times) EndTime(endTime int64) {
	str := sInts.ToString(endTime)
	o.endTime = &str
}

// Params returns the order parameters
func (o *times) GetParams(cp map[string]string) map[string]string {

	if o.startTime != nil {
		cp["startTime"] = *o.startTime
	}

	if o.endTime != nil {
		cp["endTime"] = *o.endTime
	}

	return cp
}
