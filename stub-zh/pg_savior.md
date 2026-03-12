

## 用法

> [pg_savior: 防止误操作的 Postgres 扩展](https://github.com/viggy28/pg_savior)

`pg_savior` 扩展拦截查询执行以防止意外的数据删除。它通过钩入执行器来检测危险的 DELETE 操作并阻止它们。

### 工作原理

当处理 DELETE 语句时，`pg_savior` 检查：

- DELETE 命令中**缺失 WHERE 子句**
- DELETE 查询计划中的**索引扫描操作**
- DELETE 操作中涉及 CTE 和子查询的**复杂查询**

当检测到有风险的 DELETE 时，扩展会阻止执行并返回一条信息消息，影响行数为零。

### 示例

```sql
CREATE EXTENSION pg_savior;

-- 尝试不带 WHERE 子句的 DELETE
DELETE FROM emp;
-- INFO:  pg_savior: DELETE statement detected
-- INFO:  pg_savior: WHERE clause is mandatory for a DELETE statement
-- DELETE 0  （无行受影响，数据已保留）

-- 带 WHERE 子句的正常 DELETE 按预期工作
DELETE FROM emp WHERE id = 42;
-- DELETE 1
```

### 注意事项

- 该扩展通过 PostgreSQL 执行器钩子运行，无需更改应用程序代码
- 当前仅拦截 DELETE 语句；UPDATE 操作不受影响
- 计划中的功能包括阻止不带 `CONCURRENTLY` 的 `CREATE INDEX` 和阻止 `DROP DATABASE`
