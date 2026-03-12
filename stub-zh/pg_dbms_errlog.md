

## 用法

> [pg_dbms_errlog: 模拟 Oracle DBMS_ERRLOG 模块，将 DML 错误记录到专用表中](https://github.com/HexaCluster/pg_dbms_errlog)

使 DML 操作在遇到错误后可以继续执行，将失败记录到错误表中，而非中止事务。

### 启用

在 `postgresql.conf` 中添加到 `shared_preload_libraries`：

```ini
shared_preload_libraries = 'pg_dbms_errlog'
```

```sql
CREATE EXTENSION pg_dbms_errlog;
LOAD 'pg_dbms_errlog';
```

### 创建错误日志表

```sql
BEGIN;
CALL dbms_errlog.create_error_log('employees');
END;
-- 创建表 "ERR$_employees"，包含错误日志列

-- 使用自定义名称和模式：
BEGIN;
CALL dbms_errlog.create_error_log('hr.employees', '"ERRORS"."ERR$_EMPTABLE"');
END;
```

### 配置

```sql
SET pg_dbms_errlog.enabled TO true;       -- 启用错误日志
SET pg_dbms_errlog.query_tag TO 'daily_load';  -- 用于标识语句的标签
SET pg_dbms_errlog.reject_limit TO 10;    -- 回滚前的最大错误数（-1=无限）
SET pg_dbms_errlog.synchronous TO 'transaction'; -- 'transaction'、'query' 或 'off'
SET pg_dbms_errlog.no_client_error TO true;      -- 抑制客户端错误消息
```

### 与 pg_statement_rollback 配合使用

```sql
LOAD 'pg_dbms_errlog';
LOAD 'pg_statement_rollback';

CREATE TABLE hr.raises (emp_id integer, sal integer CHECK(sal > 8000));

BEGIN;
CALL dbms_errlog.create_error_log('hr.raises');
END;

SET pg_dbms_errlog.query_tag TO 'daily_load';
SET pg_dbms_errlog.reject_limit TO 10;
SET pg_dbms_errlog.enabled TO true;

BEGIN;
SET pg_statement_rollback.enabled TO on;
INSERT INTO hr.raises VALUES (145, 15400);  -- 成功
INSERT INTO hr.raises VALUES (161, 7700);   -- 失败（已记录）
ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
INSERT INTO hr.raises VALUES (175, 9680);   -- 成功
COMMIT;
```

### 查看错误日志

```sql
SELECT * FROM "ERR$_raises";
-- pg_err_number$  | 23514
-- pg_err_mesg$    | new row for relation "raises" violates check constraint
-- pg_err_optyp$   | I
-- pg_err_tag$     | daily_load
-- pg_err_query$   | INSERT INTO hr.raises VALUES (161, 7700);
```

### 刷新排队错误

```sql
SELECT dbms_errlog.publish_queue();
```
