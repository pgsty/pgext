

## Usage

> [fuzzystrmatch: Determine similarities and distance between strings](https://www.postgresql.org/docs/current/fuzzystrmatch.html)

The `fuzzystrmatch` module provides functions to determine similarities and distance between strings.

```sql
CREATE EXTENSION fuzzystrmatch;
```

### Soundex

Converts a string to its Soundex code (useful for matching similar-sounding names):

```sql
SELECT soundex('Anne'), soundex('Ann'), difference('Anne', 'Ann');
-- A500, A500, 4

SELECT soundex('Anne'), soundex('Andrew'), difference('Anne', 'Andrew');
-- A500, A536, 2

SELECT soundex('Anne'), soundex('Margaret'), difference('Anne', 'Margaret');
-- A500, M626, 0
```

The `difference` function returns 0–4, where 4 means most similar.

### Daitch-Mokotoff Soundex

Returns a set of Daitch-Mokotoff soundex codes (better for non-English names):

```sql
SELECT daitch_mokotoff('George');
-- {595000}

SELECT daitch_mokotoff('John');
-- {160000,460000}

-- Find names sounding like 'Schwartzenegger'
SELECT * FROM s WHERE daitch_mokotoff(nm) && daitch_mokotoff('Schwartzenegger');
```

Supports indexing with GIN:

```sql
CREATE INDEX ON s USING gin (daitch_mokotoff(nm) gin__int_ops);
```

### Levenshtein Distance

Computes edit distance between two strings (insertions, deletions, substitutions):

```sql
SELECT levenshtein('GUMBO', 'GAMBOL');
-- 2

SELECT levenshtein('GUMBO', 'GAMBOL', 2, 1, 1);
-- 3 (custom costs: insert=2, delete=1, substitute=1)

-- Bounded version (faster, stops early)
SELECT levenshtein_less_equal('extensive', 'exhaustive', 2);
-- 3 (actual distance exceeds threshold, returns actual)

SELECT levenshtein_less_equal('extensive', 'exhaustive', 4);
-- 4
```

### Metaphone

Returns a metaphone code for a string:

```sql
SELECT metaphone('GUMBO', 4);
-- KM
```

### Double Metaphone

Returns primary and alternate codes (handles more name variations):

```sql
SELECT dmetaphone('gumbo');
-- KMP

SELECT dmetaphone_alt('gumbo');
-- KMP
```
