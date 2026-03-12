


## Usage

> [index_advisor: Query index advisor](https://github.com/supabase/index_advisor)

`index_advisor` analyzes a SQL query and recommends indexes that would improve its execution time, reporting before/after cost estimates and the DDL statements to create the suggested indexes.

### Function

```sql
index_advisor(query text)
RETURNS TABLE (
    startup_cost_before jsonb,
    startup_cost_after  jsonb,
    total_cost_before   jsonb,
    total_cost_after    jsonb,
    index_statements    text[],
    errors              text[]
)
```

### Basic Example

```sql
CREATE TABLE book (
    id    int  PRIMARY KEY,
    title text NOT NULL
);

SELECT * FROM index_advisor('
    SELECT book.id FROM book WHERE title = $1
');
```

Returns cost improvements and the recommended `CREATE INDEX` statement.

### Multi-Table Example

```sql
SELECT * FROM index_advisor('
    SELECT book.id, book.title, publisher.name, author.name, review.body
    FROM book
    JOIN publisher ON book.publisher_id = publisher.id
    JOIN author    ON book.author_id = author.id
    JOIN review    ON book.id = review.book_id
    WHERE author.id = $1 AND publisher.id = $2
');
```

The output includes multiple index recommendations across the joined tables with before/after cost comparisons.

### Features

- Supports parameterized queries using `$1`, `$2` syntax
- Handles materialized views
- Resolves underlying tables and columns behind views
- Depends on HypoPG for hypothetical index evaluation
