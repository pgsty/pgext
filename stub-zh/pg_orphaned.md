

## 用法

> [pg_orphaned: 处理孤立文件](https://github.com/bdrouvot/pg_orphaned)

pg_orphaned 提供检测和管理 PostgreSQL 中孤立数据文件的函数。它通过使用脏快照处理进行中的事务可能导致误报的边角情况。

### 列出孤立文件

```sql
-- 列出孤立文件（默认：超过 1 天的标记为 "older"）
SELECT * FROM pg_list_orphaned();

-- 自定义时间阈值
SELECT * FROM pg_list_orphaned('10 seconds');
SELECT * FROM pg_list_orphaned('1 minute');
```

返回：`dbname`、`path`、`name`、`size`、`mod_time`、`relfilenode`、`reloid`、`older`（布尔值）。

### 将孤立文件移至备份

```sql
-- 将超过阈值的文件移动到 orphaned_backup 目录
SELECT pg_move_orphaned('1 minute');
```

### 列出已移动的文件

```sql
SELECT * FROM pg_list_orphaned_moved();
```

### 将文件移回（如果仍然是孤立的）

```sql
SELECT pg_move_back_orphaned();
```

### 删除已移动的文件

```sql
SELECT pg_remove_moved_orphaned();
```

### 典型工作流程

```sql
-- 1. 检查孤立文件
SELECT * FROM pg_list_orphaned('1 minute');

-- 2. 将它们移到备份（仅超过阈值的）
SELECT pg_move_orphaned('1 minute');

-- 3. 验证移动了什么
SELECT * FROM pg_list_orphaned_moved();

-- 4. 确认后删除备份文件
SELECT pg_remove_moved_orphaned();
```

注意：函数操作的是你当前连接的数据库中的孤立文件。在移动或删除文件前请务必仔细检查。
