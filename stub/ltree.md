

## Usage

> [ltree: hierarchical tree-like label data type](https://www.postgresql.org/docs/current/ltree.html)

The `ltree` extension provides data types for hierarchical tree-structured data with extensive search facilities.

```sql
CREATE EXTENSION ltree;
```

### Data Types

- **`ltree`**: Label path (e.g., `Top.Science.Astronomy`)
- **`lquery`**: Regex-like pattern for matching ltree values
- **`ltxtquery`**: Full-text-search-like pattern

### Basic Usage

```sql
CREATE TABLE categories (path ltree);
INSERT INTO categories VALUES
    ('Top'), ('Top.Science'), ('Top.Science.Astronomy'),
    ('Top.Hobbies'), ('Top.Collections.Pictures');

-- Find descendants
SELECT path FROM categories WHERE path <@ 'Top.Science';

-- Pattern matching
SELECT path FROM categories WHERE path ~ '*.Astronomy.*';

-- Full-text search
SELECT path FROM categories WHERE path @ 'Science & !Pictures';
```

### Operators

| Operator | Description |
|----------|-------------|
| `@>` | Is ancestor of (or equal) |
| `<@` | Is descendant of (or equal) |
| `~` | Matches lquery pattern |
| `?` | Matches any lquery in array |
| `@` | Matches ltxtquery |
| `\|\|` | Concatenate paths |

### lquery Patterns

```sql
'*.Science.*'           -- any path containing Science
'Top.*{1,2}.Astronomy'  -- 1-2 labels between Top and Astronomy
'*.astro*'              -- prefix matching
'*.Astro*@'             -- case-insensitive prefix
```

### Functions

```sql
SELECT nlevel('Top.Science.Astronomy');                     -- 3
SELECT subltree('Top.Science.Astronomy', 1, 2);            -- Science
SELECT subpath('Top.Science.Astronomy', 1);                 -- Science.Astronomy
SELECT index('a.b.c.d', 'b.c');                             -- 1
SELECT lca('1.2.3', '1.2.3.4.5');                          -- 1.2
SELECT lca(ARRAY['1.2.3'::ltree, '1.2.4'::ltree]);        -- 1.2
```

### Index Support

```sql
-- GiST index (supports @>, <@, ~, ?, @)
CREATE INDEX path_gist_idx ON categories USING gist (path);

-- B-tree index (supports <, <=, =, >=, >)
CREATE INDEX path_idx ON categories USING btree (path);
```
