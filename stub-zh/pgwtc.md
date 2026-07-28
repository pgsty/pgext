## 用法

来源：

- [官方上游 README](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/README.md)
- [官方扩展控制文件 (pgwtc.control)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/pgwtc.control)
- [官方扩展 SQL (pgwtc--0.1.sql)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/pgwtc--0.1.sql)

`pgwtc` — 该目录包含两个针对 PostgreSQL 17 的扩展。当应用程序需要此特定数据库功能时，请使用此目录。在安装扩展依赖项并验证它们之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION pgwtc;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `lilypond_book_subjects()` 是一个扩展函数，返回 `SETOF`。
- `lilypond_voice(src text , vox vox , start_id int DEFAULT NULL , max_count bigint DEFAULT NULL , add_key boolean DEFAULT false , add_time boolean DEFAULT false , add_clef boolean DEFAULT false , add_rest boolean DEFAULT false)` 是一个扩展函数，返回 `text`。
- `answer` 是一个扩展定义的类型。
- `notes_pretty` 是一个扩展定义的视图。
- `subject_occurrences` 是一个扩展定义的视图。
- `subject_occurrences_pretty` 是一个扩展定义的视图。
- `subject_patterns` 是一个扩展定义的视图。
- `subjects` 是一个扩展定义的视图。
- `subjects_pretty` 是一个扩展定义的视图。
- `metadata` 是一个由扩展安装或管理的表。
- `notes` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 请首先安装并验证确认的扩展依赖项：`ly2pg`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
