## 用法

来源：

- [官方上游 README](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/README.md)
- [官方扩展控制文件 (myext.control)](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/myext.control)
- [官方扩展 SQL (myext--1.0--noop.sql)](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/myext--1.0--noop.sql)

`myext` — 在安装此示例扩展后，使用 psql、pgAdmin 或一个使用 SQL 内核的 IPython 笔记本运行这些命令。在管理或自动化上述数据库行为时使用它。在目标 PostgreSQL 构建上使用链接的固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION myext;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `myext_reload()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `noop`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
