## Usage

Sources:

- [Pinned version 1.0 installation SQL](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc--1.0.sql)
- [Pinned extension control file](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/nvlfunc.control)
- [Pinned upstream README](https://github.com/rongfengliang/nvl-pg-extension/blob/b2c153a2ff86e5fb2acb4b094448a8835700f031/README.md)

`nvlfunc` version `1.0` provides one Oracle-style `NVL` convenience function. Its actual SQL surface is deliberately narrow: `public.NVL(smallint, smallint)` returns the first non-null argument by calling PostgreSQL `COALESCE`.

### Core Workflow

```sql
CREATE EXTENSION nvlfunc;

SELECT public.NVL(NULL::smallint, 7::smallint);  -- 7
SELECT public.NVL(3::smallint, 7::smallint);     -- 3
SELECT public.NVL(NULL::smallint, NULL::smallint);
```

### Important Objects

- `public.NVL(smallint, smallint) RETURNS smallint`: an `IMMUTABLE` SQL wrapper around `COALESCE($1, $2)`. PostgreSQL folds the unquoted declaration to function name `nvl`.

### Operational Notes

This extension does not implement Oracle's broad `NVL` type-conversion behavior and has no overloads for `integer`, `bigint`, numeric, text, date, or timestamp values. Cast both arguments to `smallint` when literals or `NULL` make overload resolution ambiguous, and use built-in `COALESCE` for other types. The function is installed explicitly in `public` and the extension is non-relocatable, so review name collisions and search-path policy before installation. The upstream repository contains only this function and minimal documentation; validate compatibility on the target PostgreSQL version rather than inferring a wider Oracle-compatibility layer.
