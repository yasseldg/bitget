package model

type BookInfo struct {
	Asks     []interface{} `json:"asks"`
	Bids     []interface{} `json:"bids"`
	Checksum string        `json:"checksum"`
}

// @todo
func (bi *BookInfo) Merge(book BookInfo) BookInfo {
	var nb BookInfo

	nb.Asks = make([]interface{}, 0, len(bi.Asks)+len(book.Asks))
	nb.Asks = append(nb.Asks, bi.Asks...)
	nb.Asks = append(nb.Asks, book.Asks...)

	nb.Bids = make([]interface{}, 0, len(bi.Bids)+len(book.Bids))
	nb.Bids = append(nb.Asks, bi.Bids...)
	nb.Bids = append(nb.Asks, book.Bids...)

	return nb
}

func (bi *BookInfo) CheckSum(cs uint32) bool {
	// @todo
	return true
}
