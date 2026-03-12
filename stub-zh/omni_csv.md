


## 用法

> [omni_csv: CSV 工具箱](https://docs.omnigres.org/omni_csv/)

`omni_csv` 扩展提供 CSV 解析、检查和生成函数。

### 检查 CSV 结构

```sql
SELECT * FROM omni_csv.csv_info(E'name,age,city\nJohn,25,NYC\nJane,30,LA');
-- 返回: name, age, city（每列一行）
```

### 将 CSV 解析为记录

```sql
SELECT * FROM omni_csv.parse(E'name,age,city\nJohn,25,NYC\nJane,30,LA')
    AS t(name text, age text, city text);
```

需要使用 `AS t(...)` 子句指定列结构，并应与 CSV 结构匹配。

### 从查询结果生成 CSV

```sql
SELECT omni_csv.csv_agg(t)
FROM (SELECT name, age, city FROM employees ORDER BY name) t;
```

返回带标题和正确转义值的 CSV 文本。空结果集返回空字符串。

### 限制

- 大型 CSV 字符串完全在内存中处理（无流式处理）
- 列类型必须在 `AS t(...)` 子句中声明
- 对于重复访问，建议将解析后的数据物化
