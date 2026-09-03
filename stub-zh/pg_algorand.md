## 用法

来源：

- [官方文档](https://github.com/AlgoNode/pg_algorand/blob/7bc3552bdb84a684b75dde2c7c8e91aaafa089cc/README.md)
- [扩展控制文件](https://github.com/AlgoNode/pg_algorand/blob/7bc3552bdb84a684b75dde2c7c8e91aaafa089cc/pg_algorand.control)
- [官方仓库](https://github.com/AlgoNode/pg_algorand)

`pg_algorand` 提供 Algorand 地址、应用、资产、交易与校验编码工具。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_algorand`：

```sql
CREATE EXTENSION pg_algorand;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT
  COUNT(*) FROM account
WHERE
  addr = AddressTxt2Bin('ALGONODEIBJTET5OSEAXIHDSIEG7C2DOFB2WDYLRZTXN3NXVJ3NJD26L4E');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `AddressBin2Txt` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `AddressTxt2Bin` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `GetNFDSigNameLSIG` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `GetNFDSigRevAddressBinLSIG` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `GetNFDSigRevAddressLSIG` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `algoaddr` | TYPE | 扩展创建的用户数据类型。 |
| `algoaddr_in` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `algoaddr_out` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
