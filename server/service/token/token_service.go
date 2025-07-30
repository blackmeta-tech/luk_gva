package token

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/model/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/model/token"
	tokenReq "github.com/flipped-aurora/gin-vue-admin/server/model/token/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/utils"
	"go.uber.org/zap"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	enum "github.com/flipped-aurora/gin-vue-admin/server/emun/luk"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	util "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	client    *ethclient.Client
	err       error
	rpcClient *rpc.Client // 为了 获取更详细信息
)

// 重写
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

type EventInfo struct {
	From   string          `json:"from" `   // 从（转帐地址）
	To     string          `json:"to" `     // 到（收款用户地址）
	Amount decimal.Decimal `json:"amount" ` // 转帐金额
	TxHash string          `json:"txHash" ` // 交易Hash
	Status int             `json:"status"`  // 确认状态 0 不成功 1成功
}

type AllBlockData struct {
	Status  *string    `json:"status"`
	Message *string    `json:"message"`
	Result  []BlockLog `json:"result"`
}

type BlockInfo struct {
	TxHash      string `json:"txHash" `      // 交易Hash
	BlockNumber string `json:"blockNumber" ` // 区块号
	From        string `json:"from" `        // msg.sender 消息发送者
	To          string `json:"to" `          // 合约地址或用户地址
}

type BlockData struct {
	BlockNumber       string `json:"blockNumber"`
	BlockHash         string `json:"blockHash"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	Input             string `json:"input"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Txreceipt_status  string `json:"txreceipt_status"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	IsError           string `json:"isError"`
}

type BlockLog struct {
	Address          string `json:"address"`
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash"`
	TransactionIndex string `json:"transactionIndex"`
	TransactionHash  string `json:"transactionHash"`
	GasPrice         string `json:"gasPrice"`
	GasUsed          string `json:"gasUsed"`
	TimeStamp        string `json:"timeStamp"`
}

var (
	// 代币合约 abi
	TransferEvent = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")) //0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	ApprovalEvent = crypto.Keccak256Hash([]byte("Approval(address,address,uint256)")) //0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92
	// PairABI 配对合约abi
	SwapEvent          = crypto.Keccak256Hash([]byte("Swap(address,uint256,uint256,uint256,uint256,address)")) //0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822
	MintEvent          = crypto.Keccak256Hash([]byte("Mint(address,uint256,uint256)"))                         // 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f
	BurnEvent          = crypto.Keccak256Hash([]byte("Burn(address,uint256,uint256,address)"))                 // 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496
	logTransferSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	// abi相关
	ContractCreateMethodId = "0x60806040" // 创建合约方法的方法Id,过滤掉

	LukContractABI *abi.ABI
	UniswapABI     *abi.ABI
)

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

type BscTokenService struct {
}

func NewToken(tokenEnum enum.TokenType) common.Address {
	var tokenAddress = ""
	switch tokenEnum {
	case enum.NFT:
		tokenAddress = global.GVA_CONFIG.Bsc.Tiger
	case enum.USDT:
		tokenAddress = global.GVA_CONFIG.Bsc.Usdt
	case enum.LUK:
		tokenAddress = global.GVA_CONFIG.Bsc.Luk
	default:
		tokenAddress = ""
	}
	if client == nil {
		client = NewWss()
	}
	// 创建调用合约实例
	tokenHexAddress := common.HexToAddress(tokenAddress)
	return tokenHexAddress
}

// 创建节点连接
func NewWss() *ethclient.Client {
	// 连接QuickNode 币安节点
	rpcClient, err = rpc.DialContext(context.Background(), global.GVA_CONFIG.Bsc.Wss)
	if err != nil {
		fmt.Println("连接失败")
	}

	client = ethclient.NewClient(rpcClient)
	return client
}

func GetTransactOpts(key string) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	adminAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), adminAddress)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	//接下来，我们创建一个新的keyed transactor，它接收私钥。
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(global.GVA_CONFIG.Bsc.ChainId))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                     // in wei
	auth.GasLimit = global.GVA_CONFIG.Bsc.GasLimit // in units
	auth.GasPrice = gasPrice
	return
}

