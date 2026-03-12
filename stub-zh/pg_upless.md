

## 用法

> [pg_upless: 检测无效 UPDATE](https://github.com/rodo/pg_upless)

pg_upless 检测不改变任何实际值的 UPDATE 语句（ORM 中常见）。它通过在监控表上安装触发器并收集统计信息来工作。它作为诊断工具使用，不适合长期启用。

### 开始监控

```sql
-- 监控特定表
SELECT pg_upless_start('public', 'boats');

-- 监控模式中的所有表
SELECT pg_upless_start('public');
```

### 停止监控

```sql
-- 停止监控特定表
SELECT pg_upless_stop('public', 'boats');

-- 停止监控模式中的所有表
SELECT pg_upless_stop('public');
```

### 查看统计信息

该扩展在 `pg_upless` 模式中创建两个表：

- **`pg_upless_stats`** -- 存储每个表的总更新次数与无效更新次数
- **`pg_upless_start_time`** -- 记录监控开始时间（用于速率计算）
