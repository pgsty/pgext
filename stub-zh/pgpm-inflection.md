## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/README.md)
- [官方扩展控制文件 (pgpm-inflection.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/pgpm-inflection.control)
- [官方扩展 SQL (pgpm-inflection--0.30.0.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/sql/pgpm-inflection--0.30.0.sql)

`pgpm-inflection` — PostgreSQL 命名规范的字符串变形实用工具。当 SQL 需要这些特殊函数或聚合时使用它。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-inflection";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `inflection.camel(str text)` 是一个扩展函数，返回 `text`。
- `inflection.dashed(str text)` 是一个扩展函数，返回 `text`。
- `inflection.dns_1123(value text)` 是一个扩展函数，返回 `text`。
- `inflection.no_consecutive_caps(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_consecutive_caps_till_end(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_consecutive_caps_till_lower(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_single_underscores(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_single_underscores_at_end(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_single_underscores_in_beginning(str text)` 是一个扩展函数，返回 `text`。
- `inflection.no_single_underscores_in_middle(str text)` 是一个扩展函数，返回 `text`。
- `inflection.pascal(str text)` 是一个扩展函数，返回 `text`。
- `inflection.pg_slugify(text)` 是一个扩展函数，返回 `text`。
- `inflection.pg_slugify(value text, allow_unicode boolean)` 是一个扩展函数，返回 `text`。
- `inflection.plural(str text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.30.0`。
- 请先安装并验证确认的扩展依赖项：`plpgsql`, `unaccent`, `pgpm-verify`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
