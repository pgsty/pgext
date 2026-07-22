## 用法

来源：

- [官方 README](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/README.md)
- [官方扩展 SQL](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/jalali_utils--0.0.1.sql)
- [官方 C 实现](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/jalali_utils.c)

`jalali_utils` 将 PostgreSQL 公历日期与时间戳转换成太阳回历（Jalali）文本或整数日期字段。时间戳重载始终在 `Asia/Tehran` 中解释瞬时时间，不使用会话的 `TimeZone`。

### 核心流程

```sql
CREATE EXTENSION jalali_utils;

SELECT format_jalali(timestamptz '2019-07-07 14:10:52.84937+04:30');
SELECT format_jalali(date '2021-03-20');
SELECT jalali_part('year', date '2021-03-20');
SELECT jalali_part('doy', timestamptz '2019-07-07 14:10:52.84937+04:30');
```

可选标志为 true 时，`format_jalali(timestamptz, boolean)` 返回 `YYYY/MM/DD HH:MM:SS`；为 false 时只返回日期文本。`format_jalali(date)` 始终只返回日期文本。

### 函数索引

- `format_jalali(timestamptz, boolean DEFAULT true)` 以德黑兰民用时间格式化一个瞬时点。
- `format_jalali(date)` 转换一个公历日期。
- `jalali_part(text, timestamptz)` 从德黑兰时间的瞬时点提取 Jalali 字段。
- `jalali_part(text, date)` 从日期中提取 Jalali 字段。

接受的字段包括 `year`、`month`、`day`、`hour`、`minute`、`second`、`doy`、`dow`、`quarter`、`decade` 和 `century`。`dow` 以星期六为零。

### 注意事项

版本 `0.0.1` 将格式化结果作为 `text` 返回，而不是可排序或经过验证的 Jalali 数据类型。排序和比较应基于原始 PostgreSQL 日期或时间戳，只在展示边界做格式化。硬编码的德黑兰时区意味着时间戳结果相对于 UTC 或会话时区可能跨越日历日期。把转换用于金融或法律流程前，应在实际二进制上验证闰日、夏令时历史、负数/无穷日期和受支持的 PostgreSQL 版本。
