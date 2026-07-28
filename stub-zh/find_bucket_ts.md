## 用法

来源：

- [官方上游 README](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/README.md)
- [官方扩展控制文件 (find_bucket_ts.control)](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/src/find_bucket_ext/find_bucket_ts/find_bucket_ts.control)
- [官方扩展 SQL (find_bucket_ts--1.0.sql)](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/src/find_bucket_ext/find_bucket_ts/find_bucket_ts--1.0.sql)

`find_bucket_ts` — Artifacts of BaCon. 仅供 VLDB26 审查人使用。用于相应的调度、时间或时间序列工作流。在目标 PostgreSQL 构建上使用上述链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION find_bucket_ts;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `find_bucket_ts(val timestamp without time zone, buckets timestamp without time zone[][], need_null boolean)` 是一个扩展函数，返回 `int`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
