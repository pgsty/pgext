## Usage

Sources:

- [Official upstream documentation](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/README.md)
- [Official extension SQL](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/sql/pg_libnumbertext--0.1.0.sql)
- [Official C++ implementation](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/src/pg_libnumbertext.cpp)

`pg_libnumbertext` 0.1.0 binds PostgreSQL to `libnumbertext`, converting a number supplied as text into words for a requested language. The project describes itself as early alpha. It requires the matching C++ library and language-data files on the database server and only supports UTF-8 databases.

### Convert Numbers to Words

Create the extension and call `numbertext(text, text)` with a digit string and language code:

```sql
CREATE EXTENSION pg_libnumbertext;

SELECT numbertext('1', 'en');
SELECT numbertext('3', 'ja');
```

The documented results are `one` and `三`. Passing the number as text preserves formatting choices that a numeric cast could discard; actual accepted forms and language codes come from the installed `libnumbertext` data.

### API and Data Files

`numbertext(text, text)` is declared `IMMUTABLE` and `STRICT`: NULL inputs return NULL, and PostgreSQL may use results in expression indexes. At runtime the library loads language resources installed below PostgreSQL's extension share directory. Package the binary and those data files as one versioned unit on every server that can execute queries.

The reviewed extension SQL also executes `CREATE TYPE phone_number;` but never supplies input/output functions to complete that shell type. It is unrelated to the documented number-to-words API and should not be used as a real column type.

### Compatibility and Index Safety

The database encoding must be UTF-8. Unsupported languages, malformed numbers, missing data files, or library errors raise PostgreSQL errors. Because the SQL function is immutable while its output depends on external resource files, changing `libnumbertext` or its language data can invalidate stored expectations and expression indexes. Pin the library, resources, and extension together; compare outputs before an upgrade and rebuild dependent indexes when results can change. Test decimals, signs, large values, locale-specific grammar, and application normalization for every enabled language.
