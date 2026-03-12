

## Usage

> [explain_ui: generate visual explain plan URLs from SQL](https://github.com/davidgomes/pg-explain-ui)

explain_ui provides a function that takes a SQL query, runs `EXPLAIN` on it, and uploads the plan to Dalibo's explain visualizer, returning a shareable URL.

### Function

```sql
SELECT explain_ui($$
  SELECT b.title, a.name, p.name
  FROM books b
  INNER JOIN authors a ON b.author_id = a.author_id
  INNER JOIN publishers p ON b.publisher_id = p.publisher_id
  ORDER BY b.publication_date DESC
$$);

                    explain_ui
--------------------------------------------------
 https://explain.dalibo.com/plan/ccg2e5fedd913bb7
```

The function:
1. Runs `EXPLAIN (FORMAT JSON)` on the provided SQL query
2. Uploads the plan to [explain.dalibo.com](https://explain.dalibo.com/)
3. Returns a URL to the visual plan representation

The visualizer is built on [pev2](https://github.com/dalibo/pev2) (PostgreSQL Explain Visualizer 2), providing an interactive graphical view of query execution plans.
