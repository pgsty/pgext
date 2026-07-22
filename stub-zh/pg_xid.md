## 用法

来源：

- [Official README](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/README.md)
- [Version 1.0 SQL API](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid--1.0.sql)
- [PostgreSQL entry points](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid.c)
- [Identifier generation and encoding](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/xid.c)

`pg_xid` 1.0 生成采用经典 MongoDB ObjectID 布局的 12 字节标识符：四字节秒级时间戳、三字节主机名派生值、两字节进程 ID 和三字节进程内计数器。它不需要中央序列，并可按创建时间大致排序。`xid()` 返回原始 `bytea`；`xid_encoded()` 返回 20 字符的自定义 Base32 文本。

### 生成标识符

```sql
CREATE EXTENSION pg_xid;

SELECT encode(xid(), 'hex') AS raw_hex;
SELECT xid_encoded() AS compact_text;
```

应为一个字段选择并始终使用同一种表示：

```sql
CREATE TABLE event (
  id bytea PRIMARY KEY DEFAULT xid(),
  payload jsonb NOT NULL
);
```

二进制形式为 12 字节，普通十六进制渲染为 24 字符。编码函数使用上游的有序字母表 `0-9a-v`，生成 20 字符，并非 MongoDB 常见的十六进制文本。

### 运维与安全边界

前导时间戳只能让值大致按时间排序：同一秒生成的所有 ID 会按主机/进程/计数器分量排序，而不是按全局时钟排序。时钟变化、主机名复用、同名容器、PID 复用、计数器回绕和克隆环境都会影响碰撞假设。24 位计数器在每个初始化的主机/进程上下文中每秒生成 16,777,216 个值后回绕。

这些值不是秘密，也不是密码学随机令牌；它们会暴露创建时间以及主机和进程派生位。不要用于认证、能力 URL 或需要隐藏信息的标识符。

上游没有现代 PostgreSQL 兼容矩阵或升级系列。C 代码使用 OpenSSL MD5 生成主机名分量，并用 PostgreSQL API 初始化强随机计数器。应在每个目标架构和 PostgreSQL 大版本上构建及并发测试，并在作为分布式主键前验证备份恢复以及字节/文本排序。
