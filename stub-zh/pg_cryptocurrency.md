## 用法

来源：

- [README](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/README.md)
- [Control file / 控制文件](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/pg_cryptocurrency.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/docs/SECURITY.md)

`pg_cryptocurrency` 在 PostgreSQL 14–18 中提供精确的区块链资产金额、含网络信息的资产身份、地址与无符号 256 位整数，不访问区块链或发起网络请求。

### 核心工作流

```sql
CREATE EXTENSION pg_cryptocurrency;
SELECT '1.25 BTC'::crypto_amount;
SELECT '1500 USDC@ethereum'::crypto_amount;
SELECT crypto_to_units('10 ETH');
SELECT crypto_from_units(1000000000000000001::bigint, 'ETH');
SELECT '1 BTC'::crypto_amount + '0.25 BTC';
```

### 对象与身份

`crypto_amount` 将精确金额绑定到资产。`crypto_asset` 标识原生货币和代币，`crypto_address` 验证包含网络信息的地址，`uint256` 保存无符号 256 位值。`crypto_register_asset` 根据符号、网络、合约和小数位数注册代币。`crypto_amount_value`、`crypto_amount_asset` 与 `crypto_amount_network` 用于检查值。

代币符号不等于资产身份，网络和合约同样重要。不同资产之间的算术会报错，不会隐式兑换。转换为最小单位时会检查精度和金额范围；地址格式合法并不能证明所有权或链上存在。

### 运行

扩展可重定位，无需预加载。资产注册等管理操作应单独检查权限。`pg_money` 可选适配器独立于核心类型工作流；安装扩展不会建立钱包、托管服务或汇率来源。
