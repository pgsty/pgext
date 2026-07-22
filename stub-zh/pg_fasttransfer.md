## 用法

来源：

- [官方 README](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/README.md)
- [1.0 版本 SQL 定义](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/pg_fasttransfer--1.0.sql)
- [1.0 版本 C 实现](https://github.com/arpe-io/pg_fasttransfer/blob/a06073f50427935d1ef9ecfea101b4b77b323a98/pg_fasttransfer.c)
- [FastTransfer 产品文档](https://fasttransfer.arpe.io/)

`pg_fasttransfer` 允许 PostgreSQL 后端启动单独安装的 FastTransfer 程序，并把进程结果以行返回。它是外部授权数据传输工具的管理桥接，而不是数据库内部的复制协议。

### 核心流程

安装 `pgcrypto` 和该扩展，让 PostgreSQL 操作系统账号可执行 FastTransfer 并可读授权文件，然后使用明确的源、目标和二进制路径调用管理函数：

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_fasttransfer;

SELECT *
FROM xp_RunFastTransfer_secure(
    sourceconnectiontype := 'pgsql',
    sourceserver := '127.0.0.1:5432',
    sourceuser := 'transfer_user',
    sourcepassword := pg_fasttransfer_encrypt('source-secret'),
    sourcedatabase := 'source_db',
    sourceschema := 'public',
    sourcetable := 'orders',
    targetconnectiontype := 'msbulk',
    targetserver := '127.0.0.1,1433',
    targetuser := 'transfer_user',
    targetpassword := pg_fasttransfer_encrypt('target-secret'),
    targetdatabase := 'target_db',
    targetschema := 'dbo',
    targettable := 'orders',
    fasttransfer_path := '/opt/fasttransfer',
    license := '/opt/fasttransfer/FastTransfer.lic'
);
```

### 对象与结果

- `pg_fasttransfer_encrypt(text)` 生成运行器可接受的编码密码格式。
- `xp_RunFastTransfer_secure(...)` 把命名参数映射为 FastTransfer 命令行选项，并返回 `exit_code`、`output`、`error_message`、`total_rows`、`total_columns`、`transfer_time_ms` 和 `total_time_ms`。
- 版本 `1.0` 需要 `pgcrypto`。上游报告的已测试组合为 Linux 上的 PostgreSQL 15 至 17，以及 Windows 上的 16 至 17；其他组合没有经过文档化的验证。

### 安全与运维

已复核的 C 源码会根据 SQL 参数构造 shell 命令并调用 `popen`；参数值虽被放在双引号中，但并未转义内嵌引号和 shell 元字符。应撤销应用角色和不可信角色的执行权，只允许管理员提交经审核的固定输入。同一源码把密码加密密钥嵌入为常量 `key`，因此 `pg_fasttransfer_encrypt(text)` 只是混淆，无法防护可检查二进制或源码的人。进程输出也可能包含敏感连接或传输信息。

传输使用数据库服务器账号的文件系统、进程和网络权限，并可执行远程批量写入。启用破坏性加载模式前，应使用最小权限源账号、严格限定的目标账号、受保护的授权和设置文件、扩展外部的显式超时，以及经演练的恢复计划。
