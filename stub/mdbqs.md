## Usage

Sources:

- [Official extension SQL](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs--0.1.sql)
- [Official parser implementation](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/mdbqs_gram.y)
- [Official regression examples](https://github.com/NikitOS94/MDBQT/blob/9b4c5b5b17973660ddf83a02d3b4ca2f8537af94/sql/mdbqs_test.sql)

`mdbqs` 0.1 parses a subset of MongoDB query syntax into the API of the separate `jsquery` PostgreSQL extension and evaluates it against JSONB values. It does not connect to MongoDB. Use it only when an old application needs this exact query dialect and can strictly control the query strings.

### Install and Match JSONB

The control file does not declare the runtime dependency, so install `jsquery` first:

```sql
CREATE EXTENSION jsquery;
CREATE EXTENSION mdbqs;

SELECT '{"a": 2}'::jsonb
       <=> '{ a: { $gte: 2 } }'::mdbquery;

SELECT '{"tags": ["ssl", "security"]}'::jsonb
       <=> '{ tags: { $all: ["ssl", "security"] } }'::mdbquery;
```

The `<=>` operator is defined in both argument orders for `jsonb` and `mdbquery` and returns a boolean.

### Supported Surface

The grammar includes scalar equality and `$eq`, `$ne`, `$lt`, `$lte`, `$gt`, `$gte`; arrays with `$in`, `$nin`, and `$all`; `$size`, `$exists`, `$type`, `$not`, and `$mod`; and expression trees with `$and`, `$or`, and `$nor`. It translates the parsed tree to `jsquery`, so supported value types and matching behavior ultimately depend on the installed `jsquery` version.

### Safety and Compatibility

This is an old parser coupled to PostgreSQL and `jsquery` internal function names. Several MongoDB constructs and types are rejected, and it is not a complete MongoDB compatibility layer. The reviewed parser's error handler calls the C process-exit function on invalid syntax, which can terminate the current backend connection instead of reporting a normal parse error. Never accept arbitrary `mdbquery` text from untrusted users; validate it in an isolated process and bind only known templates. The extension defines no index operator class, so test query plans and expect per-row evaluation unless another compatible index expression is explicitly provided. Pin PostgreSQL and `jsquery`, and test malformed input, nesting depth, numeric conversion, Unicode, NULL/missing fields, and backend failure recovery.
