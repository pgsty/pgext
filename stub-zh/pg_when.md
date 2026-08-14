## 用法

来源：

- [PGXN 上的 pg_when 0.1.10](https://pgxn.org/dist/pg_when/0.1.10/)
- [pg_when 0.1.10 README](https://github.com/frectonz/pg-when/blob/0.1.10/README.md)
- [pg_when 0.1.10 Cargo 清单](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/Cargo.toml)
- [pg_when 0.1.10 control 文件](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/pg_when.control)
- [pg_when 0.1.10 导出函数源码](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_is.rs)
- [pg_when 0.1.10 相对日期实现](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_relative_date.rs)

`pg_when` 0.1.10 解析受限的自然语言日期与时间表达式，返回 PostgreSQL `timestamptz`，或按指定精度返回 Unix epoch 值。

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('5 days ago at this hour in Asia/Tokyo');
SELECT millis_at('in 2 months at midnight in UTC-8');
SELECT micros_at('December 31, 2026 at evening');
SELECT nanos_at('last monday at 22:30');
```

### 查询结构

查询可包含日期、时间和时区，并通过 `at` 与 `in` 连接：

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<date> in <timezone>');
SELECT when_is('<time>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

省略时区时，解析器使用 UTC。支持的输入包括 `tomorrow`、`last month`、`5 days ago` 等相对日期，常见数字及月份名称形式的准确日期，`noon`、`midnight`、`next hour` 等相对时间，时钟时间、IANA 时区名与 UTC 偏移量。

### 函数索引

- `when_is(text)` 返回 `timestamptz`。
- `seconds_at(text)` 返回 Unix epoch 秒数。
- `millis_at(text)` 返回 Unix epoch 毫秒数。
- `micros_at(text)` 返回 Unix epoch 微秒数。
- `nanos_at(text)` 返回 Unix epoch 纳秒数。

### 兼容性与边界

- 解析器实现的是文档中定义的语法，并非通用自然语言解释器。
- 上游 0.1.10 提供 PostgreSQL 13–18 的构建特性并固定使用 pgrx 0.18.1；Pigsty 软件包覆盖 PostgreSQL 14–18，并应用锁定依赖的 pgrx 0.19.1 兼容更新。
- `pg_when` 不可重定位，其 control 文件要求超级用户执行 `CREATE EXTENSION`。
- 非法文本会触发错误。这五个函数都声明为 `STRICT`，因此空值输入返回空值；当 epoch 纳秒数无法放入 `bigint` 时，`nanos_at(text)` 也会报错。
- 0.1.10 的 SQL 函数声明为 `IMMUTABLE`，但 `now`、`tomorrow`、`5 days ago` 等相对表达式会读取当前时钟。不要把相对输入调用用于表达式索引或生成列，也不要假定它们会在缓存计划中重新求值；只有完整指定日期、时间与时区的输入才与当前时间无关。
