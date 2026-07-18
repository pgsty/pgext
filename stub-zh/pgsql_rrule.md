## 用法

来源：

- [固定提交的 recurrence-rule SQL 实现](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/sql/pgsql_rrule.sql)
- [固定提交的扩展 control 文件](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/pgsql_rrule.control)
- [固定提交的上游 README](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/README.md)

`pgsql_rrule` 是一套纯 SQL 数据类型与函数，用于解析和展开 RFC 5545 recurrence rule。它定义了复合类型 `rrule`、frequency 与 weekday 类型、解析器 `rrule(text)`、返回集合的 `rrule_expand(...)`，以及兼容函数 `get_occurrences(...)`。

```sql
CREATE EXTENSION pgsql_rrule;

SELECT *
FROM rrule_expand(
  rrule('FREQ=WEEKLY;BYDAY=MO,WE'),
  DATE '2026-07-01',
  DATE '2026-07-31'
);
```

解析器会校验字段组合，但已复核的展开代码有意保留了未实现范围。它会对不支持的 frequency 报错，并明确拒绝 weekly `BYSETPOS` 与 weekly `INTERVAL > 1`；应测试应用使用的每种 recurrence 形态，包括 inclusivity、time zone、fractional second、畸形输入和超大或无边界结果集。

control 文件记录版本 `0.0.4`，它从仓库的未版本化 SQL 源生成，而项目元数据把该版本线标为 unstable。应固定准确 commit，不要只依赖版本标签；将结果与受维护的 RFC 5545 实现比对，并在应用层限制展开上限。
