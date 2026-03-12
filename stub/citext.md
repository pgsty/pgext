

## Usage

> [citext: case-insensitive character string type](https://www.postgresql.org/docs/current/citext.html)

The `citext` extension provides a case-insensitive text type that eliminates the need for `lower()` calls in queries.

```sql
CREATE EXTENSION citext;
```

### Basic Usage

```sql
CREATE TABLE users (
    nick citext PRIMARY KEY,
    pass text NOT NULL
);

INSERT INTO users VALUES ('Larry', 'secret123');

-- Case-insensitive matching
SELECT * FROM users WHERE nick = 'larry';   -- matches 'Larry'
SELECT * FROM users WHERE nick = 'LARRY';   -- matches 'Larry'
```

### Behavior

`citext` performs comparisons by internally converting strings to lowercase. The following operations are case-insensitive with `citext`:

- Comparison operators: `=`, `<>`, `<`, `>`, `<=`, `>=`
- Pattern matching: `LIKE`, `ILIKE`, `~~`, `~~*`
- Regular expressions: `~`, `~*`, `!~`, `!~*`

### Case-Insensitive Functions

When arguments are `citext`, these functions perform case-insensitive matching:

`regexp_match()`, `regexp_matches()`, `regexp_replace()`, `regexp_split_to_array()`, `regexp_split_to_table()`, `replace()`, `split_part()`, `strpos()`, `translate()`

### Advantages Over lower()

- Eliminates verbose `lower()` calls in WHERE clauses
- Supports case-insensitive PRIMARY KEY and UNIQUE constraints
- No need for functional indexes
- Transparent case-folding in all operations

### Limitations

- Case-folding depends on `LC_CTYPE` at database creation
- Slightly less efficient than `text` (copying and conversion overhead)
- Does not support B-tree deduplication
- For better Unicode handling, consider nondeterministic collations instead
