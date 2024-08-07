package errors

import "github.com/pkg/errors"

var (
	ErrInvalidProductType   = errors.New("Invalid Product type")
	ErrInvalidSubAccountUid = errors.New("Invalid SubAccountUid")
	ErrInvalidSymbol        = errors.New("Invalid Symbol")
	ErrInvalidNewClientOid  = errors.New("Invalid NewClientOid")
	ErrInvalidSide          = errors.New("Invalid Side")
	ErrInvalidSize          = errors.New("Invalid Size")
	ErrInvalidMarginMode    = errors.New("Invalid MarginMode")
	ErrInvalidMarginCoin    = errors.New("Invalid MarginCoin")
	ErrInvalidOrderType     = errors.New("Invalid OrderType")
	ErrInvalidPosMode       = errors.New("Invalid PosMode")

	ErrMissingOrderIdOrClientOid = errors.New("Missing OrderId or ClientOid")
)
