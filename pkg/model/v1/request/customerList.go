package request

type CustomerList struct {
	times
	pages
}

// NewCustomerList returns a new CustomerList
func NewCustomerList() *CustomerList {
	return &CustomerList{}
}

// Params returns the order parameters
func (o *CustomerList) GetParams() map[string]string {
	cp := make(map[string]string)

	cp = o.times.GetParams(cp)

	return o.pages.GetParams(cp)
}
