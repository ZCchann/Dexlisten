# 常用信息

## BEP20节点信息

### 正式网节点

| 节点名称   | 主页                                                    | RPC URL                           | 请求限制/分钟 | 备注                   |
| ---------- | ------------------------------------------------------- | --------------------------------- | ------------- | ---------------------- |
| BSC主网    | https://docs.binance.org/smart-chain/developer/rpc.html | https://bsc-dataseed.binance.org/ | 60            | 超出1秒1次返回假数据   |
| getblock   | https://getblock.io/cn/pricing/                         |                                   | 30/s          | 付费无请求限制         |
| Chainstack | https://chainstack.com/                                 |                                   | 无            | 990美金支持BSC专用节点 |
| QuickNode  | https://www.quicknode.com/                              |                                   | 无            |                        |
| moralis    | https://moralis.io/                                     |                                   |               | 不支持txpool API       |
| ankr       | https://www.ankr.com/                                   |                                   |               | 免费版不支持txpool API |

### 测试网节点

| 节点名称  | 主页                                                    | RPC URL                                         | 请求限制/分钟 | 备注          |
| --------- | ------------------------------------------------------- | ----------------------------------------------- | ------------- | ------------- |
| BSC测试网 | https://docs.binance.org/smart-chain/developer/rpc.html | https://data-seed-prebsc-1-s1.binance.org:8545/ | 60            | 返回值3秒一次 |
| getblock  | https://getblock.io/cn/pricing/                         |                                                 | 27            |               |

## 合约

### 正式网

| 合约名称                | 合约地址                                   | 备注                | 交易手续费       |
| ----------------------- | ------------------------------------------ | ------------------- | ---------------- |
| WBNB                    | 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c | WBNB代币            |                  |
| BUSD                    | 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 | BUSD代币            |                  |
| BUSD-T Stablecoin       | 0x55d398326f99059fF775485246999027B3197955 | BEP20 USDT          |                  |
| USD Coin: USDC Token    | 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d | BEP20 USDC          |                  |
| PancakeSwap: Router v2  | 0x10ED43C718714eb63d5aA57B78B54704E256024E | PancakeSwap路由合约 | 千分之2.5(0.25%) |
| PancakeSwap: Factory v2 | 0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73 | PancakeSwap工厂合约 |                  |
| Biswap: Router          | 0x3a6d8cA21D1CF76F653A67577FA0D27453350dD8 | Biswap路由合约      | 千分之一(0.1%)   |
| Biswap: Factory         | 0x858E3312ed3A876947EA49d572A7C42DE08af7EE | Biswap工厂合约      |                  |
| ApeSwap: ApeRouter      | 0xcF0feBd3f17CEf5b47b0cD257aCf6025c5BFf3b7 | ApeSwap路由合约     | 千分之二(0.2%)   |
| ApeSwap: ApeFactory     | 0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6 | ApeSwap工厂合约     |                  |

### 测试网

| 合约名称                   | 合约地址                                   | LP合约地址                                 |
| -------------------------- | ------------------------------------------ | ------------------------------------------ |
| WBNB                       | 0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd |                                            |
| PancakeSwap: Router v2     | 0x9Ac64Cc6e4415144C455BD8E4837Fea55603e5c3 |                                            |
| PancakeSwap: Factory v2    | 0xB7926C0430Afb07AA7DEfDE6DA862aE0Bde767bc |                                            |
| FFF(测试代币)              | 0x76c50C192FcceDEf1D7324019172914EC9fF6C37 | 0x54E84AbAa3D9B1921E831086F27a4cF883b5F678 |
| BABYDAO（测试代币 带燃烧） | 0x5d73f006E54bbD64635F6c8b9e85858139894f7e |                                            |

## 常用网址

### 正式网

| 网址名称        | url                          | 备注 |
| --------------- | ---------------------------- | ---- |
| bsc区块链浏览器 | https://bscscan.com/         |      |
| pancake交易所   | https://pancakeswap.finance/ |      |
| Biswap          | https://biswap.org/          |      |
| ApeSwap         | https://apeswap.finance/     |      |

### 测试网

| 网址名称              | url                                          | 备注 |
| --------------------- | -------------------------------------------- | ---- |
| bsc测试网区块链浏览器 | https://testnet.bscscan.com/                 |      |
| bsc测试网水龙头       | https://testnet.binance.org/faucet-smart     |      |
| PancakeSwap测试网     | https://pancake.kiemtienonline360.com/#/swap |      |