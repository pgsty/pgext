## 用法

来源：

- [官方上游 README](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/README.md)
- [官方扩展控制文件 (ly2pg.control)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/ly2pg.control)
- [官方扩展 SQL (ly2pg--0.1.sql)](https://github.com/gciolli/pgwtc/blob/4ecaabd7ebe48aae2e627fed470fd52e4b7281fd/ly2pg--0.1.sql)

`ly2pg` — 该目录包含两个针对 PostgreSQL 17 的扩展：使用它来移动、转换或集成相应的数据。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION ly2pg;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `clavis2lilypond(IN clavis , OUT key text , OUT is_maior boolean)` 是一个扩展函数，返回 `record`。
- `clavis2nota(IN clavis , nota OUT nota , is_maior OUT boolean)` 是一个扩展函数，返回 `record`。
- `duration2ticks(text)` 是一个扩展函数，返回 `int`。
- `int_error(text)` 是一个扩展函数，返回 `int`。
- `lilypond(nota)` 是一个扩展函数，返回 `text`。
- `lilypond(nota, text)` 是一个扩展函数，返回 `text`。
- `lilypond(nota, text[])` 是一个扩展函数，返回 `text`。
- `lilypond_book_subjects()` 是一个扩展函数，返回 `SETOF text`。
- `lilypond_voice(src text , vox vox , start_id int DEFAULT NULL , max_count bigint DEFAULT NULL , add_key boolean DEFAULT false , add_time boolean DEFAULT false , add_clef boolean DEFAULT false , add_rest boolean DEFAULT false)` 是一个扩展函数，返回 `text`。
- `nota(jsonb)` 是一个扩展函数，返回 `nota`。
- `nota_add(nota, int)` 是一个扩展函数，返回 `nota`。
- `nota_at_clavis(nota, clavis)` 是一个扩展函数，返回 `int`。
- `tempo(text)` 是一个扩展函数，返回 `tempo`。
- `tempo2text(tempo)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
