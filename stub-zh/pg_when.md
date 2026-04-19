## 用法

来源: [official README](https://github.com/frectonz/pg-when/blob/main/README.md), [official repo](https://github.com/frectonz/pg-when)

`pg-when` 解析受限的自然语言时间表达式，并返回 PostgreSQL `timestamp with time zone`，或返回不同精度的 Unix epoch 值。

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('next friday at 8:00 pm in America/New_York');
SELECT millis_at('next friday at 8:00 pm in America/New_York');
SELECT micros_at('next friday at 8:00 pm in America/New_York');
SELECT nanos_at('next friday at 8:00 pm in America/New_York');
```

### 支持的查询形态

解析器最多接受三部分：

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

如果没有提供时区，上游说明默认使用 UTC。

### 常见输入

- 相对日期：`today`, `tomorrow`, `last month`, `this friday`, `5 days ago`, `in 2 years`
- 精确日期：`YYYY-MM-DD`, `DD/MM/YYYY`, `January 10, 2004`, `10 Jan 2004`
- 相对时间：`noon`, `midnight`, `morning`, `evening`, `next hour`
- 精确时间：`8:30 pm`, `15:45`
- 时区：`America/New_York`, `Europe/London`, `UTC-08:00`, `UTC+05:30`

### 示例

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('December 31, 2026 at evening');
```

### 注意事项

- 该扩展面向上面记录的语法，而不是任意英文表达。
- 上游提供了 PostgreSQL 13 到 18 的现成 Docker 镜像，但这个 stub 应聚焦 SQL 用法，而不是容器部署。
