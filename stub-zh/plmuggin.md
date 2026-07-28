## 用法

来源：

- [官方上游 README](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/README.md)
- [官方扩展控制文件 (plmuggin.control)](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/plmuggin.control)
- [官方扩展 SQL (plmuggin--0.1.sql)](https://github.com/tatut/plmuggin/blob/77a0c4d35fa1441e3b86dec06ae75ffaa610281d/plmuggin--0.1.sql)

`plmuggin` — PL/Muggin 是一个受 pug 启发的 HTML 模板引擎，定义为 PostgreSQL 语言处理器。Muggin 模板是空白敏感的，看起来像这样：. 使用它时，请确保数据库代码必须在或与该过程语言进行交互运行。上游明确表示该项目尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION plmuggin;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `plmuggin_call_handler()` 是一个扩展函数，返回 `language_handler`。
- `plmuggin_get_metadata(template_name TEXT, meta_key TEXT)` 是一个扩展函数，返回 `TEXT`。
- `plmuggin_templates()` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 控制文件标记该扩展为可信的。
- 上游明确表示该项目尚未准备好用于生产环境。
- 上游描述该项目仍处于开发阶段。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
