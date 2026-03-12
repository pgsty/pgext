

## 用法

> [pg_auditor: 带闪回功能的数据变更审计](https://github.com/kouber/pg_auditor)

`pg_auditor` 记录指定表上的每一次数据修改（INSERT、UPDATE、DELETE），并允许部分或完整的事务闪回。

```sql
CREATE EXTENSION pg_auditor CASCADE;  -- 同时安装 hstore
```

### 审计控制

```sql
-- 开始审计一个表（默认审计所有 DML）
SELECT auditor.attach('fruit');

-- 仅审计特定操作
SELECT auditor.attach('fruit', ARRAY['INSERT', 'UPDATE']);

-- 仅审计特定列
SELECT auditor.attach('fruit', ARRAY['INSERT', 'UPDATE', 'DELETE'], ARRAY['name', 'weight']);

-- 停止审计
SELECT auditor.detach('fruit');

-- 管理单个语句/列
SELECT auditor.attach_statement('fruit', 'DELETE');
SELECT auditor.detach_statement('fruit', 'DELETE');
SELECT auditor.attach_column('fruit', 'weight');
SELECT auditor.detach_column('fruit', 'weight');

-- 防止 TRUNCATE
SELECT auditor.forbid_truncate('fruit');
```

### 查看审计日志

```sql
SELECT transaction_id, operation, old_rec, new_rec FROM auditor.log;
```

### 闪回函数

```sql
-- 撤销当前会话中最后 N 个事务
SELECT auditor.undo();          -- 撤销最后1个
SELECT auditor.undo(3);         -- 撤销最后3个
SELECT auditor.undo(1, true);   -- 覆盖其他会话

-- 取消特定事务
SELECT auditor.cancel(5557);

-- 将数据恢复到特定事务或时间戳
SELECT auditor.flashback(5556);
SELECT auditor.flashback('2021-02-08 14:40:00'::timestamp);
```

### 列演化追踪

```sql
SELECT * FROM auditor.evolution('fruit'::regclass, 'weight', 'orange'::text);
-- 显示给定主键的某列值的完整历史
```
