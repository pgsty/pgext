## 用法

来源：

- [官方介绍](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/docs/source/intro.rst)
- [扩展控制文件](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/pgsfti.control)
- [扩展 SQL](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/sql/pgsfti--1.0.sql)
- [类型与关系实现](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/src/pgsfti.c)

`pgsfti` 增加了用于简单梯形模糊时间区间的 `sfti` 类型。历史或领域事件的支撑区间与核心边界不确定，而查询又需要带程度的 Allen 时间关系而不仅是精确 PostgreSQL 日期范围时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION pgsfti;

CREATE TABLE historical_event (
  event_name text PRIMARY KEY,
  happened sfti NOT NULL
);

INSERT INTO historical_event VALUES
  ('period_a', sfti_makeSFTI(1900, 1925, 1975, 1999)),
  ('period_b', sfti_fuzzify(1980, 1990, 3, 5));

SELECT event_name
FROM historical_event
WHERE happened && sfti_makeSFTI(1970, 1985);

SELECT allen_overlaps(
  sfti_makeSFTI(1900, 1925, 1975, 1999),
  sfti_makeSFTI(1970, 1980, 2000, 2010),
  0.0,
  0.0
);
```

一个 `sfti` 保存五个浮点值：支撑起点、核心起点、核心终点、支撑终点和最大隶属度。`sfti_makeSFTI` 提供年份与日期重载；`sfti_fuzzify` 用模糊边界扩展精确年份、日期或区间。`smallint`、`integer`、`date` 都有隐式转换。

关系函数包括 Allen 系列 `allen_before`、`allen_meets`、`allen_overlaps`、`allen_during`、`allen_starts`、`allen_finishes`、`allen_equals` 及其逆关系，还包括 `kvd_before`、`kvd_after`、`kvd_during`、`kvd_contains` 与 `kvd_intersects`。它们返回浮点隶属度并接受两个容差参数。布尔操作符 `<`、`>`、`=`、`&&`、`@`、`~` 提供固定阈值的快捷判断。

### 表示注意事项

日期会转换成浮点年份位置，而不是作为日历日期保存；不存在零年，所以公元前 50 年映射为 -49。转换与输出使用浮点数，不能精确保留所有日历表示。输入函数只解析五个数字，不强制支撑与核心顺序，也不检查隶属度范围，因此应在表约束或应用代码中验证 `support_start <= core_start <= core_end <= support_end` 以及适合领域的隶属度值。

SQL 定义了操作符，但没有索引操作符类，因此若不另外加入可索引的粗过滤条件，关系谓词通常需要扫描。不需要预加载。应在确切 PostgreSQL 大版本上测试 1.0；由于表依赖自定义基本类型，转储恢复期间必须保留扩展二进制。
