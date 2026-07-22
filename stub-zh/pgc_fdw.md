## 用法

来源：

- [官方 README](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/README.md)
- [1.0 版本 SQL 对象](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/pgc_fdw--1.0.sql)
- [FDW 选项校验](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/option.c)
- [FoundationDB 缓存实现](https://github.com/fengttt/pgc_fdw/blob/5993bb88886c0a27d2ee6086f68255ba9a8460c2/cache.c)

`pgc_fdw` 是一个旧的 `postgres_fdw` 分支，用 FoundationDB 缓存远程查询结果。它创建名为 `pgc_fdw` 的外部数据包装器，因此可以与核心 `postgres_fdw` 共存；但它跟随的是旧 PostgreSQL 源码树，应视为实验性软件。

### 创建缓存外部表

必须安装 FoundationDB 服务端与客户端库，且 PostgreSQL 操作系统账号必须能够使用默认 FoundationDB 集群文件。

```sql
CREATE EXTENSION pgc_fdw;

CREATE SERVER cached_remote
  FOREIGN DATA WRAPPER pgc_fdw
  OPTIONS (host 'db.internal', dbname 'app', cache_timeout '60');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER cached_remote
  OPTIONS (user 'remote_user', password 'secret');

CREATE FOREIGN TABLE cached_items (id bigint, payload text)
  SERVER cached_remote
  OPTIONS (schema_name 'public', table_name 'items', cache_timeout '30');

SELECT * FROM cached_items;
```

`cache_timeout` 可用于服务器或外部表，以秒为单位，默认值为 `3600`；设为 `0` 可禁用缓存。负值会被拒绝。

### 缓存操作与风险

```sql
SELECT * FROM pgc_fdw_cache_info();
SELECT pgc_fdw_invalidate('cache-entry-sha');
```

`pgc_fdw_cache_info()` 报告 `sha`、时间戳、元组数与查询。扩展还暴露低层的 `pgc_fdw_set()` 与 `pgc_fdw_watch()` 函数。除非调用者明确需要，否则应撤销 `PUBLIC` 对缓存修改函数的权限。

缓存读取可能在整个超时时间内保持过期，且扩展没有把远端写入与通用失效机制关联起来。必须测试事务可见性、故障转移、FoundationDB 故障与限制、凭据处理、缓存键隔离及手工失效。该分支依赖 PostgreSQL 内部接口，并在每个加载后端中启动 FoundationDB 网络，因此只能针对计划部署的准确 PostgreSQL/FoundationDB 组合构建并做压力测试。
