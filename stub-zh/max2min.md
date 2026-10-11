## 用法

来源：

- [README.md](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/README.md)
- [max2min--1.0.sql](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/max2min--1.0.sql)
- [LICENSE](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/LICENSE)
- [max2min.control](https://github.com/rdnkrkmz/max2min/blob/6de4805841dfd3797c3f967acfd2939fecedb6c9/max2min.control)

`max2min` 1.0 提供 `max_to_min()` 聚合函数，将最大值和最小值合并为一个文本结果，并可指定分隔符。

### 核心用法

```sql
CREATE EXTENSION max2min;
SELECT max_to_min(v) FROM (VALUES (5), (3), (10), (NULL::integer)) AS t(v);
SELECT max_to_min(v, ' .. ') FROM (VALUES (5), (3), (10)) AS t(v);
```

### 运行边界

版本化 SQL 为数值、整数、浮点、字符和日期时间类型提供重载。聚合函数处理 SQL NULL 输入，大小关系依据输入类型，而非格式化后的文本。结果是展示用文本，不是带类型的范围值。这是纯 SQL 扩展，无共享库和预加载要求。control 未将安装标记为可信，上游也未公布 PostgreSQL 主版本兼容矩阵。
