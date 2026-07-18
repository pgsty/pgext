## 用法

来源：

- [上游 pg_comparator 手册入口](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/README.pg_comparator)
- [扩展 SQL 定义](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp.sql)
- [C 支撑实现](https://github.com/zx80/pg_comparator/blob/01cd47b2fae021b3320eea2f0885297a320d6ec4/pgcmp.c)

`pgcmp` 是 `pg_comparator` 命令行数据比对与同步工具在 PostgreSQL 端的支撑扩展。它提供外部程序使用的服务端校验和与类型转换。

```sql
CREATE EXTENSION pgcmp;
SELECT extversion FROM pg_extension WHERE extname = 'pgcmp';
```

源码版本为 `3.1`，但项目已经废弃，最新 Git tag 还更旧。比对与同步工具可能在两端数据库产生大量扫描和写入；应先只读比对，保护凭据，核实排序规则与类型语义，并在启用任何修复操作前演练恢复。当前系统应优先使用受维护的逻辑复制或专用对账方案。
