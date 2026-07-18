## 用法

来源：

- [已复核 commit 的归档上游 README](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/README.md)
- [扩展 control 文件](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/src/clickhouse_fdw.control)
- [FDW 实现](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/src/clickhouse_fdw.c)

`clickhouse_fdw` 是一个已归档的 ClickHouse 外部数据包装器。已复核版本支持 HTTP 与二进制驱动、远端模式导入、读取、插入和部分聚合下推；上游建议新用户改用其持续维护的后继项目。

```sql
CREATE EXTENSION clickhouse_fdw;
CREATE SERVER clickhouse_svr
  FOREIGN DATA WRAPPER clickhouse_fdw
  OPTIONS (dbname 'analytics', host 'clickhouse.internal', port '8123');
CREATE USER MAPPING FOR CURRENT_USER
  SERVER clickhouse_svr OPTIONS (user 'reporter', password 'secret');
IMPORT FOREIGN SCHEMA "analytics"
  FROM SERVER clickhouse_svr INTO clickhouse_ext;
```

部署 SQL 中不要保留明文密码，应使用受限的配置流程和目录访问权限。HTTP 驱动会向远程服务发送查询，因此应在支持时强制 TLS，使用最小权限 ClickHouse 凭据、网络白名单、超时与资源限制。按上游要求让 ClickHouse 使用 UTC，并验证类型、空值、时间戳与聚合映射。

该仓库已归档，新部署应使用后继项目。遗留系统必须固定 libcurl、UUID 等依赖，通过 `EXPLAIN VERBOSE` 检查远端 SQL，并在依赖它之前测试事务语义、部分写入、重试、取消、故障转移、谓词下推和大版本升级。
