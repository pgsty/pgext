## 用法

来源：

- [官方 README](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/README.md)
- [2.0.0 版本扩展 SQL](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/sql/pgtt_rsl--2.0.0.sql)
- [解析器钩子实现](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_rsl.c)
- [后台工作进程实现](https://github.com/darold/pgtt-rsl/blob/f48f1a90bc09a80cb6f9d596535277960f658be9/pgtt_bgw.c)

`pgtt_rsl` 使用永久非日志底层表、行级安全策略和安全屏障视图实现 Oracle/DB2 风格全局临时表。表定义会持久存在，并且可以位于普通模式中；行则按会话或事务隔离。它的主要取舍是避免重复目录膨胀，但数据路径性能低于原生临时表。

### 核心流程

创建或使用扩展前，先预加载两个库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pgtt_bgw,pgtt_rsl'
pgtt_bgw.naptime = 5
pgtt_bgw.analyze = off
pgtt_bgw.chunk_size = 250000
```

```sql
CREATE EXTENSION pgtt_rsl;

CREATE GLOBAL TEMPORARY TABLE work_items (
    id integer,
    label text
) ON COMMIT PRESERVE ROWS;

INSERT INTO work_items VALUES (1, 'session private');
SELECT * FROM work_items;

DROP TABLE work_items;
```

当生成的 SQL 无法使用被拦截语法时，可以调用显式辅助函数：

```sql
SELECT pgtt_schema.pgtt_create_table(
    'txn_items',
    'id integer, label text',
    false,
    'public'
);
```

`false` 表示事务级行（相当于 `ON COMMIT DELETE ROWS`），`true` 表示在会话期间保留行。

### 对象索引

- `pgtt_schema.pgtt_create_table(name, definition, preserved DEFAULT false, schema DEFAULT 'public')` 创建全局临时表外观。
- `pgtt_schema.pgtt_drop_table(name, schema DEFAULT 'public')` 删除它。
- `lsid` 和 `get_session_id()` 在内部标识所属会话。
- `pgtt_bgw.naptime`、`pgtt_bgw.analyze` 和 `pgtt_bgw.chunk_size` 控制陈旧行清理。
- `pgtt_schema.pgtt_maintenance(...)` 是内部接口，只能由后台工作进程调用。

### 运维说明

版本 `2.0.0` 不可重定位，并且要求以重启级方式预加载。后台工作进程扫描所有包含该扩展的数据库，并删除已结束会话或事务的行。应监控工作进程健康、死元组、自动清理和清理延迟；行级安全、视图、索引和删除操作使大规模负载比原生临时表更慢。

底层表是 `UNLOGGED`：其内容不具备崩溃持久性，也不适合可复制或可恢复状态。创建辅助函数会向 `PUBLIC` 授予底层表 DML，并依靠强制 RLS 与外观视图隔离；应审查权限，且绝不要直接访问 `pgtt_schema.pgtt_*` 底层表。迁移负载前，应测试会话断开、事务回滚、重启、故障转移、连接池行为和扩展升级。
