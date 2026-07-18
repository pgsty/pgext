## 用法

来源：

- [项目 README](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/README.md)
- [0.9 版安装 SQL](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/build/json_query--0.9.sql)
- [上游示例](https://github.com/dgillis/pg_json_query/blob/6aa96058553cfd01a6c9a508d4d925fcbf38a064/example.sql)

`json_query` 0.9 是一个纯 SQL 查询构建扩展，可将经过校验的 JSONB 过滤对象转换为针对已注册表或复合行类型的谓词。它适合让 SQL 函数接受由应用传入的一组可选过滤条件，从而避免维护大量近似重复的函数变体。

### 注册行类型并进行过滤

```sql
CREATE EXTENSION json_query;

CREATE TABLE customer (id integer, status text);
SELECT jq_register_row_type('customer');

SELECT *
FROM customer AS c
WHERE jq_filter(c, '{"status": "active", "id__ge": 100}'::jsonb);
```

`jq_filter(anyelement,jsonb)` 接受 `id__ge` 这类 Django 风格键；`jq_filter_raw(anyelement,jsonb)` 接受显式的字段、操作符和值对象数组。可用操作包括相等与排序、集合成员、模式匹配，以及 JSONB 包含或存在性测试。嵌套 JSONB 字段可使用 `->` 和 `->>` 路径语法。

### 注册生命周期

`jq_register_row_type(text)` 会生成类型专用的辅助函数。表定义变化后，应先执行 `jq_unregister_row_type(text)`，再重新注册。自定义列类型还可能需要 `jq_register_column_type(text)`。扩展会校验字段名和操作符名，但应用仍应限制调用者可用的过滤条件并控制查询成本。README 发布的最低要求是 PostgreSQL 9.4；已审查源码没有当前版本兼容矩阵。
