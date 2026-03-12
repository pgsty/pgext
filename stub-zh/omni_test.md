


## 用法

> [omni_test: 测试框架](https://docs.omnigres.org/omni_test/guide/)

`omni_test` 扩展支持在 PostgreSQL 数据库内执行测试。它是一个模板化扩展。

### 初始化

```sql
CREATE DATABASE myapp_test;
UPDATE pg_database SET datistemplate = true WHERE datname = 'myapp_test';
```

### 编写测试

**函数测试：**

```sql
CREATE FUNCTION my_test() RETURNS omni_test.test
    LANGUAGE plpgsql AS $$
BEGIN
    -- 测试断言
END;
$$;
```

**过程测试**（用于需要 COMMIT/ROLLBACK 的非原子测试）：

```sql
CREATE PROCEDURE my_test(INOUT omni_test.test)
    LANGUAGE plpgsql AS $$
BEGIN
    -- 测试逻辑
END;
$$;
```

### 测试设置

```sql
CREATE PROCEDURE tx_test(INOUT test omni_test.test)
    SET omni_test.transaction_isolation = serializable
    LANGUAGE plpgsql AS $$ ... $$;
```

### 运行测试

```sql
SELECT * FROM omni_test.run_tests('myapp_test');
```

返回：`name`、`description`、`start_time`、`end_time`、`error_message`。

### 过滤测试

```sql
-- 排除标记为 @slow 的测试：
SELECT * FROM omni_test.run_tests('myapp_test', filter => '^(?!.*@slow).*');
```

每个测试在模板数据库的全新副本中运行。
