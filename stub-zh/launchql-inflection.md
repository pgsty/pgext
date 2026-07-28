## 用法

来源：

- [官方上游 README](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/readme.md)
- [官方扩展控制文件 (launchql-inflection.control)](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/launchql-inflection.control)
- [官方扩展 SQL (launchql-inflection--0.0.2.sql)](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/sql/launchql-inflection--0.0.2.sql)

`launchql-inflection` — inflection 是将 Ruby on Rails 的 Active Support Inflection 类的功能移植到 PostgreSQL 的扩展。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。在安装扩展及其依赖项之前，请确保已正确安装并验证了所有依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-inflection";

select inflection.plural( 'child' );
-- children

select inflection.singular( 'children' );
-- child

select inflection.camel( 'message_properties' );
-- messageProperties

select inflection.pascal( 'web acl' );
-- WebAcl

select inflection.underscore( 'WebACL' );
-- web_acl
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `inflection.camel(str text)` 是一个扩展函数，返回 `text`。
- `inflection.dashed(str text)` 是一个扩展函数，返回 `text`。
- `inflection.lower(str text)` 是一个扩展函数，返回 `text`。
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

- 控制文件声明默认版本为 `0.0.2`。
- 请首先安装确认的扩展依赖项：`plpgsql`, `unaccent`, `uuid-ossp`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
