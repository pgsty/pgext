## Usage

Sources:

- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/README.md)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor.control)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor--0.2.0--0.2.1.sql)

`index_advisor` 0.2.1 uses HypoPG to compare planner costs with hypothetical indexes. It returns recommendations, not measured execution-time guarantees, and does not create the recommended physical indexes.

### Core Workflow

Install the required `hypopg` extension first, with the privileges its installation requires. This SQL extension needs no preload or restart.

```sql
CREATE EXTENSION hypopg;
CREATE EXTENSION index_advisor;
CREATE TABLE book (id bigint PRIMARY KEY, title text NOT NULL);
SELECT * FROM index_advisor('SELECT id FROM book WHERE title = $1');
```

`index_advisor(query text)` accepts a parameterized query and returns before/after startup and total cost values, `index_statements` and `errors`. Inspect the suggested DDL and real workload before creating an index. View and materialized-view support depends on what the advisor can resolve to underlying tables.

### 0.2.1 Corrections and Limits

Version 0.2.1 fixes qualification of HypoPG calls when HypoPG uses another schema, non-greedy removal of comments, and cleanup of the advisor's own prepared statement. Update installed databases with `ALTER EXTENSION index_advisor UPDATE TO '0.2.1'` after installing the new SQL files.

Run analysis with deliberate table permissions and representative statistics. Planning cost does not account for every operational cost: additional indexes consume disk and slow writes, and should be validated against production query shapes before deployment.
