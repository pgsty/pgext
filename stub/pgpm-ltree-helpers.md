## Usage

Sources:

- [packages/ltree-helpers/sql/pgpm-ltree-helpers--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/sql/pgpm-ltree-helpers--0.47.0.sql)
- [packages/ltree-helpers/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/Makefile)
- [packages/ltree-helpers/pgpm-ltree-helpers.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/ltree-helpers/pgpm-ltree-helpers.control)

`pgpm-ltree-helpers` provides three small SQL conversion helpers: `ltree_helpers.to_path`, `ltree_helpers.to_slash` and `ltree_helpers.to_query`.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-ltree-helpers";
SELECT ltree_helpers.to_path('/team/docs');
SELECT ltree_helpers.to_slash('team.docs'::ltree);
SELECT ltree_helpers.to_query('/team/**');
```

### Operational Boundaries

Path conversion strips leading slashes and replaces separators with dots. Query conversion maps a single star to one label and a double star to the lquery multi-label wildcard. These are direct string substitutions followed by casts, not filesystem glob parsing or label escaping. Inputs must satisfy ltree/lquery syntax; invalid labels raise cast errors. Functions are strict and immutable. Dependencies are `plpgsql`, `ltree`, `pgpm-verify`. The installation SQL grants privileges to the pre-existing roles `authenticated`, `anonymous`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