// 重写 模块方法
func TransactionByHash(ctx context.Context, hash common.Hash) (rpc *rpcTransaction, isPending bool, err error) {

	var json *rpcTransaction
	err = rpcClient.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, err
	} else if json == nil {
		return nil, false, ethereum.NotFound
	} else if _, r, _ := json.tx.RawSignatureValues(); r == nil {
		return nil, false, fmt.Errorf("server returned transaction without signature")
	}

	return json, json.BlockNumber == nil, nil
}

//获取哈希状态
func (b *BscTokenService) QueryTxHash(txHashStr common.Hash) (receipt *rpcTransaction, status int) {
	var err error
	retryACount := 0
RetryA:
	receipts, err := client.TransactionReceipt(context.Background(), txHashStr)
	if err != nil && retryACount < 4 {
		retryACount++
		time.Sleep(time.Millisecond * 3000)
		goto RetryA
	}
	if receipts != nil {
		if receipts.Status == 0 && retryACount < 4 {
			retryACount++
			time.Sleep(time.Millisecond * 3000)
			goto RetryA
		}
		status = int(receipts.Status)
	}
	return
}

//直接获取哈希状态
func (b *BscTokenService) DirectQueryTxHash(txHashStr common.Hash) (amounts decimal.Decimal, address string, status int) {
	if client == nil {
		client = NewWss()
	}
	// 加载abi文件
	if LukContractABI == nil {
		if tokenAbi, err := abi.JSON(strings.NewReader(token.LukTokenABI)); err != nil {
			global.GVA_LOG.Error("加载abi文件失败", zap.Error(err))
		} else {
			LukContractABI = &tokenAbi
		}
	}
	tx, _, err := client.TransactionByHash(context.Background(), txHashStr)
	if err != nil {
		global.GVA_LOG.Error("哈希获取失败", zap.Error(err))
		return
	}
	input := tx.Data()
	// input前俩位'0x'去除掉
	_, inputsMap, err := b.DecodeTransactionInputData(LukContractABI, input)
	if err != nil {
		global.GVA_LOG.Error("解析MethodId失败", zap.Error(err))
		return
	}
	if a, ok := inputsMap["amount"]; ok {
		amount := *abi.ConvertType(a, new(*big.Int)).(**big.Int)
		amounts = utils.ToDecimal(amount)
	}
	chainId := new(big.Int).SetInt64(global.GVA_CONFIG.Bsc.ChainId)
	if msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), nil); err == nil {
		address = strings.ToLower(msg.From().Hex())
	}

	receipts, err := client.TransactionReceipt(context.Background(), txHashStr)
	if err != nil {
		return
	}
	if receipts != nil {
		status = int(receipts.Status)
	}
	return
}

//GetBalanceOf 查询余额
func (bscTokenService *BscTokenService) GetBalanceOf(transfer token.Transfer) (valueD decimal.Decimal, err error) {
	tokenHexAddress := NewToken(transfer.Type)
	address := common.HexToAddress(transfer.From)
	switch transfer.Type {
	case enum.LUK:
		valueD, err = (&LukTokenService{}).GetBalanceOf(tokenHexAddress, address)
	case enum.USDT:
		valueD, err = (&UsdtTokenService{}).GetBalanceOf(tokenHexAddress, address)
	default:
		err = errors.New("没有相关币种合约")
	}
	return
}

// 转账
// 管理员权限（部署账户）
func (bscTokenService *BscTokenService) TransferFrom(search tokenReq.TransferFrom) (txHash common.Hash, err error) {
	tokenHexAddress := NewToken(search.TokenType)
	amount := utils.ToWei(search.Amount)
	opts, err := GetTransactOpts(search.PrivateKey)
	if err != nil {
		return
	}

	switch search.TokenType {
	case enum.LUK:
		txHash, err = (&LukTokenService{}).TransferFrom(opts, tokenHexAddress, search.ToAddress, amount)
	case enum.USDT:
		txHash, err = (&UsdtTokenService{}).TransferFrom(opts, tokenHexAddress, search.ToAddress, amount)
	default:
		err = errors.New("没有相关币种合约")
	}
	return
}

