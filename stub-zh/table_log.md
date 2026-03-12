

## 用法

> [table_log: 记录表修改日志并实现表/行级别的时间点恢复](https://github.com/df7cb/table_log)

`table_log` 扩展将对表的 INSERT、UPDATE 和 DELETE 操作记录到单独的日志表中，实现表或行级别的时间点恢复。

### 初始化日志

```sql
CREATE EXTENSION table_log;

-- 基本设置：为 'my_table' 创建日志表和触发器
-- Level 5 = 记录 trigger_id + trigger_user + trigger 列
SELECT table_log_init(5, 'my_table');

-- 指定日志模式
SELECT table_log_init(5, 'my_table', 'log_schema');

-- 完整形式，包含所有选项
SELECT table_log_init(
    5,                -- 级别: 3=最小, 4=+用户, 5=+id+用户
    'public',         -- 源模式
    'my_table',       -- 源表
    'log_schema',     -- 日志表模式
    'my_table_log',   -- 日志表名（默认: {table}_log）
    'SINGLE',         -- 分区模式: 'SINGLE' 或 'PARTITION'
    false,            -- basic_mode（更简单的触发器）
    '{INSERT, UPDATE, DELETE}'::text[]  -- 要记录的操作
);
```

### 日志表结构

日志表镜像原始表的列，加上元数据：

| 列 | 描述 |
|--------|-------------|
| `trigger_mode` | 操作类型：INSERT、UPDATE、DELETE |
| `trigger_tuple` | 元组版本：'old' 或 'new' |
| `trigger_changed` | 变更时间戳 |
| `trigger_id` | 序列 ID（级别 4+） |
| `trigger_user` | 执行变更的用户（级别 5） |

### 时间点恢复

```sql
-- 将表恢复到特定时间点
SELECT table_log_restore_table(
    'my_table',           -- 原始表名
    'my_table_log',       -- 日志表名
    'id',                 -- 主键列
    'trigger_changed',    -- 日志中的时间戳列
    'trigger_tuple',      -- 日志中的元组类型列
    '2024-01-15 10:30:00' -- 恢复到此时间戳
);
```

### 触发器函数

| 函数 | 描述 |
|----------|-------------|
| `table_log()` | 完整触发器函数，记录所有列 |
| `table_log_basic()` | 基本触发器函数，更简单的记录 |
| `table_log_restore_table(...)` | 将表状态恢复到给定时间戳 |
