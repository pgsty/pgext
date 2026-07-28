## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/README.md)
- [Official extension control file (pgchess.control)](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/pgchess.control)
- [Official extension SQL (pgchess.sql)](https://api.pgxn.org/src/pgchess/pgchess-0.1.7/sql/pgchess.sql)

`pgchess` — pgchess is a PostgreSQL 9.1+ extension for the game of Chess. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgchess;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `c_score(IN b game)` is an extension function and returns `double`.
- `chess_letter_to_x(char)` is an extension function and returns `int2`.
- `chess_x_to_letter(int2)` is an extension function and returns `text`.
- `int_to_int_to_location(IN x int , IN y int)` is an extension function and returns `location`.
- `is_game_ended(IN b game)` is an extension function and returns `boolean`.
- `is_king_safe(IN b game)` is an extension function and returns `boolean`.
- `location_to_location_to_move(IN a location , IN b location)` is an extension function and returns `move`.
- `piece_display_ascii(character(1))` is an extension function and returns `text`.
- `piece_display_utf8(character(1))` is an extension function and returns `text`.
- `piece_value(character(1))` is an extension function and returns `double`.
- `valid_moves(IN b game)` is an extension function and returns `SETOF move`.
- `game` is an extension-defined type.
- `location` is an extension-defined type.
- `move` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.7`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