// 获取最新区块高度
func (b *BscTokenService) GetBscTokenNewBlockHigth() (number uint64, err error) {
	if client == nil {
		client = NewWss()
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}
	number = uint64(header.Number.Int64())
	// fmt.Println(header.Number) // 5671744
	return
}

// 获取最初的区块高度
func (b *BscTokenService) GetBscTokenInitBlockHigth() (uint64, error) {
	var startBlock int64 = 0
	blockInfo, err := b.GetBscBlockData(startBlock)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return 0, err
	}
	for _, v := range blockInfo.Result {
		blockNumber, err := strconv.ParseUint(v.BlockNumber[2:], 16, 64)
		if err == nil {
			number := blockNumber
			return number, err
		}
	}
	return 0, err
}

// 解析input参数
func (b *BscTokenService) DecodeTransactionInputData(contractABI *abi.ABI, data []byte) (m *abi.Method, inputsMap map[string]interface{}, err error) {
	methodSigData := data[:4]
	inputsSigData := data[4:]
	method, err := contractABI.MethodById(methodSigData)
	m = method
	if err != nil {
		global.GVA_LOG.Error("ABI中没有这个方法")
		return
	}
	inputsMap = make(map[string]interface{})
	if err = method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		global.GVA_LOG.Error("解析input参数失败")
		return
	}
	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
	return
}

