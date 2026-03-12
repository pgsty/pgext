


## 用法

> [pg_stat_kcache: PostgreSQL 内核缓存统计](https://github.com/powa-team/pg_stat_kcache)

pg_stat_kcache 收集文件系统层实际读写的统计信息，以及每个查询的 CPU 使用情况。它依赖 `pg_stat_statements`。

### 视图

**`pg_stat_kcache`** -- 按数据库汇总的统计信息：

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `datname` | name | 数据库名称 |
| `exec_user_time` | double precision | 执行语句的用户 CPU 时间（秒） |
| `exec_system_time` | double precision | 执行语句的系统 CPU 时间（秒） |
| `exec_reads` | bigint | 文件系统层读取的字节数 |
| `exec_reads_blks` | bigint | 文件系统层读取的 8K 块数 |
| `exec_writes` | bigint | 文件系统层写入的字节数 |
| `exec_writes_blks` | bigint | 文件系统层写入的 8K 块数 |
| `plan_user_time` | double precision | 规划阶段的用户 CPU 时间（启用追踪时） |
| `plan_system_time` | double precision | 规划阶段的系统 CPU 时间（启用追踪时） |

**`pg_stat_kcache_detail`** -- 按查询的统计信息，包含查询文本、角色和数据库。

### 函数

```sql
-- 重置所有已收集的统计信息
SELECT pg_stat_kcache_reset();
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_stat_kcache.linux_hz` | -1 | Linux CONFIG_HZ（自动检测） |
| `pg_stat_kcache.track` | `top` | 追踪级别：`top`、`all` 或 `none` |
| `pg_stat_kcache.track_planning` | `off` | 追踪规划统计信息（PG 13+） |

### 示例

```sql
SELECT datname, exec_user_time, exec_system_time, exec_reads, exec_writes
FROM pg_stat_kcache;

SELECT query, exec_user_time, exec_system_time, exec_reads
FROM pg_stat_kcache_detail
ORDER BY exec_user_time DESC
LIMIT 10;
```
