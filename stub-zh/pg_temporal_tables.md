## 用法

来源：

- [README.md](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/README.md)
- [sql/pg_temporal_tables--0.1.1.sql](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/sql/pg_temporal_tables--0.1.1.sql)
- [pg_temporal_tables.control](https://github.com/GerardSmit/pg_temporal_tables/blob/bd849266516293ca7e3b41225fdc61284ef66843/pg_temporal_tables.control)

`pg_temporal_tables` 记录行历史，并按选定的历史时刻改写查询。当前实现面向 PostgreSQL 18，要求预加载并重启服务器。

### 核心用法

```ini
shared_preload_libraries = 'pg_temporal_tables'
```

```sql
CREATE EXTENSION pg_temporal_tables;
CREATE TABLE users (id bigint PRIMARY KEY, name text);
SELECT temporal.enable('public.users');
SET temporal.user_id = 'operator';
INSERT INTO users VALUES (1, 'Alice');
UPDATE users SET name = 'Alicia' WHERE id = 1;
```

### 运行边界

被追踪表必须有主键。`temporal.enable()` 建立历史追踪，`temporal.disable()` 停止追踪。`temporal.changes()` 与 `temporal.column_history()` 查看变化，保留与压缩函数维护历史。用 `temporal.user_id` 标识操作用户，历史查询后重置 `temporal.as_of`。分区表应追踪根表，不支持传统继承表和 TimescaleDB 超表。更改或复用主键之前须核对谱系选项。安装要求超级用户。历史与基础数据应一同备份，删除历史前先明确保留要求。
