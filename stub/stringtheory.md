## Usage

Sources:

- [Official upstream documentation](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/README.md)
- [Official extension SQL](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/stringtheory--1.0.2.sql)
- [Official implementation](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/src/text.cpp)

`stringtheory` provides SIMD-oriented C++ functions for exact text equality and substring search. It uses SSE4.2 on x86_64 and SWAR64 on aarch64/arm64. The functions live in the `stringtheory` schema and are useful for explicitly evaluated comparisons, not as drop-in operators or index access methods.

### Core Workflow

Create the extension and call its schema-qualified functions:

```sql
CREATE EXTENSION stringtheory;

SELECT stringtheory.equals('postgres', 'postgres');
SELECT stringtheory.strstr('postgresql', 'sql');
```

The first query returns true. The second returns `8`, because `stringtheory.strstr(text, text)` reports a zero-based byte offset; it returns `-1` when no match is found.

### Important Objects

- `stringtheory.equals(text, text)` returns a boolean for exact byte equality.
- `stringtheory.strstr(text, text)` returns the first zero-based byte offset or `-1`.

Both functions are declared `STRICT`, `IMMUTABLE`, and parallel safe. A NULL argument therefore returns NULL without entering the C++ function.

### Semantic and Performance Caveats

The reviewed implementation returns false when either input to `stringtheory.equals(text, text)` is empty, including `equals('', '')`. `stringtheory.strstr(text, text)` returns `-1` for an empty haystack or empty needle. These differ from ordinary equality and common substring semantics. Offsets are bytes rather than Unicode character positions, so multibyte text requires special care.

The extension defines neither operators nor an index operator class. A predicate that calls these functions is evaluated as a function expression and does not automatically use an ordinary text B-tree or trigram index. Upstream benchmark numbers came from one AMD host and are not a general performance guarantee. Verify CPU instruction support, exact PostgreSQL ABI, empty-string and Unicode behavior, selectivity, and end-to-end query plans on representative data before replacing built-in operations.
