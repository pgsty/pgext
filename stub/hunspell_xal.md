## Usage

Sources:

- [Official README](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/README.md)
- [Official extension control file](https://github.com/ontaev/kalmyk-hunspell/blob/92f10f374186899e658ef06fd823adae089efaee/hunspell_xal.control)

`hunspell_xal` installs a Kalmyk-language Hunspell dictionary for PostgreSQL full-text search. Its word list is based on the Kalmyk-Russian dictionary edited by B. D. Muniev in 1977, and the reviewed upstream version supports noun affixes only.

### Core Workflow

After installing the extension files and creating `hunspell_xal`, inspect the objects installed by the package and attach its Hunspell dictionary to a Kalmyk text-search configuration. A typical PostgreSQL configuration workflow is:

```sql
CREATE EXTENSION hunspell_xal;

SELECT dictname
FROM pg_ts_dict
WHERE dictname LIKE '%xal%';

CREATE TEXT SEARCH CONFIGURATION xal (COPY = simple);
ALTER TEXT SEARCH CONFIGURATION xal
  ALTER MAPPING FOR word, asciiword
  WITH hunspell_xal, simple;
```

Verify tokenization and normalization with representative Kalmyk text before building an index:

```sql
SELECT to_tsvector('xal', 'кальмг келн');
SELECT to_tsquery('xal', 'кальмг');
```

### Linguistic and Operational Boundaries

- The upstream README explicitly says this release supports only noun affixes; verb-affix support was future work. Do not assume complete Kalmyk morphology.
- Dictionary quality depends on the bundled `.aff` and `.dic` data and on the token types mapped in the text-search configuration. Test expected inflections, proper names, spelling variants, and false positives with domain reviewers.
- Updating dictionary files does not rewrite existing stored `tsvector` values. Rebuild generated vectors and related indexes when search semantics change.
- Confirm the dictionary's data provenance and licensing are acceptable for the intended redistribution and deployment context.
