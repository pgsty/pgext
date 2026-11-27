## 用法

来源：

- [0.2.3](https://crates.io/crates/pg_wildcard/0.2.3)
- [pg_wildcard/README.md](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/pg_wildcard/README.md)
- [pg_wildcard/pg_wildcard.control](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/pg_wildcard/pg_wildcard.control)
- [pg_wildcard/Cargo.toml](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/pg_wildcard/Cargo.toml)
- [docs/sql-api.md](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/docs/sql-api.md)
- [docs/fdw.md](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/docs/fdw.md)
- [docs/operations.md](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/docs/operations.md)
- [pg_wildcard/sql/pg_wildcard--0.2.2--0.2.3.sql](https://github.com/f1b0-gh/pg_wildcard/blob/v0.2.3/pg_wildcard/sql/pg_wildcard--0.2.2--0.2.3.sql)

`pg_wildcard` 0.2.3 使用通配符和正则表达式搜索遥测字符串，可提取指标并生成 ClickHouse/Antalya 查询，同时提供只读扫描的 `wildcard_fdw`。已发布的 crate 使用 `antalya-rs` 0.1.1；经过验证的服务器主版本是 PostgreSQL 18 和 19。声明的 PostgreSQL 14–17 构建特性需要另行验证。

### 本地字符串搜索

扩展不受信任且不可重定位，需要超级用户安装。上游没有为普通用法规定共享预加载或服务器重启。

```sql
CREATE EXTENSION pg_wildcard;
SELECT wc_match('https://example.org/admin', '*/admin', 'raw');
SELECT wc_match('GET /%77p-ADMIN/a.PHP', '*/wp-admin/*.php', 'decoded');
SELECT * FROM wc_indicator('hxxps://Login.Evil[.]com/');
SELECT wc_explain('*/wp-admin/*.php', 'wildcard', 'decoded');
```

转换视图包括 `raw`、`lower`、`decoded`、`revdomain` 和 `winpath`。`wc_match`、`wc_like`、`wc_regex` 匹配文本，`wc_regex_bytes` 处理字节。`wc_url_decode` 保留解码后的字节，`wc_view` 则生成显示文本。`wc_explain` 描述保守的候选集合，最终匹配仍须由精确谓词判定。正则语义遵循所支持的 RE2 方言。

### 远程查询与数据湖元数据

`wc_ch_predicate` 等辅助函数返回 SQL，由调用者自行执行；其可选过滤表达式属于受信任 SQL 片段。`wc_lake_index`、`wc_lake_file` 和 `wc_lake_gram` 保存导入的剪枝元数据，候选文件始终需要精确复查，覆盖不足时必须扩大扫描范围。

`wildcard_fdw` 需要 Antalya 端点、带有 `bearer_token` 的用户映射、兼容的标量列以及 UTF-8 PostgreSQL 数据库。普通 ClickHouse HTTP 服务可用并不足够。映射对应的身份需要访问源数据及文档规定的查询完成证据，通常来自 `system.query_log`。应保护映射凭据，并独立维护令牌刷新。

包装器下推受支持的谓词，同时保留本地检查。连接和聚合在本地执行，外部表没有本地 PostgreSQL 索引。不支持的排序规则或查询形态会限制下推。即使已经向客户端传输部分行，后续远程错误或完成证据错误仍会令语句失败。

### 资源限制与升级

`pg_wildcard.matcher_cache_entries` 与 `pg_wildcard.matcher_cache_kb` 限制每个后端保留的模式缓存，`wc_matcher_cache_stats()` 提供统计。这些限制不是后端总内存上限。

从 0.2.0、0.2.1 或 0.2.2 升级时，先安装匹配的 0.2.3 库和 SQL 文件，重新连接，再执行：

```sql
ALTER EXTENSION pg_wildcard UPDATE TO '0.2.3';
```

该更新保留 SQL 签名并记录新扩展版本。独立库、数据湖命令行工具和远程引擎配置是分别维护的组件。

可选的批处理和分阶段扫描在 FDW 之外运行，执行、准入限制和清理由调用者负责。FDW 不会自动并行执行查询，也未将压缩作为正式 FDW 选项开放。
