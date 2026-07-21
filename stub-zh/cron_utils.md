## 用法

来源：

- [pg_cron_utils 0.1.0 README](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/README.md)
- [Extension control file](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils.control)
- [SQL definitions](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils--0.1.0.sql)

`cron_utils` 解析五字段 cron 表达式并在 PostgreSQL 中计算触发时间戳。它是一个调度工具，而不是作业执行器：使用它可以预览、验证或查询一个计划，并单独使用 `pg_cron` 等调度程序来执行工作。

### 核心流程

```sql
CREATE EXTENSION cron_utils;

-- First trigger at or after the supplied time.
SELECT cron_first_trigger('0 9 * * 1-5', now());

-- Last trigger before the supplied time (strict is true by default).
SELECT cron_last_trigger('0 9 * * 1-5', now());

-- Next five hourly triggers.
SELECT *
FROM cron_iterate_n('0 * * * *', now(), false, 'next', 5);
```

要检查一个限定窗口：

```sql
SELECT *
FROM cron_first_last_triggers(
  '0 0 * * *',
  date_trunc('month', now()),
  date_trunc('month', now()) + interval '1 month'
);
```

当表达式在窗口中没有触发时，返回的边界可以是 `NULL`。

### 重要对象

- `cron_parts` 是分钟、小时、日、月和周字段的解析表示。
- `parse_cron(text)` 解析 `*`、列表、范围以及步进语法。
- `cron_first_trigger(expr, base_time, strict)` 向前搜索。当 `strict = true` 时，会跳过在 `base_time` 精确匹配的情况。
- `cron_last_trigger(expr, base_time, strict)` 向后搜索，默认为严格匹配。
- `cron_first_last_triggers(expr, start_time, end_time)` 返回窗口中的第一个和最后一个匹配项。
- `cron_iterate_n(expr, base_time, strict, direction, max_matches)` 按 `next` 或 `prev` 方向返回连续的匹配项。

### 语义与注意事项

表达式使用标准的五个字段 `minute hour day month dow`；秒不被接受。周字段使用 `1 = Monday` 到 `7 = Sunday` 的表示方式。结果为 `timestamptz`，因此会话时区会影响显示的本地时间，并且对于本地时间计划应测试夏令时转换。

该扩展是纯 SQL/PL/pgSQL 实现的，可重定位且没有依赖于 `pg_cron`。其函数声明为不可变且并行安全。0.1.0 版本在控制元数据中标记为 `unstable`，因此在将其嵌入关键调度程序之前，请锁定行为并重新测试边缘情况。
