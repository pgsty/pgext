## 用法

来源：

- [官方文档](https://github.com/whitslack/pg_bitcoin_address/blob/1ab081b6a55cc4bfc8b01c63c0637fcb4d8ae6b6/README.md)
- [扩展控制文件](https://github.com/whitslack/pg_bitcoin_address/blob/1ab081b6a55cc4bfc8b01c63c0637fcb4d8ae6b6/pg_bitcoin_address.control)
- [官方仓库](https://github.com/whitslack/pg_bitcoin_address)

`pg_bitcoin_address` 提供比特币地址类型以及 Base58Check、Bech32/Bech32m 编解码。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_bitcoin_address`：

```sql
CREATE EXTENSION pg_bitcoin_address;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
=> SELECT pg_column_size('1BitcoinEaterAddressDontSendf59kuE'::text);
pg_column_size | 38

=> SELECT pg_column_size('1BitcoinEaterAddressDontSendf59kuE'::bitcoin_address);
pg_column_size | 26

=> SELECT pg_column_size('bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4'::text);
pg_column_size | 46

=> SELECT pg_column_size('bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4'::bitcoin_address);
pg_column_size | 26
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `bitcoin_address` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `program` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `base58check` | TYPE | 扩展创建的用户数据类型。 |
| `version` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `hrp` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `is_blinding` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `is_liquidtestnet` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `is_liquidv1` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
