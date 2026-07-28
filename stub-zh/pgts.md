## 用法

来源：

- [官方上游 README](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/README.md)
- [官方扩展控制文件 (pgts.control)](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/src/pgts.control)
- [官方扩展 SQL (pgts--0.1.0.sql)](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/src/pgts--0.1.0.sql)

`pgts` — time encode for PostgreSQL。使用它来进行相应的调度、时间或时间序列工作流。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION pgts;

select
  hostname,
  unnest(    ts.timestamp_decode(ctime)   )                 as ctime
from
  x
group by
  hostname,
  ctime
order by
  ctime;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `ts.f8_decode(v bytea)` 是一个扩展函数，返回 `table`。
- `ts.f8_encode(v double precision[])` 是一个扩展函数，返回 `bytea`。
- `ts.timestamp_decode(v bytea)` 是一个扩展函数，返回 `timestamp[]`。
- `ts.timestamp_encode(v timestamp[])` 是一个扩展函数，返回 `bytea`。
- `ts.u8_decode(v bytea)` 是一个扩展函数，返回 `bigint[]`。
- `ts.u8_encode(v bigint[])` 是一个扩展函数，返回 `bytea`。
- `ts` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为可重定位。
- 控制文件不要求超级用户安装。
- 上游将该项目描述为一个概念验证。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
