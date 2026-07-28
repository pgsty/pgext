## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pgwaffles/pgwaffles-1.0.1/README.md)
- [Official extension control file (pgwaffles.control)](https://api.pgxn.org/src/pgwaffles/pgwaffles-1.0.1/pgwaffles.control)
- [Official extension SQL (pgwaffles--1.0.1.sql)](https://api.pgxn.org/src/pgwaffles/pgwaffles-1.0.1/pgwaffles--1.0.1.sql)

`pgwaffles` — This extension was created for education purpose only. It will look for several recipes, let you choose the one you'd like to see and display the ingredients and step to make the recipe. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgwaffles;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgwaffles.displayIngredients(out yield integer, out quantity decimal(5,3), out unit text, out name text)` is an extension function and returns `setof`.
- `pgwaffles.displayRecipe()` is an extension function and returns `setof`.
- `pgwaffles.ingredient` is a table installed or managed by the extension.
- `pgwaffles.ingredientInRecipe` is a table installed or managed by the extension.
- `pgwaffles.recipe` is a table installed or managed by the extension.
- `pgwaffles.step` is a table installed or managed by the extension.
- `pgwaffles` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
