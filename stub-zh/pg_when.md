## 用法
> 来源: [README](https://github.com/frectonz/pg-when/blob/main/README.md) 和 [项目仓库](https://github.com/frectonz/pg-when)。

`pg-when` 是一个 PostgreSQL 扩展，用于根据自然语言短语生成时间值。
它通过同一份解析结果提供多种返回格式：`when_is`、`seconds_at`、`millis_at`、`micros_at` 和 `nanos_at`。

查询语法最多由三部分组成：

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

如果未指定时区，扩展默认使用 UTC。

### 支持的组件

`<date>` 可以是相对日期，也可以是精确日期。

README 中列出的相对日期示例包括：

- `today`
- `yesterday`
- `tomorrow`
- `next week`
- `last month`
- `this friday`
- `5 days ago`
- `in 2 years`

精确日期可以写成以下格式：

- `YYYY-MM-DD` 或 `YYYY/MM/DD`
- `DD-MM-YYYY` 或 `DD/MM/YYYY`
- `Month D, YYYY`
- `D Month YYYY`

`<time>` 也可以是相对时间或精确时间。

相对时间示例包括：

- `noon`
- `midnight`
- `morning`
- `evening`
- `next hour`
- `previous minute`
- `this hour`

精确时间可以用 12 小时制或 24 小时制表示，例如 `8:30 pm` 或 `15:45`。

时区既支持 IANA 名称，也支持 UTC 偏移量，例如 `America/New_York` 或 `UTC-08:00`。

### 示例

```sql
SELECT when_is('5 days ago at this hour in Asia/Tokyo');
SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT when_is('in 2 months at midnight in UTC-8');
SELECT when_is('last monday at 22:30');
SELECT when_is('December 31, 2026 at evening');
```

### 部署

上游 README 展示了 PostgreSQL 13 到 18 的 Docker 镜像。
这与本仓库中的包元数据一致。
