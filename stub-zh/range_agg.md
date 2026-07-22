## 用法

来源：

- [官方 README](https://github.com/pjungwir/range_agg/blob/a9e825be48954ff1d83b619f7b7be7eaf29fc3be/README.md)
- [官方 1.2.1 版 SQL](https://github.com/pjungwir/range_agg/blob/a9e825be48954ff1d83b619f7b7be7eaf29fc3be/range_agg--1.2.1.sql)
- [官方 PGXN 版本](https://pgxn.org/dist/range_agg/1.2.1/)
- [PostgreSQL 内置聚合文档](https://www.postgresql.org/docs/current/functions-aggregate.html)

`range_agg` 版本 `1.2.1` 合并相邻或重叠的 PostgreSQL 范围，并可拒绝间隙或重叠。它早于 PostgreSQL 同名的内置多范围聚合，因此现代安装必须处理名称解析与不同返回类型。

### 核心流程

在适合使用该扩展的服务器上，按房间合并连续预订：

```sql
CREATE EXTENSION range_agg;

SELECT room_id, range_agg(booked_during)
FROM reservations
GROUP BY room_id;
```

单参数形式返回一个范围，遇到间隙或重叠时抛出错误。要在允许间隙的同时合并范围，可使用双参数形式，并展开返回的数组：

```sql
SELECT room_id, merged_range
FROM (
  SELECT room_id, range_agg(booked_during, true) AS merged_ranges
  FROM reservations
  GROUP BY room_id
) AS grouped
CROSS JOIN LATERAL unnest(grouped.merged_ranges) AS merged_range;
```

### 聚合形式

- `range_agg(anyrange)` 返回一个范围，并拒绝间隙和重叠。
- `range_agg(range, permit_gaps)` 返回数组，参数为 true 时允许间隙，但仍拒绝重叠。
- `range_agg(range, permit_gaps, permit_overlaps)` 返回数组，可分别控制两种条件。
- 双参数和三参数形式为 `int4range`、`int8range`、`numrange`、`tsrange`、`tstzrange` 与 `daterange` 声明。自定义范围类型需要配套的转换/终结函数与聚合声明。

### 运维注意事项

PostgreSQL 14 及以上版本提供返回多范围的 `pg_catalog.range_agg(anyrange)`。该内置函数语义不同，通常会在未限定名称查找中优先匹配，而本扩展的单参数聚合返回单个范围。应使用 `regprocedure` 确认实际选择的函数；如果有意共存，应限定模式；当前 PostgreSQL 应优先使用仍在维护的内置实现。当范围重叠或包含不允许的间隙时，聚合错误会中止语句。依赖合并结果前，应测试空分组、无限边界、离散与连续范围，以及自定义规范化行为。