func (b *BscTokenService) GetBscBlockData(startBlock int64) (blockInfo *AllBlockData, err error) {
	params := url.Values{}
	params.Add("module", "logs")
	params.Add("action", "getLogs")
	params.Add("address", strings.ToLower(global.GVA_CONFIG.Bsc.Luk))
	params.Add("topic0", TransferEvent.String())
	params.Add("fromblock", util.ToString(startBlock))
	params.Add("page", "1")
	params.Add("sort", "asc")
	params.Add("apikey", global.GVA_CONFIG.Bsc.Apikey)

	resp, err := util.HttpRequestGetBsc(global.GVA_CONFIG.Bsc.BscAddr, "/api", params)
	if err != nil || resp.StatusCode != 200 {
		global.GVA_LOG.Sugar().Errorf("获取币安网数据失败, ststusCode:%d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	// var blockInfo = AllBlockData{}
	util.Unmarshal(resp.Body, &blockInfo)
	return
}

// 从区块高度区块扫描信息
func (b *BscTokenService) BlockScan(startBlock int64) (records []luk.LukBlockScanRecord, endNumber uint64, err error) {
	// 加载abi文件
	if UniswapABI == nil {
		if tokenAbi, err := abi.JSON(strings.NewReader(token.UniswapTokenABI)); err != nil {
			global.GVA_LOG.Error("加载abi文件失败", zap.Error(err))
		} else {
			UniswapABI = &tokenAbi
		}
	}
	// 获取区块信息
	blockInfo, err := b.GetBscBlockData(startBlock)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	txHash := make([]string, 0)
	// 解析区块数据
	//记录最后的一个区块
	if len(blockInfo.Result) > 0 {
		endNumber, err = strconv.ParseUint(blockInfo.Result[len(blockInfo.Result)-1].BlockNumber[2:], 16, 64)
		if err != nil {
			fmt.Println("=====", err)
		}
	}
	for _, v := range blockInfo.Result {
		//如果表里面有这个哈希就不需要执行后续操作
		if config.InStringArray(v.TransactionHash, txHash) {
			continue
		}
		txHash = append(txHash, v.TransactionHash)
		record := luk.LukBlockScanRecord{}
		_ = global.GVA_DB.Where("tx_hash = ?", v.TransactionHash).First(&record).Error
		if record.ID > 0 {
			global.GVA_LOG.Error("已存在再表数据里", zap.Error(err))
			continue
		}
		hash := common.HexToHash(v.TransactionHash)
		tx, _, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			global.GVA_LOG.Error("哈希获取失败", zap.Error(err))
			continue
		}
		input := tx.Data()
		// input前俩位'0x'去除掉
		method, inputsMap, err := b.DecodeTransactionInputData(UniswapABI, input)
		if err != nil {
			global.GVA_LOG.Error("解析MethodId失败", zap.Error(err))
			continue
		}
		block, err := strconv.ParseUint(v.BlockNumber[2:], 16, 64)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		record.Block = &block
		record.TxHash = v.TransactionHash
		//用哈希获取状态
		_, status := b.QueryTxHash(hash)
		record.Status = &status
		if *record.Status == 0 {
			global.GVA_LOG.Error("status is 0")
			continue
		}

		record.To = strings.ToLower(tx.To().String()) // 薄饼地址
		chainId := new(big.Int).SetInt64(global.GVA_CONFIG.Bsc.ChainId)
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), nil); err == nil {
			record.From = strings.ToLower(msg.From().Hex())
		}
		if record.To != strings.ToLower(global.GVA_CONFIG.Bsc.Contractaddress) {
			continue
		}
		if record.From == strings.ToLower(global.GVA_CONFIG.Wallet.Formalities) {
			continue
		}
		switch method.Name {
		// 添加流动性
		case string(enum.TokenMethodAddLiquidity):
			var lukAmount, usdtAmount *big.Int
			amountADesired := *abi.ConvertType(inputsMap["amountADesired"], new(*big.Int)).(**big.Int)
			amountBDesired := *abi.ConvertType(inputsMap["amountBDesired"], new(*big.Int)).(**big.Int)
			tokenA := *abi.ConvertType(inputsMap["tokenA"], new(common.Address)).(*common.Address)
			fmt.Println("====", amountADesired, amountBDesired)
			if strings.ToLower(tokenA.String()) == strings.ToLower(global.GVA_CONFIG.Bsc.Luk) {
				lukAmount = amountADesired
				usdtAmount = amountBDesired
			} else {
				lukAmount = amountBDesired
				usdtAmount = amountADesired
			}

			lpAmount, err := b.processAddLiquidity(hash, record.From)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				continue
			}
			record.Type = enum.MethodTypeAddLiquidity
			record.LukAmount = utils.ToDecimal(lukAmount)
			record.UsdtAmount = utils.ToDecimal(usdtAmount)
			record.LpAmount = lpAmount
			record.MethodName = enum.TokenMethodAddLiquidity
			//修改中心化数据
			(&luk.LukUserAddress{}).UpdateDataLp(record.From, record.LpAmount, true)
			break
		case string(enum.TokenMethodRemoveLiquidity), string(enum.TokenMethodRemoveLiquidityWithPermit):
			// 移除流动性
			lpAmount := *abi.ConvertType(inputsMap["liquidity"], new(*big.Int)).(**big.Int)
			record.LukAmount, record.UsdtAmount, record.LukServiceFee, record.UsdtServiceFee, err = b.processRemoveLiquidity(hash, record.From)
			fmt.Println("====", record.LukAmount, record.UsdtAmount, record.LukServiceFee, record.UsdtServiceFee)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				continue
			}
			record.LpAmount = utils.ToDecimal(lpAmount)
			record.MethodName = enum.TokenMethodRemoveLiquidity
			record.Type = enum.MethodTypeRemoveLiquidity
			//修改中心化数据
			(&luk.LukUserAddress{}).UpdateDataLp(record.From, record.LpAmount, false)
			break
		case string(enum.TokenMethodSwapTokensForExactTokens), string(enum.TokenMethodSwapExactTokensForTokens), string(enum.TokenMethodSwapExactTokensForTokensSupportingFeeOnTransferTokens):
			//兑换
			record.LukAmount, record.UsdtAmount, record.LukServiceFee, record.Type, err = b.processSwap(hash, record.From, inputsMap)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				continue
			}
		default:
			continue
		}
		tempTimeStamp, _ := strconv.ParseUint(v.TimeStamp[2:], 16, 64)
		timeStamp := int64(tempTimeStamp)
		record.MethodName = enum.TokenMethodName(method.Name)
		record.TimeStamp = &timeStamp
		record.Time = time.Unix(timeStamp, 0)
		records = append(records, record)
	}

	return
}

