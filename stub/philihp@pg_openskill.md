## Usage

Sources:

- [Official README](https://github.com/philihp/pg_openskill/blob/v0.1.0/README.md)
- [Versioned install SQL](https://github.com/philihp/pg_openskill/blob/v0.1.0/pg_openskill--0.1.0.sql)
- [database.dev package](https://database.dev/philihp/pg_openskill)

`philihp@pg_openskill` is a database.dev TLE package that implements OpenSkill Plackett–Luce ratings as pure SQL and PL/pgSQL in the `openskill` schema.

### Enablement

The documented identity is namespaced even though the source files use the `pg_openskill` stem. Install it through dbdev and create the quoted extension name:

```sql
SELECT dbdev.install('philihp@pg_openskill');
CREATE EXTENSION "philihp@pg_openskill" VERSION '0.1.0';
```

The package requires the database.dev installer and `pg_tle`. Its CI validates PostgreSQL 16; other major versions are not claimed by the official source.

### Rating and Ranking

`openskill.rating` returns the default `(mu, sigma)` pair. `openskill.rate` updates teams in finish order, while an optional rank array represents wins and ties. `openskill.ordinal` computes the conservative score `mu - 3 * sigma`.

```sql
SELECT openskill.rating();

SELECT openskill.rate(
  '[[{"mu":25,"sigma":8.333333333333334}],
    [{"mu":25,"sigma":8.333333333333334}]]'::jsonb
);

SELECT openskill.ordinal(openskill.rating());
```

For one player per team, `openskill.rate_1v1` accepts an array of `openskill.rating` values and returns the updated array.

### Numerical Boundary

The implementation preserves the operation order and exponential behavior of openskill.js 5.0.1, and its test suite compares floating-point bit patterns. Ratings still depend on input order and prior results; apply matches deterministically and persist the returned state when incremental history matters.

