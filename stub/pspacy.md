## Usage

Sources:

- [Extension SQL](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy--1.0.sql)
- [Python implementation](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy.py)
- [Pinned Python dependencies](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/requirements.txt)

`pspacy` 1.0 is the PL/Python extension in the chajda repository that turns spaCy tokenization and lemmatization into PostgreSQL full-text-search values. It can prototype multilingual search normalization, but it executes an old Python dependency stack inside database backends and requires strict privilege and resource controls.

### Core Workflow

Install `plpython3u` and the exact Python/native dependencies into the PostgreSQL server environment, then create the extension and test a language:

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pspacy;

SELECT spacy_lemmatize('en', 'The striped bats were hanging on their feet');
SELECT spacy_tsvector('en', 'The striped bats were hanging on their feet');
SELECT spacy_tsquery('en', 'striped bats');
```

The first call in each backend obtains PostgreSQL's extension directory through `pg_config`, imports the installed Python module, and caches it in PL/Python session state. Language pipelines are loaded lazily; unsupported language codes fall back to the multilingual `xx` pipeline.

### Function Index

- `spacy_load()` initializes the session module cache.
- `spacy_lemmatize(...)` controls lower-casing, special-character removal, stop-word removal, and positional output.
- `spacy_lemmatize_query(...)` joins lemmas with AND operators for query construction.
- `spacy_tsvector(lang, text)` returns a `tsvector`.
- `spacy_tsquery(lang, text)` passes generated query text to the `simple` text-search configuration.

Tokens are truncated to 500 characters. Some processing errors return NULL, so monitor indexing failures explicitly.

### Security and Maintenance Caveats

`plpython3u` is untrusted, and the functions are executable by PUBLIC unless privileges are changed; revoke them and grant only a controlled search role. Model loading can consume substantial backend memory and CPU. All functions are declared immutable and parallel-safe even though initialization runs subprocess/import work and results depend on external Python packages and language data; do not use them in expression indexes or stored generated columns without correcting and validating those volatility/parallel labels. The pinned dependency set and custom spaCy fork date from the 2020-era ecosystem, so isolate it and plan migration before any modern deployment.
