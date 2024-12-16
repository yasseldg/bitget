package request

type InterPages interface {
	PageNo(pageNo int64)
	PageSize(pageSize int64)
}

type InterTimes interface {
	StartTime(startTime int64)
	EndTime(endTime int64)
}
