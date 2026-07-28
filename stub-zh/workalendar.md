## 用法

来源：

- [官方上游 README](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/README.md)
- [官方扩展控制文件 (workalendar.control)](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/workalendar.control)
- [官方扩展 SQL (workalendar--1.0.0.sql)](https://github.com/elemoine/pg_workalendar/blob/d17dac19ef87d0ca1fc4f2cad6f99eceed544cc6/workalendar--1.0.0.sql)

`workalendar` — pg_workalendar 是一个基于 plpythonu 的 PostgreSQL 扩展，用于工作日历相关的调度、时间间隔或时间序列工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION workalendar;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `holidays(year int, continent text, country text)` 是一个扩展函数，返回 `SETOF`。
- `workon(venv text)` 是一个扩展函数，返回 `void`。
- `holiday` 是一个扩展定义的类型。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 先安装并验证确认的扩展依赖项：`plpython3u`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
