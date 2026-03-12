

## Usage

> [hstore: key-value pair data type](https://www.postgresql.org/docs/current/hstore.html)

The `hstore` extension provides a key/value pair data type for storing semi-structured data in a single column.

```sql
CREATE EXTENSION hstore;
```

### Basic Usage

```sql
SELECT 'name => Alice, age => 30'::hstore;
SELECT 'name => Alice'::hstore -> 'name';           -- Alice
SELECT 'a => 1, b => 2'::hstore ? 'a';              -- true
SELECT 'a => 1'::hstore || 'b => 2'::hstore;        -- "a"=>"1", "b"=>"2"
```

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `->` | Get value by key | `h -> 'key'` |
| `\|\|` | Concatenate | `h1 \|\| h2` |
| `?` | Contains key | `h ? 'key'` |
| `?&` | Contains all keys | `h ?& ARRAY['a','b']` |
| `?\|` | Contains any key | `h ?\| ARRAY['a','b']` |
| `@>` | Contains | `h @> 'a=>1'` |
| `<@` | Contained by | `h <@ 'a=>1, b=>2'` |
| `-` | Delete key(s) | `h - 'key'` or `h - ARRAY['a','b']` |

### Subscript Access

```sql
SELECT h['name'] FROM mytable;
UPDATE mytable SET h['age'] = '31';
```

### Functions

```sql
-- Construction
SELECT hstore('key', 'value');
SELECT hstore(ARRAY['a','b'], ARRAY['1','2']);
SELECT hstore(ROW(1, 'hello'));

-- Extraction
SELECT akeys(h);                    -- text[] of keys
SELECT avals(h);                    -- text[] of values
SELECT skeys(h);                    -- set of keys
SELECT svals(h);                    -- set of values
SELECT each(h);                     -- set of (key, value) records

-- Query
SELECT exist(h, 'key');             -- boolean
SELECT defined(h, 'key');           -- true if non-NULL value

-- Modification
SELECT delete(h, 'key');
SELECT slice(h, ARRAY['a','b']);    -- subset of keys

-- JSON conversion
SELECT hstore_to_json(h);
SELECT hstore_to_jsonb(h);
SELECT hstore_to_json_loose(h);    -- distinguishes numbers/booleans

-- Record conversion
SELECT populate_record(NULL::my_table, h);
```

### Index Support

```sql
CREATE INDEX idx ON t USING gin (h);    -- supports @>, ?, ?&, ?|
CREATE INDEX idx ON t USING gist (h);   -- supports @>, ?, ?&, ?|
CREATE INDEX idx ON t USING btree (h);  -- supports =
CREATE INDEX idx ON t USING hash (h);   -- supports =
```
