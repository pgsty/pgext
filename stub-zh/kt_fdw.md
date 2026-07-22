## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/README.md)
- [0.0.1 版本 SQL 对象](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/sql/kt_fdw.sql)
- [FDW 实现](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/src/kt_fdw.c)
- [可选的 Kyoto Tycoon 事务过程](https://github.com/cloudflarearchive/kt_fdw/blob/add3f13721de44d8bbe7ab84ad15cc6e475270ee/transactions.lua)

`kt_fdw` 把 Kyoto Tycoon 键值数据库映射成可写的 PostgreSQL 外部表。它支持读取以及插入、更新和删除，但上游仓库已经归档，代码面向旧版 PostgreSQL 与 Kyoto Tycoon API。考虑使用前，应在隔离环境中针对确切版本重新构建并完成集成测试。

### 核心流程

定义服务器端点，并创建严格包含两列的外部表：第一列是键，第二列是值。

```sql
CREATE EXTENSION kt_fdw;

CREATE SERVER kt_cache
  FOREIGN DATA WRAPPER kt_fdw
  OPTIONS (host '127.0.0.1', port '1978', timeout '5');

CREATE USER MAPPING FOR CURRENT_USER SERVER kt_cache;

CREATE FOREIGN TABLE cache_entry (
  key text,
  value text
) SERVER kt_cache;

SELECT value FROM cache_entry WHERE key = 'session:42';
INSERT INTO cache_entry VALUES ('session:42', '{"state":"ready"}');
UPDATE cache_entry SET value = '{"state":"done"}' WHERE key = 'session:42';
DELETE FROM cache_entry WHERE key = 'session:42';
```

已复核的服务器选项只有 `host`（默认 `127.0.0.1`）、`port`（默认 `1978`）和 `timeout`（默认 `-1`）。扩展没有身份认证、TLS 或远端数据库编号选项，因此端点应只放在可信私有网络中。

### 查询与写入语义

第一列上的简单等值条件可以转为一次远程键查询。其他扫描会遍历完整的远端游标，再由 PostgreSQL 过滤。两列都不能为 NULL。插入使用 Kyoto Tycoon 的 add 操作，键已存在时失败；更新只能修改值，不能修改键。

事务性写入要求在编译扩展时启用其事务支持、使用带 Lua 的 Kyoto Tycoon，并在服务器端加载随附的 `transactions.lua` 过程。即使如此，也不支持保存点和预备事务。没有这些设置时，远程写入不会跟随 PostgreSQL 事务回滚。跨系统故障与重试必须作为应用可见的一致性事件处理。
