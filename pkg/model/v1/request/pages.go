package request

import "github.com/yasseldg/go-simple/types/sInts"

type pages struct {
	pageNo   *string
	pageSize *string
	uid      *string
}

func (o *pages) PageNo(pageNo int64) {
	str := sInts.ToString(pageNo)
	o.pageNo = &str
}

func (o *pages) PageSize(pageSize int64) {
	str := sInts.ToString(pageSize)
	o.pageSize = &str
}

func (o *pages) Uid(uid string) {
	o.uid = &uid
}

// Params returns the order parameters
func (o *pages) GetParams(cp map[string]string) map[string]string {

	if o.pageNo != nil {
		cp["pageNo"] = *o.pageNo
	}

	if o.pageSize != nil {
		cp["pageSize"] = *o.pageSize
	}

	if o.uid != nil {
		cp["uid"] = *o.uid
	}

	return cp
}
