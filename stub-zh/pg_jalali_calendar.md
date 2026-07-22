## 用法

来源：

- [官方 Rust 实现](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/src/lib.rs)
- [官方 Cargo 清单](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/Cargo.toml)
- [扩展 control 文件](https://github.com/alisol911/pg_jalali_calendar/blob/9682677a9d23b357b5a25db605c7ead5eb1abf50/pg_jalali_calendar.control)

`pg_jalali_calendar` 提供基于文本的波斯（Jalali）历转换与计算，并附带两个波斯语数字辅助函数。其日期 API 接收和返回字符串，而不是 PostgreSQL `date` 或 `timestamp` 值。

### 核心流程

创建扩展后，可在 Jalali `YYYY/MM/DD` 字符串与 Gregorian `YYYY-MM-DD` 字符串之间转换；对于不可信输入，应先校验，再调用那些会因无效日期报错的函数。

```sql
CREATE EXTENSION pg_jalali_calendar;

SELECT jalali_date_to_gregorian('1404/01/01');
SELECT gregorian_date_to_jalali('2025-03-21');
SELECT jalali_date_is_valid('1403/12/30');
```

日期算术返回 Jalali 文本。天数加法接受正负值，但 `0.1.0` 版的月份加法只接受正值。

```sql
SELECT jalali_date_add_days('1404/01/01', -10);
SELECT jalali_date_add_months('1404/01/31', 7);
SELECT jalali_date_diff('1404/01/01', '1404/02/01');
```

### 重要对象

- `jalali_date_to_gregorian(text)`、`gregorian_date_to_jalali(text)`：历法转换。
- `jalali_date_add_days(text, integer)`、`jalali_date_add_months(text, integer)`：日期算术。
- `jalali_date_diff(text, text)`、`jalali_date_diff_with_addition(text, text, integer)`：带符号的天数差。
- `jalali_date_now()`：按 UTC 日期计算当前 Jalali 日期。
- `jalali_date_is_valid(text)`、`jalali_date_is_leap_year(text)`：有效性与闰年检查。
- `jalali_date_period_state(text, integer)`：针对源码定义的月度周期边界返回 `Start`、`End`、`Middle` 或 `Unknown`。
- `number_to_farsi_word(bigint)`、`farsi_word_add_suffix(text)`：波斯语数字词与序数后缀。

### 语义与兼容性

月份加法遇到目标月份不存在的日期时，会收敛到该月最后一个有效日。本版本拒绝零和负的月份数量。大多数函数遇到格式错误或越界文本时会以 PostgreSQL 错误结束；`jalali_date_is_valid(text)` 是筛查输入的安全谓词。`jalali_date_now()` 使用 UTC，而不是会话时区。

固定的 `0.1.0` Cargo 清单默认只启用 pgrx PostgreSQL 18 特性；control 文件不可迁移，并同时标记为需要超级用户和 trusted。应在目标平台确认实际打包构建与权限。如果数值需要参与 PostgreSQL 索引、范围查询、时区或 interval 逻辑，应在 Jalali 展示字符串之外同时保存原生 Gregorian `date`。
