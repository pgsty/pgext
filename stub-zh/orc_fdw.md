## 用法

来源：

- [官方 README](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/README.md)
- [官方扩展 SQL](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw--1.0.sql)
- [官方 FDW 实现](https://github.com/gokhankici/orc_fdw/blob/b72f3883cf5d10058f60f187b4c9190607258e44/orc_fdw.c)

`orc_fdw` 1.0 是只读外部数据包装器，用于查询本地 Apache ORC 格式文件。外部表把 PostgreSQL 列映射到 ORC 文件，并要求提供服务端 `filename` 选项。

### 核心流程

```sql
CREATE EXTENSION orc_fdw;

CREATE SERVER orc_files
  FOREIGN DATA WRAPPER orc_fdw;

CREATE FOREIGN TABLE orc_events (
  event_id bigint,
  event_name text
)
SERVER orc_files
OPTIONS (filename '/srv/postgresql/orc/events.orc');

SELECT event_id, event_name
FROM orc_events
WHERE event_id >= 1000;
```

该路径在数据库服务器而非客户端解析。运行 PostgreSQL 的操作系统账号必须能够打开文件并读取其元数据。

### 重要对象

- `orc_fdw_handler()` 提供 PostgreSQL FDW 回调表。
- `orc_fdw_validator(text[], oid)` 校验对象选项，并要求每个外部表都设置 `filename`。
- `orc_fdw` 是安装后的外部数据包装器。
- `filename` 是固定源码中唯一实现的外部表选项。

### 运维说明

实现注册了扫描与 EXPLAIN 回调，但没有 insert、update 或 delete 回调，因此只能用于读取。文件位于每个服务器进程的本地文件系统；若故障切换后查询需要保持一致，备节点必须具有相同路径和内容。上游构建说明面向 protobuf/protobuf-c 以及与 Hive 0.12 相关的 ORC 数据，表明这是较旧且范围有限的格式实现。采用前必须验证列类型映射、压缩、谓词行为以及精确 ORC 写入器生成的文件；不要仅凭格式名称假设它兼容当前 ORC 生态。
