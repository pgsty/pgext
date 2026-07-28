## 用法

来源：

- [官方上游 README](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/README.md)
- [官方扩展控制文件 (pg_table_bloat.control)](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/pg_table_bloat.control)
- [官方实现源代码](https://github.com/lucasfrederico/pg_table_bloat/blob/686385d88dfc8318aa57351ae260c87399f70868/src/lib.rs)

`pg_table_bloat` — > 一个基于 PostgreSQL 的扩展，用于从 pg_class 和 pg_stats 估算表的膨胀情况，无需使用 pgstattuple，无需逐页加锁，适用于任何大小的表。在进行数据库管理或自动化上述描述的行为时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_table_bloat;

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    payload jsonb
);
INSERT INTO users (email, payload)
SELECT 'user'||g||'@example.com', jsonb_build_object('data', repeat('x', 500))
FROM generate_series(1, 50000) g;
DELETE FROM users WHERE id > 500;   -- delete 99% of rows
ANALYZE users;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该扩展的目录记录版本 `0.1.0`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
