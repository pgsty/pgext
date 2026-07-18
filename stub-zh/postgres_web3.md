## 用法

来源：

- [项目 README](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/README.md)
- [扩展 control 文件](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/postgres_web3.control)
- [1.0 版 SQL 定义](https://github.com/Yen/postgres_web3/blob/85e356422b38390d68796483068f9c1913f6a7f3/postgres_web3--1.0.sql)

`postgres_web3` 1.0 提供定宽的 `int128`、`uint128`、`int256`、`uint256`、`hex160` 和 `hex256` 类型，用于区块链尺度的整数、地址和哈希。它定义了算术、比较和位运算符、类型转换、B-tree 与 hash 运算符类以及数值聚合。

### 建模 Web3 数值

```sql
CREATE EXTENSION postgres_web3;

CREATE TABLE chain_transfer (
    tx_hash hex256 PRIMARY KEY,
    from_address hex160 NOT NULL,
    amount uint256 NOT NULL
);

INSERT INTO chain_transfer (tx_hash, from_address, amount)
VALUES ('01'::hex256, '88ff'::hex160,
        '340282366920938463463374607431768211456'::uint256);

SELECT amount + '1'::uint256
FROM chain_transfer
WHERE from_address = '88ff'::hex160;
```

整数输入使用十进制文本。十六进制输入不区分大小写，不含 `0x` 前缀，输出为小写文本。算术运算会检查范围，无法放入目标类型的结果会报错。

### 构建与精度限制

已审查的 README 要求 C23 `_BitInt`、Clang 16，以及不带 LLVM 支持的特殊 PostgreSQL 构建，并建议使用其预构建 Linux 容器而非源码构建。这些约束既陈旧又狭窄；存储持久生产数据前，应验证编译器、架构、PostgreSQL ABI、备份、恢复、复制和客户端驱动行为。

`avg` 聚合使用双精度内部求和，上游文档警告它对大数求和会产生不准确结果。平均值影响余额或账务时应改用精确方案。标为赋值转换的 cast 仍可能因负数或越界值失败，而且定宽十六进制输出并不实现 Ethereum checksum address。
