package response

import "github.com/shopspring/decimal"

type PondData struct {
	Luk                decimal.Decimal `json:"luk"`                //池子LUK量
	Usdt               decimal.Decimal `json:"usdt"`               //池子USDT量
	DestructionLuk     decimal.Decimal `json:"destructionLuk"`     //销毁LUK量
	Address            int64           `json:"address"`            //全网用户数量
	SwapUsdt           decimal.Decimal `json:"swapUsdt"`           //买入U量
	SwapLuk            decimal.Decimal `json:"swapLuk"`            //买入L量
	SwapFee            decimal.Decimal `json:"swapFee"`            //手续费
	SellUsdt           decimal.Decimal `json:"sellUsdt"`           //买出U量
	SellLuk            decimal.Decimal `json:"sellLuk"`            //买出L量
	SellFee            decimal.Decimal `json:"sellFee"`            //买出手续费
	AddLiquidity       decimal.Decimal `json:"addLiquidity"`       //添加流动性总量
	RemoveLiquidity    decimal.Decimal `json:"removeLiquidity"`    //移除流动性总量
	RemoveLiquidityFee decimal.Decimal `json:"removeLiquidityFee"` //移除手续费
	Nft                int64           `json:"nft"`                //已售NFT个数
}

type PrecipitationData struct {
	Alliance  decimal.Decimal `json:"alliance"`  //8%联盟分红[套餐购买]
	StcPot    decimal.Decimal `json:"stcPot"`    //3%STC底池[套餐购买]
	StcNode   decimal.Decimal `json:"stcNode"`   //3%STC节点分红[套餐购买]
	Operate   decimal.Decimal `json:"operate"`   //3%运营钱包[套餐购买]
	Nft       decimal.Decimal `json:"nft"`       //1%NFT加权分红[套餐购买]
	Funding   decimal.Decimal `json:"funding"`   //1%联盟经费[套餐购买]
	Allowance decimal.Decimal `json:"allowance"` //1%联盟运营岗位津贴[套餐购买]
	LpPool    decimal.Decimal `json:"lpPool"`    //1%回流底池[Lp]
	LpOperate decimal.Decimal `json:"lpOperate"` //1%运营钱包[Lp]
}

type DividendsData struct {
	Recommend          decimal.Decimal `json:"recommend"`
	NftSwap            decimal.Decimal `json:"nftSwap"`
	NftSell            decimal.Decimal `json:"nftSell"`
	NftRemoveLiquidity decimal.Decimal `json:"nftRemoveLiquidity"`
	NftWeighted        decimal.Decimal `json:"nftWeighted"`
	NftValue           decimal.Decimal `json:"nftValue"`
	Combo7             decimal.Decimal `json:"combo7"`
	Combo15            decimal.Decimal `json:"combo15"`
	Combo30            decimal.Decimal `json:"combo30"`
	Combo60            decimal.Decimal `json:"combo60"`
	ComboRecommend     decimal.Decimal `json:"comboRecommend"`
	Linkage1           decimal.Decimal `json:"linkage1"`
	Linkage2           decimal.Decimal `json:"linkage2"`
	Linkage3           decimal.Decimal `json:"linkage3"`
	Linkage4           decimal.Decimal `json:"linkage4"`
	Community1         decimal.Decimal `json:"community1"`
	Community2         decimal.Decimal `json:"community2"`
	Community3         decimal.Decimal `json:"community3"`
	Community4         decimal.Decimal `json:"community4"`
	LpSwap             decimal.Decimal `json:"lpSwap"`
	LpSell             decimal.Decimal `json:"lpSell"`
	LpRemoveLiquidity  decimal.Decimal `json:"lpRemoveLiquidity"`
	LpRebate           decimal.Decimal `json:"lpRebate"`
}

type OtherData struct {
	Wss            string          `json:"wss"`
	LukToken       string          `json:"lukToken"`
	UsdtToken      string          `json:"usdtToken"`
	NftToken       string          `json:"nftToken"`
	LiquidityToken string          `json:"liquidityToken"`
	WithdrawUsdt   decimal.Decimal `json:"withdrawUsdt"`
	WithdrawLuk    decimal.Decimal `json:"withdrawLuk"`
}

type HeritageData struct {
	MetaverseNum          int64           `json:"metaverseNum"`          //新元宇宙人数
	MetaverseBalance      decimal.Decimal `json:"metaverseBalance"`      //新元宇宙释放总量USDT
	MetaverseRemaining    decimal.Decimal `json:"metaverseRemaining"`    //新元宇宙未释放总量USDT
	Metaverse             decimal.Decimal `json:"metaverse"`             //新元宇宙分红LUK
	MetaverseUsdt         decimal.Decimal `json:"metaverseUsdt"`         //新元宇宙分红USDT
	MetaverseOldNum       int64           `json:"metaverseOldNum"`       //老元宇宙人数
	MetaverseOldBalance   decimal.Decimal `json:"metaverseOldBalance"`   //老元宇宙释放总量USDT
	MetaverseOldRemaining decimal.Decimal `json:"metaverseOldRemaining"` //老元宇宙未释放总量USDT
	MetaverseOld          decimal.Decimal `json:"metaverseOld"`          //老元宇宙分红LUK
	MetaverseOldUsdt      decimal.Decimal `json:"metaverseOldUsdt"`      //老元宇宙分红USDT
	LpNum                 int64           `json:"lpNum"`                 //LP矿机人数
	LpBalance             decimal.Decimal `json:"lpBalance"`             //LP矿机释放总量USDT
	LpRemaining           decimal.Decimal `json:"lpRemaining"`           //LP矿机未释放总量USDT
	Lp                    decimal.Decimal `json:"lp"`                    //LP矿机分红LUK
	LpUsdt                decimal.Decimal `json:"lpUsdt"`                //LP矿机分红USDT
}
