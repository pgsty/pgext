## 用法

来源：

- [官方 README](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/README.md)
- [版本与 PostgreSQL 特性矩阵](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/Cargo.toml)
- [官方 RPC 实现](https://github.com/uygunbodur/postgrechain/tree/dafe2307effd636a6954526cd95190779987af07/src/rpc)

`postgreChain` 0.0.0 将同步 Solana RPC 操作暴露为 PostgreSQL 函数。它可以读取余额和程序账户、创建钱包、请求测试网络空投以及提交转账；网络调用在数据库后端内执行，因此延迟、故障和秘密处理会直接影响数据库会话。

### 核心流程

```sql
CREATE EXTENSION "postgreChain";

SELECT pc_balance('WalletPublicKeyBase58', 'devnet');

SELECT public_key, data_len, lamports, rent_epoch, executable
FROM pc_get_program_accounts('ProgramPublicKeyBase58', 'testnet');
```

支持的网络字符串是 `mainnet`、`devnet`、`testnet` 与 `localhost`。固定提交中的实现会把任何其他值静默回退到 `devnet`，因此应用代码必须校验网络值。

### 主要函数

- `pc_balance(address, network)`：返回账户的 lamport 余额；地址或 RPC 出错时实现会返回零。
- `pc_token_account_balance(address, network)`：返回 SPL 代币账户余额，出错时同样返回零。
- `pc_get_program_accounts(program_id, network)`：返回账户公钥、数据长度、lamport、租金纪元和可执行标志。
- `pc_create_wallet()`：返回新生成的公钥与 Base58 编码私钥。
- `pc_airdrop(address, amount, network)`：请求 SOL，并以布尔值返回成功状态。
- `pc_transfer(public_key, secret_key, destination, amount, network)`：签名并提交 SOL 转账，以布尔值返回成功状态。

### 安全与事务边界

必须严格限制执行权。创建钱包会通过 SQL 结果返回私钥，转账调用还把私钥作为 SQL 参数接收，因此私钥可能暴露在客户端日志、语句日志、活动视图、监控或查询历史中。区块链转账和空投属于外部副作用，不会随外围 PostgreSQL 事务回滚。

RPC 调用是同步的，各命名网络使用固定的 Solana 公共端点。目录版本为 0.0.0，源码使用 pgrx 0.11.3，并为 PostgreSQL 11 至 16 提供构建特性；授权前应确认实际编译包与服务器匹配，并先在非生产网络测试。
