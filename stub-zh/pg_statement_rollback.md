

## 用法

> [pg_statement_rollback: 类似 Oracle 或 DB2 的 PostgreSQL 语句级服务器端回滚](https://github.com/lzlabs/pg_statement_rollback)

在每条语句之前自动创建服务器端保存点，允许单条语句失败而不中止整个事务。

### 启用

```sql
LOAD 'pg_statement_rollback.so';
SET pg_statement_rollback.enabled TO on;
```

或在 `postgresql.conf` 中为所有会话启用：

```ini
session_preload_libraries = 'pg_statement_rollback'
pg_statement_rollback.enabled = on
```

### 基本用法

```sql
BEGIN;
CREATE TABLE test(id integer);
INSERT INTO test SELECT 1;
SELECT COUNT(*) FROM test;         -- 返回 1

INSERT INTO test SELECT 'wrong';   -- ERROR: invalid input syntax
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";  -- 仅回滚失败的语句

SELECT COUNT(*) FROM test;         -- 仍然返回 1
COMMIT;
```

没有此扩展，错误会中止整个事务，后续所有语句都会以 "current transaction is aborted" 失败。

### 配置

```sql
-- 随时在会话中启用/禁用
SET pg_statement_rollback.enabled TO off;

-- 更改保存点名称（仅超级用户）
SET pg_statement_rollback.savepoint_name TO 'my_savepoint';

-- 将保存点限制为仅写操作语句（默认：on）
SET pg_statement_rollback.enable_writeonly TO off;
```

### 关键行为

- 以最小性能开销在服务器端自动创建保存点
- 默认仅在写语句（INSERT、UPDATE、DELETE 等）后创建保存点
- 当 `enable_writeonly` 开启时，SELECT 语句不会触发自动保存点
- 客户端在处理错误时仍需调用 `ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt"`
