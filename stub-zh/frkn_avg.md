## 用法

来源：

- [0.0.1 版本 SQL 定义](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg--0.0.1.sql)
- [C 实现](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg.c)
- [扩展控制文件](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg.control)

`frkn_avg` 定义了由自定义 C 类型 `my_state` 支撑的 `my_avg(bigint)` 与 `my_avg(float)` 聚合。两者都用单精度字段累计总和与计数。

```sql
CREATE EXTENSION frkn_avg;
SELECT my_avg(v::bigint), my_avg(v::float)
FROM generate_series(1, 10) AS g(v);
```

该实现具有实验性质并存在明显正确性风险：SQL `bigint` 转移值由 32 位 C 访问器读取，浮点输入会被降精度，数值或行数增大时状态精度也会下降；声明的内部长度还与两个 float 组成的 C 结构不一致。不要把它开放给不可信工作负载，也不要用于财务或精确分析；应先用代表性数据与 PostgreSQL `avg` 对比结果。
