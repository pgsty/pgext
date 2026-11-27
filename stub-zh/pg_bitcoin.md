## 用法

来源：

- [README.md](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/README.md)
- [Cargo.toml](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/Cargo.toml)
- [pg_bitcoin.control](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/pg_bitcoin.control)
- [src/lib.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/lib.rs)
- [src/config.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/config.rs)
- [src/client_wrapper.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/client_wrapper.rs)
- [src/electrum.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/electrum.rs)
- [tests/pg_regress/sql/setup.sql](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/tests/pg_regress/sql/setup.sql)
- [LICENSE](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/LICENSE)

`pg_bitcoin` 通过 Electrum 服务器查询比特币地址余额、历史和未花费输出。已核对源码声明的版本为 0.0.0，属于早期源码版本，尚无公开发行版。构建清单声明支持 PostgreSQL 16 和 18，默认使用 18。

### 启用与配置

该扩展不受信任且不可重定位，需要超级用户安装。普通用法不要求共享预加载。先准备与所选网络匹配、可访问的 Electrum 索引服务，再配置当前后端：

```sql
CREATE EXTENSION pg_bitcoin;
SET pg_bitcoin.network_type = 'regtest';
SET pg_bitcoin.electrum_url = 'tcp://127.0.0.1:60401';
SELECT bitcoin_ping();
```

两项设置均需超级用户权限，默认值分别为 `regtest` 和 `tcp://127.0.0.1:60401`。网络可选 `mainnet`、`regtest`、`testnet`、`testnet4` 和 `signet`。地址必须属于所选网络；修改任一设置后，下次调用会重新创建缓存连接。

### SQL 接口

`get_confirmed_balance(text)` 返回已确认余额。`get_balance(text)` 和 `get_balance(text[])` 返回地址、已确认金额和未确认金额。`get_history(text)` 返回高度、交易哈希和可选手续费，数组重载还会标明地址。`list_unspent(text)` 返回高度、交易哈希、输出位置和金额。调用时应提供当前网络的真实地址。

金额以整数聪为单位。函数同步查询索引服务，不在本地复制区块链，也不实现钱包。`bitcoin_ping()` 在连接或探测失败时返回 false；数据函数遇到无效地址或传输失败时会抛出 PostgreSQL 错误。

### 运行边界

每个后端缓存一个 Electrum 客户端，连接配置采用五秒超时和一次重试。结果与延迟取决于所选索引服务，示例 TCP 端点未加密。应限制 SQL 调用权限，并按预期的信任边界选择端点。扩展提供查询函数，不保管私钥，也不提交交易。当前修订未提供 SQL 升级脚本。
