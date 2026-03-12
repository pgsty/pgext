

## Usage

> [dict_int: integer text search dictionary](https://www.postgresql.org/docs/current/dict-int.html)

Controls the indexing of integers in full-text search to prevent excessive unique-word growth that degrades search performance.

```sql
CREATE EXTENSION dict_int;
```

### Configuration Parameters

| Parameter | Description | Default |
|---|---|---|
| `maxlen` | Maximum number of digits allowed | 6 |
| `rejectlong` | If `true`, reject overlength integers (stop word). If `false`, truncate. | `false` |
| `absval` | If `true`, strip leading `+`/`-` before applying maxlen | `false` |

### Examples

```sql
-- Test the default dictionary
SELECT ts_lexize('intdict', '12345678');
-- {123456}  (truncated to 6 digits by default)

-- Configure to reject long integers
ALTER TEXT SEARCH DICTIONARY intdict (MAXLEN = 4, REJECTLONG = true);
SELECT ts_lexize('intdict', '12345');
-- {}  (rejected as stop word)

SELECT ts_lexize('intdict', '1234');
-- {1234}  (accepted)

-- Apply to a text search configuration
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR int, uint WITH intdict;
```
