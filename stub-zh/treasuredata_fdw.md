## 用法

来源：

- [当前上游固定版本 README](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/README.md)
- [固定版本扩展控制文件](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/treasuredata_fdw.control)
- [固定版本发行元数据](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/META.json)
- [固定版本选项与凭据处理](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/option.c)
- [固定版本 Rust 桥接依赖](https://github.com/treasure-data/treasuredata_fdw/blob/afbecb707e3757b43a9c7dcfe729721f8269cbd0/bridge_td_client_rust/Cargo.toml)
- [正式 PGXN 发行页](https://pgxn.org/dist/treasuredata_fdw/)

treasuredata_fdw 通过 Hive 或 Presto 作业把 PostgreSQL 连接到 Treasure Data 服务。外部表可以查询远端表或固定远端 SQL，下推部分过滤与聚合，以流式或下载方式读取结果，分块插入数据，并通过 IMPORT FOREIGN SCHEMA 生成表定义。

### 定义只读外部表

所审验证器要求把凭据和连接选项直接放在外部表上：

```sql
CREATE EXTENSION treasuredata_fdw;

CREATE SERVER treasuredata_server
FOREIGN DATA WRAPPER treasuredata_fdw;

CREATE FOREIGN TABLE sample_access (
    time integer,
    host text,
    path text,
    code integer
) SERVER treasuredata_server
OPTIONS (
    apikey '<api-key>',
    database 'sample_datasets',
    query_engine 'presto',
    table 'www_access'
);

SELECT code, count(*)
FROM sample_access
WHERE time BETWEEN 1412121600 AND 1414800000
GROUP BY code;
```

扩展对象版本为 1.2，与本目录一致。固定仓库元数据标为发行版 1.2.16，而采集到的 PGXN 页面把 1.2.15 列为最新稳定版。不要把这些包版本直接与 SQL 对象版本比较。

### 远端事务与秘密边界

必需的 API 密钥保存在外部表选项，而不是受保护的用户映射中。目录访问、模式转储、备份、支持包、IMPORT 生成的命令和管理日志都可能暴露它。应限制目录与表所有权，使用最小权限密钥，不共享定义，定期轮换并验证服务端点。

查询与导入都是同步远端作业，可能运行数小时并持续占用 PostgreSQL 后端。可选的结果下载目录会在数据库服务器写压缩文件。atomic_import 关闭时，INSERT 上传的分块可能在本地语句提交前就对远端可见；PostgreSQL 中止不会回滚已经完成的远端导入。原子模式会使用远端临时表并最终复制，但仍不是 PostgreSQL 两阶段提交，失败后可能残留远端资源。

Rust 桥接使用陈旧的 HTTP、时间、重试与 Treasure Data 客户端依赖，包含 unwrap 与 panic 路径，仓库最后源码提交停留在 2023 年。必须测试 TLS、取消、超时、重试重复、部分导入、清理、不支持类型、远端 SQL 注入边界、本地磁盘耗尽、模式漂移和当前 Treasure Data API。只有在记录成本与数据出站策略后才授予外部表访问权。
