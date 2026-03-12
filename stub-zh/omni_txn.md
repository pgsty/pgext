


## 用法

> [omni_txn: 事务管理](https://docs.omnigres.org/omni_txn/linearize/)

`omni_txn` 扩展提供事务线性化功能，在标准 PostgreSQL 隔离级别之上提供更强的一致性保证。

### 启用线性化

```sql
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;
SELECT omni_txn.linearize();
-- 执行操作；冲突会引发序列化错误
COMMIT;
```

### 检查线性化状态

```sql
SELECT omni_txn.linearized();  -- 返回布尔值
```

### 行为

线性化确保每个操作看起来都在与实时顺序一致的原子顺序中发生。该扩展拦截变更语句（INSERT、UPDATE、DELETE、MERGE）并强制执行：

- 跨可序列化线性化事务的写后读冲突检测
- 提交后读冲突检测
- 快照冲突检测

检测到冲突时抛出序列化错误异常。可能出现误报；请配合重试机制使用（兼容 `omni_txn.retry`）。
