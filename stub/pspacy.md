## Usage

Sources:

- [Official extension control file](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy.control)
- [Official upstream README](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/README.md)

`pspacy` — PL/Python full-text-search helpers that use spaCy for multilingual lemmatization, tsvector generation and tsquery generation.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`.
The curated compatibility set is `12,13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pspacy";
SELECT extversion
FROM pg_extension
WHERE extname = 'pspacy';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
