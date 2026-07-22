## 用法

来源：

- [Official extension control file](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/explanation.control)
- [Official PGXN distribution page](https://pgxn.org/dist/explanation/)
- [Official upstream README](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/README.md)

`explanation` 0.3.0 把 PostgreSQL `EXPLAIN` 输出解析成关系行：每个计划节点一行，子节点通过父标识符关联。当邻接树结构比格式化计划文本更便于查询时，它可用于在 SQL 内分析执行计划。

### 核心流程

```sql
CREATE EXTENSION explanation;

SELECT *
FROM explanation('SELECT * FROM orders WHERE customer_id = 42');
```

向 `explanation(text)` 传入一条查询字符串。函数运行普通 `EXPLAIN`，解析其 XML 表示，并以表格形式返回计划节点，供过滤、连接或可视化使用。

### 要求与注意事项

上游版本要求 PostgreSQL 9.0 或更高版本且编译时启用 XML；其代码早于现代 JSON 格式计划工具。它不会运行 `EXPLAIN ANALYZE`，因此结果是优化器估算而非实测执行数据。只接受可信调用方提供的查询文本，并针对实际 PostgreSQL 大版本测试解析器，因为计划字段会随版本变化。