// 添加流动性
func (b *BscTokenService) processAddLiquidity(txHash common.Hash, userAddress string) (lp decimal.Decimal, err error) {
	tokenInfo, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		global.GVA_LOG.Sugar().Info(err.Error())
		return
	}
	for _, vLog := range tokenInfo.Logs {
		// 判断是否是流动性地址
		if !strings.EqualFold(vLog.Address.String(), global.GVA_CONFIG.Bsc.Liquidity) {
			continue
		}
		switch vLog.Topics[0].Hex() {
		case TransferEvent.Hex(): // 转账事件,LP的转移本质也是转账
			TokenFrom := strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex())
			TokenTo := strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex())
			// 从0地址转到用户地址代表添加了流动性
			if strings.EqualFold(TokenFrom, config.EMPTY_USERADDRESS) && strings.EqualFold(TokenTo, userAddress) {
				lp, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
				if err != nil {
					global.GVA_LOG.Error(err.Error())
					continue
				}
				return
			}
		}
	}
	return
}

// 移除流动性
func (b *BscTokenService) processRemoveLiquidity(txHash common.Hash, userAddress string) (lukAmount, usdtAmount, lukServiceFee, usdtServiceFee decimal.Decimal, err error) {
	records := luk.LukExchangeRecords{}
	info, _ := records.GetByTxHash(txHash.Hex())
	if info.ID > 0 {
		err = errors.New("txHash is existence")
		return
	}
	tokenInfo, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		global.GVA_LOG.Sugar().Info(err.Error())
		return
	}
	for _, vLog := range tokenInfo.Logs {
		TokenLuk := strings.ToLower(global.GVA_CONFIG.Bsc.Luk)
		TokenLiquidity := strings.ToLower(global.GVA_CONFIG.Bsc.Liquidity)
		TokenRemovelq := strings.ToLower(global.GVA_CONFIG.Wallet.Formalities)
		userAddress = strings.ToLower(userAddress)
		// 移除流动性-Luk
		if strings.EqualFold(strings.ToLower(vLog.Address.String()), TokenLuk) {
			switch vLog.Topics[0].Hex() {
			case TransferEvent.Hex():
				TokenFrom := strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex())
				TokenTo := strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex())
				// 从薄饼地址转到用户地址表示获得的Luk
				if strings.EqualFold(TokenFrom, TokenLiquidity) && strings.EqualFold(TokenTo, userAddress) {
					lukAmount, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				}
				// 从薄饼地址转到移除流动性钱包地址表示Luk手续费
				if strings.EqualFold(TokenFrom, TokenLiquidity) && strings.EqualFold(TokenTo, TokenRemovelq) {
					lukServiceFee, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				}
			}
		}
		// 移除流动性-USDT
		if strings.EqualFold(strings.ToLower(vLog.Address.String()), strings.ToLower(global.GVA_CONFIG.Bsc.Usdt)) {
			switch vLog.Topics[0].Hex() {
			case TransferEvent.Hex():
				TokenFrom := strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex())
				TokenTo := strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex())
				// 从合约地址转到用户地址表示获得的Usdt
				if strings.EqualFold(TokenFrom, TokenLiquidity) && strings.EqualFold(TokenTo, userAddress) {
					usdtAmount, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				}
				// 从合约地址转到移除流动性钱包地址表示usdt手续费
				if strings.EqualFold(TokenFrom, TokenLiquidity) && strings.EqualFold(TokenTo, TokenRemovelq) {
					usdtServiceFee, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				}
			}
		}
	}
	records.Type = enum.MethodTypeRemoveLiquidity
	records.Usdt = usdtAmount
	records.Luk = lukAmount
	records.ChargeLuk = lukServiceFee
	records.Address = userAddress
	records.TxHash = txHash.String()
	err = global.GVA_DB.Debug().Model(luk.LukExchangeRecords{}).Create(&records).Error
	return
}

