package order

/**
 * @Author: bitget-sdk-team
 * @Date: 2022-09-30 10:46
 * @DES: cancel the order request
 */
type CancelAllOrderReq struct {
	ProductType string `json:"productType"`
	MarginCoin  string `json:"marginCoin"`
}
