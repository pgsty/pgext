## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed.control)
- [Official README](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/README.md)
- [Official extension SQL](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed--1.0.sql)

`wildspeed` 1.0 provides the GIN operator class `wildcard_ops` for accelerating `LIKE` searches with leading or internal wildcards. It uses a permuterm-style index so patterns such as `%suffix` and `%middle%` can become indexable prefix searches. It does not add support for `ILIKE` or regular-expression operators.

### Core Workflow

Create a GIN index with `wildcard_ops`, then write ordinary `LIKE` predicates on the indexed text expression.

```sql
CREATE EXTENSION wildspeed;

CREATE TABLE words (
    word text NOT NULL
);

CREATE INDEX words_wildcard_idx
    ON words USING gin (word wildcard_ops);

ANALYZE words;

SELECT word
FROM words
WHERE word LIKE '%graph%';
```

Use `EXPLAIN (ANALYZE, BUFFERS)` with representative data and patterns to confirm that the planner chooses the index and that the total cost is favorable.

### Objects and Pattern Behavior

- `wildcard_ops` is a GIN operator class for `text` and PostgreSQL’s `LIKE` operator `~~`.
- `permute(text)` exposes the generated rotations for testing; applications normally do not call it.
- Internal support functions extract permuted keys and turn wildcard terms into GIN partial matches.

Permuterm indexing rotates each indexed word around an end marker. This can make leading and internal wildcards much faster than a sequential scan. Plain prefix patterns may still be better served by a B-tree `text_pattern_ops` index, especially for short prefixes.

### Size, Write, and Compatibility Boundaries

Every rotation becomes an index entry, so a `wildcard_ops` index can be many times larger than its table or a comparable B-tree index. The official example demonstrates substantial build-time and size amplification; treat it as a warning, not a performance guarantee. Budget for GIN build, insert/update, vacuum, backup, and replication cost.

A pattern matching nearly everything, such as `%`, can require a full GIN index scan and offer little selectivity. Test short terms, repeated wildcards, non-ASCII data, database encoding, and collation on the exact deployment. Do not assume `ILIKE`, regex, or locale-specific case folding uses this operator class.

The project was modernized as an extension in 2016 and has seen no recent compatibility work. Build and regression-test it against the target PostgreSQL release before adoption, and compare it with current built-in indexing or application search alternatives.
