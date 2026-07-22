## Usage

Sources:

- [Official README](https://gitlab.com/depesz/short_ids/-/blob/26d4f30f1b0eef49b1ac81cf9ffae295d4cd33f2/README)
- [Official extension SQL](https://gitlab.com/depesz/short_ids/-/blob/26d4f30f1b0eef49b1ac81cf9ffae295d4cd33f2/short_ids--0.1.0.sql)
- [Official PGXN release](https://pgxn.org/dist/short_ids/0.1.0/)

`short_ids` generates short random text identifiers for use as a column default. It checks the target table for collisions and uses transaction-scoped advisory locks so concurrent callers of the same function do not return the same candidate; a unique or primary-key constraint must still be the authoritative guard.

### Core Workflow

```sql
CREATE EXTENSION short_ids;

CREATE TABLE public.links (
  id text PRIMARY KEY DEFAULT generate_random_id(
    'public', 'links', 'id', 4
  ),
  target_url text NOT NULL
);

INSERT INTO public.links(target_url)
VALUES ('https://example.com/docs')
RETURNING id;
```

For a custom alphabet, pass the optional fifth argument:

```sql
SELECT generate_random_id(
  'public', 'links', 'id', 6,
  'abcdefghijklmnopqrstuvwxyz'
);
```

On collision or advisory-lock contention, `generate_random_id()` increases the candidate length by one and tries again.

### Function Index

- `generate_random_id(table_schema text, table_name text, column_name text, string_length integer, possible_chars text DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz')` returns a currently unused candidate.
- `get_random_string(string_length integer, possible_chars text DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz')` generates a string without checking a table.

### Caveats

Version `0.1.0` uses PostgreSQL `random()`, not a cryptographically secure generator. Identifiers are guessable and must not serve as passwords, bearer tokens, or access-control secrets. Safety depends on all writers respecting the database uniqueness constraint; direct concurrent inserts can race with the pre-check, and callers must retry a uniqueness violation. Validate positive lengths and a nonempty alphabet, schema-qualify the function if installed outside `public`, monitor collision-driven growth, and note that the project is old and cataloged as abandoned.
