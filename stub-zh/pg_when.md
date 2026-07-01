


## 用法

来源：[README](https://github.com/frectonz/pg-when/blob/main/README.md), [Cargo.toml version 0.1.9](https://github.com/frectonz/pg-when/blob/main/Cargo.toml), [META.json](https://github.com/frectonz/pg-when/blob/main/META.json)

`pg-when` 解析受限的自然语言时间表达式，并返回 PostgreSQL timestamp with time zone，或以不同精度返回 Unix epoch 值。

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('next friday at 8:00 pm in America/New_York');
SELECT millis_at('next friday at 8:00 pm in America/New_York');
SELECT micros_at('next friday at 8:00 pm in America/New_York');
SELECT nanos_at('next friday at 8:00 pm in America/New_York');
```

### 支持的查询形状

解析器接受最多三个部分：

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

如果没有提供 timezone，上游说明默认是 UTC。

### 常见输入

- relative dates：`today`、`tomorrow`、`last month`、`this friday`、`5 days ago`、`in 2 years`
- exact dates：`YYYY-MM-DD`、`DD/MM/YYYY`、`January 10, 2004`、`10 Jan 2004`
- relative times：`noon`、`midnight`、`morning`、`evening`、`next hour`
- exact times：`8:30 pm`、`15:45`
- time zones：`America/New_York`、`Europe/London`、`UTC-08:00`、`UTC+05:30`

### 示例

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('December 31, 2026 at evening');
```

### 注意事项

- 扩展面向上面记录的 grammar，不是任意英文解析器。
- 上游仍列出 PostgreSQL 13 到 18 的源码/runtime 支持和 Docker image 示例，但本仓库 package matrix 仅为 PostgreSQL 14 到 18；不要假设 Pigsty 为 PostgreSQL 13 提供包。
- 上游 `Cargo.toml` 当前固定 `pgrx` 0.15.0；本仓库 package metadata 记录了手工升级到 `pgrx` 0.17.0。
