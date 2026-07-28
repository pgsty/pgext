## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/README.md)
- [官方扩展控制文件 (trunklet.control)](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/trunklet.control)
- [官方扩展 SQL (trunklet--0.2.1--0.3.0.sql)](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/sql/trunklet--0.2.1--0.3.0.sql)

`trunklet` — 确保已经安装了 pg_config 并且在路径中。如果你使用了包管理系统如 RPM 安装 PostgreSQL，请确保也安装了 -devel 包。如果需要，可以告诉构建过程 pg_config 的位置：使用它来进行相应的 SQL 或数据库实用程序工作流。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION trunklet;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小 SQL 代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `_trunklet.attnum__get(table_name regclass , field_name name)` 是一个扩展函数，返回 `pg_attribute`。
- `_trunklet.exec(sql text)` 是一个扩展函数，返回 `void`。
- `_trunklet.function_name(language_id _trunklet.language.language_id%TYPE , function_type text)` 是一个扩展函数，返回 `text`。
- `_trunklet.language__get(language_id _trunklet.language.language_id%TYPE)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.language__get(language_name _trunklet.language.language_name%TYPE)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.language__get_id(language_name _trunklet.language.language_name%TYPE)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.language__get_loose(language_id _trunklet.language.language_id%TYPE)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.name_sanity(field_name text , value text)` 是一个扩展函数，返回 `boolean`。
- `_trunklet.template__get(template_id _trunklet.template.template_id%TYPE , loose boolean DEFAULT false)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.template__get(template_name _trunklet.template.template_name%TYPE , template_version _trunklet.template.template_version%TYPE DEFAULT 1 , loose boolean DEFAULT false)` 是一个扩展函数，返回 `_trunklet`。
- `_trunklet.verify_type(language_name _trunklet.language.language_name%TYPE , allowed_type regtype , supplied_type regtype , which_type text)` 是一个扩展函数，返回 `void`。
- `trunklet.execute_into(template_id _trunklet.template.template_id%TYPE , parameters anyelement)` 是一个扩展函数，返回 `anyelement`。
- `trunklet.execute_into(template_name _trunklet.template.template_name%TYPE , parameters anyelement)` 是一个扩展函数，返回 `anyelement`。
- `trunklet.execute_into(template_name _trunklet.template.template_name%TYPE , template_version _trunklet.template.template_version%TYPE , parameters anyelement)` 是一个扩展函数，返回 `anyelement`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.3.3`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
