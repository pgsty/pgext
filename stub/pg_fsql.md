
## Usage

> Sources: [README](https://github.com/yurc/pg_fsql/blob/main/README.md), [control file](https://raw.githubusercontent.com/yurc/pg_fsql/main/pg_fsql.control)

`pg_fsql` is a recursive SQL template engine for PostgreSQL. It combines a C-based placeholder renderer with PL/pgSQL template execution, hierarchical template composition, and optional SPI plan caching. The upstream project emphasizes that it does not require superuser privileges.

### Quick Start

```sql
CREATE EXTENSION pg_fsql;

INSERT INTO fsql.templates (path, cmd, body)
VALUES ('user_count', 'exec',
        'SELECT jsonb_build_object(''total'', count(*))
         FROM users WHERE status = {d[status]!r}');

SELECT fsql.run('user_count', '{"status":"active"}');
SELECT fsql.render('user_count', '{"status":"active"}');
```

### Catalog Tables

The extension installs two main catalog tables:

```sql
fsql.templates (
    path varchar(500) primary key,
    cmd varchar(50),
    body text,
    defaults text,
    cached boolean default false
)

fsql.params (
    key_param varchar(255) primary key,
    type_param varchar(255) not null
)
```

`path` is dot-separated and defines the template hierarchy.

### Commands and Placeholders

The README documents six command types:

- `exec` to execute SQL and return `jsonb`
- `ref` to redirect to another template
- `if` to choose a child branch
- `exec_tpl` to execute SQL and re-render the result as a template
- `map` to collect children into a JSON object
- `NULL` for text fragments inserted into parents

The renderer supports placeholders such as:

- `{d[key]}`
- `{d[key]!r}` for `quote_literal`
- `{d[key]!j}` for JSONB literals
- `{d[key]!i}` for `quote_identifier`

The special key `_self` injects the full input JSON object.

### Public API

The upstream public functions include:

- `fsql.run(path, data, debug)` to execute a template tree
- `fsql.render(path, data)` to preview rendered SQL
- `fsql.tree(path)` to inspect hierarchy
- `fsql.explain(path, data)` to trace expansion
- `fsql.validate()` to check templates
- `fsql.depends_on(path)` to inspect dependencies
- `fsql.clear_cache()` to free cached SPI plans

### Hierarchical Example

```sql
INSERT INTO fsql.templates (path, cmd, body) VALUES
    ('report', 'exec',
     'SELECT jsonb_build_object(''data'', array_agg(row_to_json(t)))
      FROM (SELECT {d[cols]} FROM {d[src]} {d[where]}) t'),
    ('report.cols',  NULL, 'id, name, email'),
    ('report.src',   NULL, 'customers'),
    ('report.where', NULL, 'WHERE city = {d[city]!r}');

SELECT fsql.run('report', '{"city":"Moscow"}');
SELECT fsql.render('report', '{"city":"Moscow"}');
```

### Notes

The README lists PostgreSQL 14+ and `plpgsql`. The control file currently declares SQL extension version `1.0`, even though the package task tracks release `1.1.0`. No official release notes were published in the repository; the user-facing README still documents the same core API and behavior.
