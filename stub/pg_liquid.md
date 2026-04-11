
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_liquid;
> SELECT liquid.query('Edge("a","b"). Edge("b","c"). Path(X,Y) :- Edge(X,Y). Path(X,Y) :- Edge(X,Z), Path(Z,Y). Path("a",Y)?');
> ```
>
> Sources: [README](https://github.com/michael-golfi/pg_liquid), [Docs site](https://michael-golfi.github.io/pg_liquid/)

`pg_liquid` maps the Liquid blog language and data model onto native PostgreSQL storage and execution. The extension exposes SQL entry points for running Liquid-style programs, querying as a principal, and managing row normalizers that project relational rows into Liquid compounds.

## Core Functions

The upstream README lists these main functions:

- `liquid.query(program text)`
- `liquid.query_as(principal text, program text)`
- `liquid.read_as(principal text, program text)`

These support plain execution, principal-aware querying, and CLS-aware reads.

## Language Features

The current README says supported program features include:

- `%` comments
- assertions and rule definitions terminated with `.`
- one terminal `?` query
- `Edge(...)`
- named compounds such as `Type@(cid=..., role=...)`
- query-local recursive rules

## Example Shape

Programs are passed as text and can define facts, rules, and a final query:

```sql
SELECT liquid.query($$
  Edge("a","b").
  Edge("b","c").
  Path(X,Y) :- Edge(X,Y).
  Path(X,Y) :- Edge(X,Z), Path(Z,Y).
  Path("a",Y)?
$$);
```

## Notes

The project README points to the VitePress documentation site as the main documentation surface and notes that operational rollout details are also documented there. The extension is currently published as PGXN package version `0.1.1` and validated against PostgreSQL 14 through 18.
