package config

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

//Bsc 合约相关配置
type Bsc struct {
	Contractaddress string `mapstructure:"contractaddress" json:"contractaddress" yaml:"contractaddress"` //薄饼地址
	Wss             string `mapstructure:"wss" json:"wss" yaml:"wss"`                                     // 节点连接
	ChainId         int64  `mapstructure:"chainId" json:"chainId" yaml:"chainId"`                         // 链ID
	GasLimit        uint64 `mapstructure:"gasLimit" json:"gasLimit" yaml:"gasLimit"`
	Apikey          string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`
	BscAddr         string `mapstructure:"bscAddr" json:"bscAddr" yaml:"bscAddr"`       // 币安接口地址
	Luk             string `mapstructure:"luk" json:"luk" yaml:"luk"`                   //LUK合约地址
	Tiger           string `mapstructure:"tiger" json:"tiger" yaml:"tiger"`             //NFT合约地址
	Usdt            string `mapstructure:"usdt" json:"usdt" yaml:"usdt"`                //USDT合约地址
	Liquidity       string `mapstructure:"liquidity" json:"liquidity" yaml:"liquidity"` // 流动性地址
}

//Wallet 钱包配置
type Wallet struct {
	Luk                string           `mapstructure:"luk" json:"luk" yaml:"luk"`                                              // 09-合约部署钱包（NFT， LUK）
	LukPrivateKey      WalletEncryption `mapstructure:"lukPrivateKey" json:"lukPrivateKey" yaml:"lukPrivateKey"`                // 09-合约部署钱包（NFT， LUK）私钥
	Complex            string           `mapstructure:"complex" json:"complex" yaml:"complex"`                                  // 01-20w综合钱包
	Contract           string           `mapstructure:"contract" json:"contract" yaml:"contract"`                               // 02-80w合约钱包
	Formalities        string           `mapstructure:"formalities" json:"formalities" yaml:"formalities"`                      // 03-DEX手续费钱包
	Nft                string           `mapstructure:"nft" json:"nft" yaml:"nft"`                                              // 04-NFT钱包
	Pledge             string           `mapstructure:"pledge" json:"pledge" yaml:"pledge"`                                     // 05-套餐质押钱包
	Destroy            string           `mapstructure:"destroy" json:"destroy" yaml:"destroy"`                                  // 06-销毁钱包
	DestroyPrivateKey  WalletEncryption `mapstructure:"destroyPrivateKey" json:"destroyPrivateKey" yaml:"destroyPrivateKey"`    // 09-合约部署钱包（NFT， LUK）私钥
	Reflow             string           `mapstructure:"reflow" json:"reflow" yaml:"reflow"`                                     // 07-回流底池钱包
	Withdraw           string           `mapstructure:"withdraw" json:"withdraw" yaml:"withdraw"`                               // 08-综合提现钱包
	WithdrawPrivateKey WalletEncryption `mapstructure:"withdrawPrivateKey" json:"withdrawPrivateKey" yaml:"withdrawPrivateKey"` // 08-综合提现钱包私钥
	Fee                string           `mapstructure:"fee" json:"fee" yaml:"fee"`                                              // 10-luk转账手续费钱包
}

type WalletEncryption string //自定义

func (t *WalletEncryption) len() int {
	return len(string(*t))
}

// 解密数据
func (t *WalletEncryption) String() string {

	data, err := base64.StdEncoding.DecodeString(string(*t))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	cryptText, err := goEncrypt.DesCbcDecrypt([]byte(data), []byte("asd12345"), []byte{})
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(cryptText)
}
