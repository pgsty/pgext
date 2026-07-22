## Usage

Sources:

- [Official README](https://github.com/byucesoy/pg_color/blob/master/README.md)
- [Type definition SQL](https://github.com/byucesoy/pg_color/blob/master/pg_color.sql)
- [Version 1.1 update SQL](https://github.com/byucesoy/pg_color/blob/master/pg_color--1.0--1.1.sql)

`pg_color` 1.1 is an educational extension that adds a compact RGB color base type. Text input is exactly six uppercase hexadecimal digits without a leading hash; output uses the same form. Version 1.1 adds equality and inequality only.

### Core Workflow

```sql
CREATE EXTENSION pg_color;

CREATE TABLE swatch (
    name text PRIMARY KEY,
    rgb color NOT NULL
);

INSERT INTO swatch VALUES
    ('salmon', 'FA8072'),
    ('navy',   '000080');

SELECT name, rgb
FROM swatch
WHERE rgb = 'FA8072'::color;
```

### Object Index

- `color` stores an RGB value using an integer-sized internal representation and accepts values from 000000 through FFFFFF.
- `color = color` and `color <> color` are the only operators added by version 1.1.
- Input such as a hash-prefixed value or lowercase hexadecimal text is rejected by the parser.

### Caveats

- The extension does not define ordering operators, casts, or a B-tree/hash operator class, so do not assume ordinary indexes can support equality predicates.
- It validates representation, not color spaces, alpha channels, profiles, or accessibility.
- Upstream presents the project as a PostgreSQL extension-development example. Test compatibility with the target server major before adopting it in stored schemas.
