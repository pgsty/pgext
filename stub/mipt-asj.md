## Usage

Sources:

- [mipt-asj README at the reviewed commit](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/README.md)
- [mipt-asj.control at the reviewed commit](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj.control)
- [Version 0.1 SQL interface](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj--0.1.sql)

`mipt-asj` implements abbreviation-aware approximate string matching based on the Tao-Deng-Stonebraker algorithm. It installs its functions in the `mipt_asj` schema: `calc_dict` derives full-form/abbreviation rules, `calc_pairs` generates candidate pairs, and `cmp` performs the final metric comparison.

### Basic Comparison

```sql
CREATE EXTENSION "mipt-asj";

CREATE TABLE abbreviation_rules (
  full_form varchar,
  abbreviation varchar
);

INSERT INTO abbreviation_rules VALUES ('Massachusetts', 'MA');

SELECT mipt_asj.cmp(
  'Massachusetts',
  'MA',
  'abbreviation_rules'::regclass,
  'full_form',
  'abbreviation',
  0.7
);
```

For a full join pipeline, generate or curate a rules table, call `calc_pairs` to obtain candidates, and then apply `cmp` to remove false positives. Use the same `0.7`-style exactness value in both stages.

### Caveats

- The repository is archived, and upstream explicitly labels version `0.1` as an alpha whose API and internals may change.
- `calc_pairs` can return false positives; it is a candidate generator, not the final equality predicate.
- The API takes relation OIDs and column names and targets old PostgreSQL extension interfaces. Review and test the source before using it on a modern server or untrusted identifiers.
