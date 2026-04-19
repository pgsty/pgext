## Usage

Source: [Docs site](https://michael-golfi.github.io/pg_liquid/), [README](https://github.com/michael-golfi/pg_liquid/blob/main/README.md), [Release v0.1.7](https://github.com/michael-golfi/pg_liquid/releases/tag/v0.1.7), [SQL install script](https://github.com/michael-golfi/pg_liquid/blob/main/sql/pg_liquid--0.1.7.sql)

`pg_liquid` brings Liquid-style graph and compound queries into PostgreSQL. It stores graph state in the `liquid` schema and exposes SQL entry points for plain queries, principal-bound queries, and least-privilege reads.

### Core query surface

```sql
CREATE EXTENSION pg_liquid;

SELECT *
FROM liquid.query($$
  Edge("a","b").
  Edge("b","c").
  Path(X,Y) :- Edge(X,Y).
  Path(X,Y) :- Edge(X,Z), Path(Z,Y).
  Path("a",Y)?
$$) AS t(y text);
```

- `liquid.query(program text)`: executes Liquid facts, rules, and one terminal query.
- `liquid.query_as(principal text, program text)`: principal-bound execution.
- `liquid.read_as(principal text, program text)`: least-privilege read wrapper; intended for application-facing reads.

### Language and modeling features

- facts and rules terminated by `.`
- one terminal `?` query per program
- graph edges via `Edge(...)`
- typed compounds such as `Type@(role=value, ...)`
- query-local recursive rules
- built-in policy compounds such as `CompoundReadByRole` and `liquid/acts_for`

### Row normalizers

```sql
SELECT liquid.create_row_normalizer(
  'account_profiles'::regclass,
  'account_profile_normalizer',
  'AccountProfile',
  '{"account_id":"id","display_name":"display_name","tier":"tier"}'::jsonb,
  true
);
```

- `liquid.create_row_normalizer(...)`: projects relational rows into Liquid compounds.
- `liquid.rebuild_row_normalizer(...)`: regenerates bindings after table changes.
- `liquid.drop_row_normalizer(...)`: removes the normalizer and optionally purges generated bindings.

### Caveats

- Upstream validates the extension on PostgreSQL 14 through 18.
- `query_as` and `read_as` are revoked from `PUBLIC` in the shipped SQL; grant them deliberately.
- `read_as(...)` is the safer application entry point when clients should not assert new facts.
