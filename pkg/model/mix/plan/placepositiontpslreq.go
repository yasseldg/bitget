package plan

type PlacePositionTPSLReq struct {
	/**
	 * Currency pair
	 */
	Symbol string `json:"symbol"`
	/**
	 * Deposit currency
	 */
	MarginCoin string `json:"marginCoin"`
	/**
	 * Plan type
	 */
	PlanType string `json:"planType"`
	/**
	 * Trigger price
	 */
	TriggerPrice string `json:"triggerPrice"`
	TriggerType  string `json:"triggerType"`
	/**
	 * Is this position long or short
	 */
	HoldSide string `json:"holdSide"`
}
