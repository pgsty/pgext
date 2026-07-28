## 用法

来源：

- [官方上游 README](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/README.md)
- [官方扩展控制文件 (approx_count.control)](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/approx_count.control)
- [官方扩展 SQL (approx_count--1.0.sql)](https://github.com/jmealo/pg_approx_count/blob/341dfa19f73e60d22a8869ccb03bd252d888cec7/sql/approx_count--1.0.sql)

`approx_count` — 在 PostgreSQL 14+ 中快速获取表和索引的近似行数，从 pg_class.reltuples 读取而不是通过磁盘密集型的精确 COUNT(*) 扫描。当统计信息过时时，使用受控的 ANALYZE 刷新。在 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION approx_count;

-- On a ~300M-row table an exact count scans for minutes:
SELECT count(*) FROM events;     -- minutes, heavy I/O
-- approx_count reads the planner's cached estimate:
SELECT approx_count.approx_count('events');   -- ~0.2 ms
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `IF` 是由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码进行比对。
