## 用法

来源：

- [官方文档](https://github.com/vijay750/pg_cost_guard/blob/490986b50898ff3ab1c0625c3c9032f165bee235/README.md)
- [扩展控制文件](https://github.com/vijay750/pg_cost_guard/blob/490986b50898ff3ab1c0625c3c9032f165bee235/cost_guard.control)
- [官方仓库](https://github.com/vijay750/pg_cost_guard)

`cost_guard` 通过规划器钩子拒绝估算代价超过阈值的语句。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `cost_guard`：

```sql
CREATE EXTENSION cost_guard;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Check query cost without executing
EXPLAIN (FORMAT TEXT, COSTS ON)
SELECT * FROM orders o
JOIN customers c ON o.customer_id = c.id
WHERE o.order_date > '2024-01-01';

-- View cost in JSON format for programmatic analysis
EXPLAIN (FORMAT JSON, COSTS ON)
SELECT * FROM large_table WHERE complex_condition;
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 11, 12, 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
