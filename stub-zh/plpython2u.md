## 用法

来源：

- [官方文档](https://www.postgresql.org/docs/12/plpython-python23.html)
- [官方扩展控制文件 (plpython2u.control)](https://github.com/postgres/postgres/blob/REL_12_STABLE/src/pl/plpython/plpython2u.control)
- [官方扩展 SQL (plpython2u--1.0.sql)](https://github.com/postgres/postgres/blob/REL_12_STABLE/src/pl/plpython/plpython2u--1.0.sql)

`plpython2u` — 历史上的不可信 PL/Python 2 系列化语言。当数据库代码必须在或与该系列化语言进行交互时，请使用此扩展。经过审核的上游项目已归档或不再维护。

### 核心工作流

```sql
CREATE EXTENSION plpython2u;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 这些 Python 2 语言名称是历史遗留的；当前 PostgreSQL 支持通过 `plpython3u` 提供的 PL/Python。
- PL/Python 是不可信的，因此只有超级用户才能创建该语言的函数。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码的一致性。
