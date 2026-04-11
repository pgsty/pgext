
## 用法

> [pg_store_plans: 执行计划存储与统计](https://github.com/ossc-db/pg_store_plans)

`pg_store_plans` 会追踪 SQL 语句的执行计划统计信息，为 `pg_stat_statements` 补充计划级别的数据。在 PostgreSQL 14+ 上可通过 `queryid` 关联两者。

### 查看计划统计

```sql
-- 查看追踪到的计划及统计
SELECT queryid, planid, plan, calls, total_time, rows
FROM pg_store_plans
ORDER BY total_time DESC;

-- 检查模块状态
SELECT * FROM pg_store_plans_info;
```

### 关键视图列

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `queryid` | bigint | 查询 ID，可与 `pg_stat_statements` 关联 |
| `planid` | bigint | 计划哈希值 |
| `plan` | text | 代表性计划文本 |
| `calls` | bigint | 执行次数 |
| `total_time` | double precision | 总执行时间（毫秒） |
| `rows` | bigint | 检索或影响的总行数 |
| `shared_blks_hit` | bigint | 共享缓冲区命中次数 |
| `shared_blks_read` | bigint | 共享块读取次数 |
| `first_call` | timestamptz | 首次执行时间 |
| `last_call` | timestamptz | 最后执行时间 |

### 函数

```sql
-- 重置所有统计信息，仅超级用户可用
SELECT pg_store_plans_reset();

-- 转换计划格式
SELECT pg_store_plans_textplan(plan);   -- 转为文本
SELECT pg_store_plans_jsonplan(plan);   -- 转为 JSON
SELECT pg_store_plans_xmlplan(plan);    -- 转为 XML
SELECT pg_store_plans_yamlplan(plan);   -- 转为 YAML

-- 计算查询哈希
SELECT pg_store_hash_query('SELECT 1');
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_store_plans.max` | 1000 | 最大追踪计划数，仅在服务器启动时生效 |
| `pg_store_plans.track` | `top` | `top`、`all`、`verbose`、`none` |
| `pg_store_plans.plan_format` | `text` | `text`、`json`、`xml`、`yaml`、`raw` |
| `pg_store_plans.min_duration` | 0 | 最小追踪执行时间（毫秒） |
| `pg_store_plans.log_analyze` | `off` | 包含 EXPLAIN ANALYZE 输出 |
| `pg_store_plans.log_buffers` | `off` | 包含缓冲区统计 |
| `pg_store_plans.log_timing` | `true` | 记录实际耗时 |
| `pg_store_plans.plan_storage` | `file` | 存储方式：`file` 或 `shmem` |
| `pg_store_plans.max_plan_length` | 5000 | 计划文本最大字节数 |
| `pg_store_plans.save` | `on` | 跨重启持久化统计 |
