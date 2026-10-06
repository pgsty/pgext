## Usage

Sources:

- [README.md](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/README.md)
- [modules/ts_parser/ts_parser--1.0.sql](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_parser/ts_parser--1.0.sql)
- [modules/ts_parser/ts_parser.control](https://github.com/za-arthur/apod_fts/blob/635ddf0dae7dde6dc2e594ef1b2a7b4f6fd21d10/modules/ts_parser/ts_parser.control)

`ts_parser` installs a text-search parser derived from the PostgreSQL 9.6 default parser. It belongs to the APOD search demonstration but can be installed as a separate extension.

### Core Workflow

```sql
CREATE EXTENSION ts_parser;
SELECT * FROM ts_token_type('ts_parser');
SELECT * FROM ts_parse('ts_parser', 'PostgreSQL full text search');
```

### Operational Boundaries

The parser uses the start, next-token, end, token-type and headline callbacks declared in the versioned SQL. A parser alone is not a complete search configuration: choose dictionaries and token mappings explicitly before using it in an application.

Installation needs a superuser and its compiled C library. No preload is required. The historical origin does not establish compatibility with current PostgreSQL releases; no current support matrix is declared.
