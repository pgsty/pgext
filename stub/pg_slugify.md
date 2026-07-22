## Usage

Sources:

- [Official extension SQL](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify--0.0.1.sql)
- [Official control file](https://github.com/pindlebot/pg_slugify/blob/fab77da2df8306eac3f8010ff2c7cad8c3dfb75e/pg_slugify.control)

`pg_slugify` version `0.0.1` provides one pure-SQL function for converting text into a lowercase, ASCII-oriented URL slug. It depends on `unaccent` and always creates the function in the `public` schema.

### Core Workflow

Install the dependency through the extension declaration and generate slugs when writing application rows:

```sql
CREATE EXTENSION pg_slugify CASCADE;

SELECT public.slugify('Crème brûlée: Summer Menu 2026');
```

The result is `creme-brulee-summer-menu-2026` with the standard upstream unaccent rules.

### Transformation Rules

`public.slugify(text)` applies `unaccent`, converts the result to lowercase, replaces each run outside `[a-z0-9_-]` with a hyphen, and removes leading or trailing hyphens. Existing underscores and hyphens are preserved. The function is strict, so NULL produces NULL.

### Operational Caveats

The SQL declares `public.slugify(text)` as `IMMUTABLE`, although its output depends on the installed `unaccent` rules. If those rules change, stored slugs and expression-index entries do not update automatically; regenerate data or reindex as appropriate. The function is ASCII-oriented, may collapse distinct source strings to the same slug, and does not enforce uniqueness or a maximum length. Add an application-specific unique constraint and collision policy where slugs are identifiers. Because the schema is hard-coded, review name conflicts and privileges in `public` before installation.
