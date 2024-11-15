package bitget

import (
	v2 "github.com/yasseldg/bitget/pkg/client/v2"
)

type InterRest interface {
	InterRestPublic

	GetApiKeyList(subAccountUid string) (*v2.GetApiKeyList, error)

	GetAccountInfo() (*v2.GetAccountInfo, error)
	GetAccountList(productType string) (*v2.GetAccountList, error)
	GetAllAccountBalance() (*v2.GetAllAccountBalance, error)

	ChangeMarginMode(productType, symbol, marginCoin, marginMode string) (*v2.ChangeMarginMode, error)
	ChangePositionMode(productType, posMode string) (*v2.ChangePositionMode, error)

	GetPendingOrders(productType string) (*v2.GetPendingOrders, error)
	PlaceOrder(productType, symbol, side, size, marginMode, marginCoin, orderType string) (*v2.PlaceOrder, error)
	ModifyOrder(productType, symbol, newClientOid string) (*v2.ModifyOrder, error)
	CancelOrder(productType, symbol string) (*v2.CancelOrder, error)
	ClosePositions(productType string) (*v2.ClosePositions, error)
}

// user

func (r *Rest) GetApiKeyList(subAccountUid string) (*v2.GetApiKeyList, error) {
	return v2.NewGetApiKeyList(r.c, subAccountUid)
}

// account

func (r *Rest) GetAccountInfo() (*v2.GetAccountInfo, error) {
	return v2.NewGetAccountInfo(r.c)
}

func (r *Rest) GetAccountList(productType string) (*v2.GetAccountList, error) {
	return v2.NewGetAccountList(r.c, productType)
}

func (r *Rest) GetAllAccountBalance() (*v2.GetAllAccountBalance, error) {
	return v2.NewGetAllAccountBalance(r.c)
}

func (r *Rest) ChangeMarginMode(productType, symbol, marginCoin, marginMode string) (*v2.ChangeMarginMode, error) {
	return v2.NewChangeMarginMode(r.c, productType, symbol, marginCoin, marginMode)
}

func (r *Rest) ChangePositionMode(productType, posMode string) (*v2.ChangePositionMode, error) {
	return v2.NewChangePositionMode(r.c, productType, posMode)
}

// order
func (r *Rest) GetPendingOrders(productType string) (*v2.GetPendingOrders, error) {
	return v2.NewGetPendingOrders(r.c, productType)
}

func (r *Rest) PlaceOrder(productType, symbol, side, size, marginMode, marginCoin, orderType string) (*v2.PlaceOrder, error) {
	return v2.NewPlaceOrder(r.c, productType, symbol, side, size, marginMode, marginCoin, orderType)
}

func (r *Rest) ModifyOrder(productType, symbol, newClientOid string) (*v2.ModifyOrder, error) {
	return v2.NewModifyOrder(r.c, productType, symbol, newClientOid)
}

func (r *Rest) CancelOrder(productType, symbol string) (*v2.CancelOrder, error) {
	return v2.NewCancelOrder(r.c, productType, symbol)
}

func (r *Rest) ClosePositions(productType string) (*v2.ClosePositions, error) {
	return v2.NewClosePositions(r.c, productType)
}
