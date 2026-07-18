## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.control)

`palloc_bench` — palloc_bench：用于基准测试 PostgreSQL palloc 内存分配开销的扩展函数。

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "palloc_bench";
SELECT extversion
FROM pg_extension
WHERE extname = 'palloc_bench';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
