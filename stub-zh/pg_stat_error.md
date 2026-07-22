## 用法

来源：

- [TencentDB for PostgreSQL 支持的扩展版本](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB for PostgreSQL 9.5 扩展矩阵](https://cloud.tencent.com/document/product/409/75123)

`pg_stat_error` 1.0 是 TencentDB for PostgreSQL 托管扩展。腾讯云公布了可用性矩阵，但没有发布权威的 SQL 对象参考或行为契约。不能从名称推测其视图、函数、保留策略、收集范围或权限要求。

### 可用性与安装

创建前检查实际 TencentDB 实例：

```sql
SELECT name, version, installed
FROM pg_available_extension_versions
WHERE name = 'pg_stat_error';

CREATE EXTENSION pg_stat_error;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'pg_stat_error';
```

当前服务商概览为 PostgreSQL 10–13 列出 1.0，而 PostgreSQL 14–18 不支持。更早服务商页面还为 TencentDB PostgreSQL 9.3 和 9.5 内核版本列出版本 1。可用性还可能取决于托管内核修订，因此应以实例目录为最终判断。

### 安全发现边界

安装后，可以在 `psql` 中使用 `\dx+ pg_stat_error`，或由管理员检查 `pg_depend`/`pg_proc`/`pg_class`，识别精确托管构建安装的对象。在授权访问或基于它构建监控之前，应向腾讯云获取关于列、重置行为、保留、可见性和开销的支持文档。

该扩展不是社区软件包，不能通过把库复制到原生 PostgreSQL 来复现。引用的服务商资料只证明可用性和版本；不能据此认定它捕获所有错误、替代 PostgreSQL 日志、跨故障转移持久化，或向普通用户暴露数据。
