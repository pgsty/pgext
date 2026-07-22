## Usage

Sources:

- [mipt-asj README at the reviewed commit](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/README.md)
- [Version 0.1 SQL interface](https://github.com/leskin-in/mipt-asj/blob/ad19f53317f1ab189e6615d5a5c5a5df77609c37/mipt-asj--0.1.sql)
- [Approximate string-join paper](http://www.vldb.org/pvldb/vol11/p53-tao.pdf)

`mipt-asj` implements abbreviation-aware approximate string joins based on the Tao-Deng-Stonebraker method. It installs into the fixed `mipt_asj` schema and provides a three-stage workflow: `calc_dict` derives abbreviation rules, `calc_pairs` generates candidate string pairs, and `cmp` applies the final similarity test. Use it only when approximate entity matching is acceptable; it does not provide exact relational equality.

### Core Workflow

Build or curate a two-column rule table, generate candidates from two source relations, and then filter every candidate with the same exactness threshold.

```sql
CREATE EXTENSION "mipt-asj";

CREATE TABLE rules(f varchar, a varchar);
INSERT INTO rules
SELECT * FROM mipt_asj.calc_dict(
  'public.full_forms'::regclass::oid, 'name',
  'public.abbreviations'::regclass::oid, 'name'
);

SELECT p.s1, p.s2
FROM mipt_asj.calc_pairs(
  'public.left_names'::regclass::oid, 'name',
  'public.right_names'::regclass::oid, 'name',
  'public.rules'::regclass::oid, 'f', 'a', 0.7
) AS p
WHERE mipt_asj.cmp(
  p.s1, p.s2,
  'public.rules'::regclass::oid, 'f', 'a', 0.7
);
```

The `exactness` argument ranges from zero to one; the README uses `0.7` as an example. Lower thresholds admit more candidates and generally more false positives. Measure recall and precision against labelled data before choosing a production threshold.

### Function Index

- `calc_dict` reads full-form and abbreviation relations identified by relation OID and column name and returns rule pairs `(f, a)`.
- `calc_pairs` reads two source relations plus the rule relation and returns candidate pairs `(s1, s2)`. Candidates may be false positives.
- `cmp` compares one pair against the same rule relation and threshold and returns the final boolean decision.

### Operational Notes

- The repository is archived, and upstream labels version `0.1` as alpha. Its API, internals, and modern PostgreSQL compatibility are not maintained.
- The interface accepts relation OIDs and column names. Use schema-qualified `regclass` expressions, keep identifiers trusted, and regenerate inputs if referenced tables are dropped and recreated with new OIDs.
- Candidate generation is not the final predicate. Always apply `cmp`, and keep the rule table and `exactness` value consistent between both stages.
- Similarity quality depends on the abbreviation corpus and threshold. Benchmark with domain-specific labelled pairs and review false matches before automating merges or updates.
