## Usage

Sources:

- [Official README](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/README.md)
- [Version 0.0.2 installation SQL](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/sql/fuzzywuzzy.sql)
- [Extension control file](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/fuzzywuzzy.control)
- [Official PGXN distribution page](https://pgxn.org/dist/fuzzywuzzy/)

`fuzzywuzzy` provides a single SQL-level similarity score for short text values. It wraps the Levenshtein distance from `fuzzystrmatch` and is useful for simple fuzzy ordering or filtering when a full search system is unnecessary.

### Core Workflow

Install the extension, then call `ratio(text, text)` to obtain an integer score from 0 to 100:

```sql
CREATE EXTENSION fuzzywuzzy;

SELECT name, ratio(name, '人棉锦纶') AS score
FROM products
ORDER BY score DESC
LIMIT 3;
```

The installation script creates `fuzzystrmatch` automatically when it is absent. Provision that dependency explicitly if extension creation is restricted in the target environment.

### Object and Semantics

- `ratio(text, text)` returns 100 for identical non-empty strings and 0 when either argument is NULL or empty.
- The implementation derives its score from `levenshtein(text, text)`, the two input lengths, and an integer return value; it is inspired by fuzzywuzzy but does not expose the wider Python fuzzywuzzy API.
- The function is declared `IMMUTABLE`, and version `0.0.2` is relocatable.

The project is small and has no documented Unicode normalization, collation, length-limit, or indexing contract. Normalize text deliberately, test multibyte and long inputs, and avoid treating the score as a calibrated probability.
