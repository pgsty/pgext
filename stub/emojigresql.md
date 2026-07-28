## Usage

Sources:

- [Official upstream README](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/README.md)
- [Official extension control file (emojigresql.control)](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/emojigresql.control)
- [Official extension SQL (emojigresql--1.0.sql)](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/emojigresql--1.0.sql)

`emojigresql` — EmojigreSQL is a **pure SQL** PostgreSQL extension designed to seamlessly encode/decode bytea (binary data) and text to/from emojis. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION emojigresql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `emojigresql.decode(text)` is an extension function and returns `bytea`.
- `emojigresql.encode(bytea)` is an extension function and returns `text`.
- `emojigresql.from_text(text)` is an extension function and returns `text`.
- `emojigresql.to_text(text)` is an extension function and returns `text`.
- `emojigresql.chars` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
