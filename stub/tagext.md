## Usage

Sources:

- [Official upstream README](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/README.md)
- [Official extension control file (tagext.control)](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/CHAPTER_12/tagext/tagext.control)
- [Official extension SQL (tagext--1.0--1.1.sql)](https://github.com/packtpublishing/learn-postgresql-second-edition/blob/9590a7604ce6656fcc6f888830a71e082c794d64/CHAPTER_12/tagext/tagext--1.0--1.1.sql)

`tagext` — Tag Programming Example Extension. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tagext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tag_path(tag_to_search text)` is an extension function and returns `TEXT`.
- `tag_path(tag_to_search text, delimiter text DEFAULT ' > ')` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
