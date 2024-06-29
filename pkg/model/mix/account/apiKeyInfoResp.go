package account

import "github.com/yasseldg/bitget/pkg/model"

type ApiKeyInfoResp struct {
	model.BaseResp `json:",inline"`
	Data           ApiKeyInfo `json:"data"`
}

type ApiKeyInfo struct {
	AffiliateID int `json:"affiliateID"`
}
