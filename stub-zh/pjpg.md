## 用法

来源：

- [官方 README](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/README.md)
- [1.0 版 SQL 实现](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg--1.0.sql)
- [扩展控制文件](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg.control)

`pjpg` 是一组纯 SQL 通用工具，主要用于在命名时区中解释 `timestamp without time zone` 值，以及构造按日或按周的 `tsrange` 窗口。它还提供 `blank_to_null(text)`，用于规范化空字符串或仅含空白的字符串。

### 时区与范围流程

```sql
CREATE EXTENSION pjpg;

SELECT in_tz(timestamp '2026-07-22 00:00:00', 'UTC', 'Asia/Shanghai');
SELECT beginning_of_day(timestamp '2026-07-22 14:30:00', 'Asia/Shanghai');
SELECT day_range(timestamp '2026-07-22 14:30:00', 'Asia/Shanghai', 7);
SELECT blank_to_null('   ');
```

双参数 `in_tz(t, to_tz)` 假定 `t` 表示 UTC。`beginning_of_week_usa()` 从星期日开始，`beginning_of_week_iso()` 从星期一开始，而 `beginning_of_week()` 是 USA 约定的别名。`day_range()`、`weekusa_range()`、`weekiso_range()` 与 `week_range()` 返回调用方指定宽度的半开 `[)` 范围。

### 时间戳语义

输入输出特意不使用带时区类型。它们按调用方约定表示本地墙上时间，因此数据库不会记住所用时区。夏令时切换可能使本地时间存在歧义或根本不存在；应为每个部署时区测试两个切换方向，并优先用 `timestamptz` 存储时间点。

时区结果依赖 PostgreSQL 安装的 tzdata，并可能随时区数据库更新而变化。应验证非正范围宽度、NULL 与空白输入、ISO 和 USA 周定义要求，以及下游对排除边界的处理。当 `search_path` 未被严格控制时，应对这些通用函数名使用模式限定。
