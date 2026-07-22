## 用法

来源：

- [v1.1.3 官方 README](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/README.md)
- [v1.1.3 官方扩展 SQL](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/sql/pg_ethiopian_calendar--1.1.sql)
- [v1.1.3 官方 control 文件](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/pg_ethiopian_calendar.control)

`pg_ethiopian_calendar` 用于在 PostgreSQL 公历时间戳与埃塞俄比亚历日期之间转换。扩展 SQL 版本是 `1.1`，本次核验的上游发行标签是 `v1.1.3`。适合需要埃塞俄比亚 13 月民用历法的应用，同时仍以普通 PostgreSQL 时间戳作为存储与交换边界。

### 核心流程

```sql
CREATE EXTENSION pg_ethiopian_calendar;

SELECT to_ethiopian_date(TIMESTAMP '2026-01-01 00:00:00');
SELECT from_ethiopian_date('2018-04-23');

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  occurred_at timestamp NOT NULL,
  ethiopian_at timestamp
    GENERATED ALWAYS AS (to_ethiopian_timestamp(occurred_at)) STORED
);
```

文本转换使用埃塞俄比亚历 `YYYY-MM-DD` 格式。应用应校验输入，并明确决定每个接口使用文本结果还是时间戳结果。

### 函数

- `to_ethiopian_date(timestamp)` 返回埃塞俄比亚历日期文本；`from_ethiopian_date(text)` 将该文本解析回公历时间戳。
- `to_ethiopian_datetime(timestamp)` 返回 `timestamptz`，而 `to_ethiopian_timestamp(timestamp)` 保留时间戳形式的结果。
- `current_ethiopian_date()` 返回当前埃塞俄比亚历日期。
- `pg_ethiopian_to_date`、`pg_ethiopian_from_date`、`pg_ethiopian_to_datetime` 与 `pg_ethiopian_to_timestamp` 是相应转换函数的别名。

转换函数声明为 `IMMUTABLE` 与 `STRICT`；`current_ethiopian_date()` 声明为 `STABLE`。因此确定性转换可用于生成表达式和索引，但仍受 PostgreSQL 类型与时区语义约束。

### 运维注意事项

`to_ethiopian_datetime(timestamp)` 的输入是 `timestamp without time zone`，结果却是 `timestamptz`；依赖这种形式前，应测试应用使用的会话时区。不希望发生隐式时区解释时，优先使用保持 timestamp 形式的函数。

仓库还包含独立的 `sql/ethiopian_calendar_plpgsql.sql` 实现，其 API 不同。扩展的 control/SQL 安装路径不会安装它，所以不能假定执行 `CREATE EXTENSION` 后就存在该文件中的无参数辅助函数。

上游记录的构建范围为 PostgreSQL 11–17。应针对目标 PostgreSQL 版本核验实际软件包；对于更新版本尤其如此，并在生产使用前测试历法边界与非法输入。
