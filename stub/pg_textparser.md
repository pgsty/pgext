## Usage

Sources:

- [Upstream README](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/README.md)
- [Extension SQL](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/pg_textparser--1.0.sql)
- [Extension control file](https://github.com/za-arthur/pg_textparser/blob/45ebf9bcb527f9af5eac0fe739880d7cd5c6a486/pg_textparser.control)
- [Upstream repository](https://github.com/za-arthur/pg_textparser)

`pg_textparser` is an abandoned prototype for a custom text-search parser named `textparser`. Inspect package availability without attempting activation:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'pg_textparser';
```

### Incomplete install script

The published version `1.0` SQL declares parser start, token, end, and lexeme-type functions, but its `CREATE TEXT SEARCH PARSER` statement references `textparser_headline` without first declaring that SQL function. Consequently, `CREATE EXTENSION pg_textparser` from the unmodified upstream files cannot complete successfully. The repository is not marked archived, but this catalog entry is classified as abandoned; a reviewed fork or patch is required before any use.