// 兑换
func (b *BscTokenService) processSwap(txHash common.Hash, address string, inputsMap map[string]interface{}) (lukNum, usdtNum, chargeNum decimal.Decimal, _type enum.MethodType, err error) {
	records := luk.LukExchangeRecords{}
	info, _ := records.GetByTxHash(txHash.Hex())
	if info.ID > 0 {
		err = errors.New("txHash is existence")
		return
	}
	if tokens, ok := inputsMap["path"].([]common.Address); ok {
		tokenInfo, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			return decimal.Zero, decimal.Zero, decimal.Zero, 0, err
		}
		var deduction, obtain, charge decimal.Decimal
		for _, vLog := range tokenInfo.Logs {
			switch vLog.Topics[0].Hex() {
			case TransferEvent.Hex():
				TokenFrom := strings.ToLower(common.HexToAddress(vLog.Topics[1].Hex()).Hex())
				TokenTo := strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).Hex())

				if strings.EqualFold(TokenFrom, address) && strings.EqualFold(TokenTo, address) {
					continue
				}
				if strings.EqualFold(TokenTo, strings.ToLower(global.GVA_CONFIG.Wallet.Formalities)) {
					//手续费
					charge, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				} else if strings.EqualFold(TokenFrom, address) {
					//用户扣除的
					deduction, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				} else if strings.EqualFold(TokenTo, address) {
					//用户得到的
					obtain, err = (&BscTokenService{}).ParseTransferData(vLog.Data)
					if err != nil {
						global.GVA_LOG.Error(err.Error())
						continue
					}
				}
			}
		}
		if tokens[0] == common.HexToAddress(global.GVA_CONFIG.Bsc.Usdt) {
			records.Type = enum.MethodTypeSwap
			records.Usdt = deduction
			records.Luk = obtain
		} else {
			records.Type = enum.MethodTypeSell
			records.Usdt = obtain
			records.Luk = deduction.Add(charge).Add(decimal.NewFromFloat(0.00001))
		}
		records.ChargeLuk = charge
		records.Address = address
		records.TxHash = txHash.String()
		err = global.GVA_DB.Debug().Model(luk.LukExchangeRecords{}).Create(&records).Error
		lukNum = records.Luk
		usdtNum = records.Usdt
		chargeNum = records.ChargeLuk
		_type = records.Type
	} else {
		err = errors.New("tokens is no []common.Address")
	}
	return
}

func (d *BscTokenService) ParseTransferData(data []byte) (result decimal.Decimal, err error) {
	if LukContractABI == nil {
		if tokenAbi, err := abi.JSON(strings.NewReader(token.LukTokenABI)); err != nil {
			global.GVA_LOG.Error("加载abi文件失败", zap.Error(err))
		} else {
			LukContractABI = &tokenAbi
		}
	}
	transfer, err := LukContractABI.Unpack("Transfer", data)
	if err != nil {
		global.GVA_LOG.Error("BlockScan::Transfer::", zap.Error(err))
		return
	}
	if len(transfer) == 0 {
		err = fmt.Errorf("transfer长度为0")
		return
	}
	v1 := *abi.ConvertType(transfer[0], new(*big.Int)).(**big.Int)
	result = utils.ToDecimal(v1)
	return
}

//获取兑换后的金额
/*func GetSwapAmount(txHash common.Hash) (amount decimal.Decimal, err error) {
	tokenInfo, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return
	}
	if len(tokenInfo.Logs) == 0 {
		err = errors.New("Swap fail")
		return
	}
	for _, vLog := range tokenInfo.Logs {
		switch vLog.Topics[0].Hex() {
		case TransferEvent.Hex():
			TokenTo := common.HexToAddress(strings.ToLower(vLog.Topics[2].Hex()))
			if strings.EqualFold(TokenTo.Hex(), strings.ToLower(global.GVA_CONFIG.Wallet.Luk)) {
				amount, err = (&UsdtTokenService{}).ParseTransferDataUsdt(vLog.Data)
				return
			}
		}
	}
	return
}
*/
