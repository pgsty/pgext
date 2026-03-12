


## 用法

> [explain_ui: 从 SQL 生成可视化执行计划 URL](https://github.com/davidgomes/pg-explain-ui)

explain_ui 提供一个函数，接收 SQL 查询，对其运行 `EXPLAIN`，并将计划上传到 Dalibo 的执行计划可视化工具，返回一个可分享的 URL。

### 函数

```sql
SELECT explain_ui($$
  SELECT b.title, a.name, p.name
  FROM books b
  INNER JOIN authors a ON b.author_id = a.author_id
  INNER JOIN publishers p ON b.publisher_id = p.publisher_id
  ORDER BY b.publication_date DESC
$$);

                    explain_ui
--------------------------------------------------
 https://explain.dalibo.com/plan/ccg2e5fedd913bb7
```

该函数：
1. 对提供的 SQL 查询运行 `EXPLAIN (FORMAT JSON)`
2. 将计划上传到 [explain.dalibo.com](https://explain.dalibo.com/)
3. 返回可视化计划的 URL

可视化工具基于 [pev2](https://github.com/dalibo/pev2)（PostgreSQL Explain Visualizer 2），提供查询执行计划的交互式图形化视图。
