## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/README.md)
- [官方扩展控制文件 (pgchess.control)](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/pgchess.control)
- [官方扩展 SQL (pgchess.sql)](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/sql/pgchess.sql)

`pgchess` — pgchess 是一个适用于 PostgreSQL 9.1+ 的国际象棋扩展。当应用程序需要此特定数据库功能时，请使用该扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgchess;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `c_score(IN b game)` 是一个扩展函数，返回 `double`。
- `chess_letter_to_x(char)` 是一个扩展函数，返回 `int2`。
- `chess_x_to_letter(int2)` 是一个扩展函数，返回 `text`。
- `int_to_int_to_location(IN x int , IN y int)` 是一个扩展函数，返回 `location`。
- `is_game_ended(IN b game)` 是一个扩展函数，返回 `boolean`。
- `is_king_safe(IN b game)` 是一个扩展函数，返回 `boolean`。
- `location_to_location_to_move(IN a location , IN b location)` 是一个扩展函数，返回 `move`。
- `piece_display_ascii(character(1))` 是一个扩展函数，返回 `text`。
- `piece_display_utf8(character(1))` 是一个扩展函数，返回 `text`。
- `piece_value(character(1))` 是一个扩展函数，返回 `double`。
- `valid_moves(IN b game)` 是一个扩展函数，返回 `SETOF move`。
- `game` 是一个扩展定义的类型。
- `location` 是一个扩展定义的类型。
- `move` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.7`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
