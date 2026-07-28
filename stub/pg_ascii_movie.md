## Usage

Sources:

- [Official upstream README](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/README.md)
- [Official extension control file (pg_ascii_movie.control)](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/pg_ascii_movie.control)
- [Official extension SQL (pg_ascii_movie--1.0.sql)](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/pg_ascii_movie--1.0.sql)

`pg_ascii_movie` — pg_ascii_movie is a PostgreSQL extension to watch ascii movies. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_ascii_movie;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_frame(in pos integer)` is an extension function and returns `text`.
- `get_wait_time(in pos integer)` is an extension function and returns `double`.
- `play_sw1` is an extension procedure.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `file_fdw`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
