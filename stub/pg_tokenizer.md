## Usage

Sources:

- [Installation, upgrades, permissions, and backups](https://vchord.pgsty.com/docs/tokenizer/installation/)
- [Tokenizer and native full-text workflows](https://vchord.pgsty.com/docs/tokenizer/overview/)
- [Text analyzer pipelines](https://vchord.pgsty.com/docs/tokenizer/text-analyzer/)
- [Models and filtered corpus training](https://vchord.pgsty.com/docs/tokenizer/models/)
- [Multilingual tokenization](https://vchord.pgsty.com/docs/tokenizer/multilingual/)
- [SQL and TOML reference](https://vchord.pgsty.com/docs/tokenizer/reference/)
- [1.0.0 control file](https://github.com/pgsty/pg_tokenizer/blob/v1.0.0/pg_tokenizer.control)

`pg_tokenizer` 1.0.0 provides configurable text analysis and model tokenization on PostgreSQL 14–18. It returns string tokens, integer token IDs, or native `tsvector` values for GIN/GiST and compatible BM25 indexes. It has no required extension dependency, needs shared preloading, and installs into the fixed `tokenizer_catalog` schema.

### Analyze and Index Text

Have an administrator add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_text_analyzer('fts_words', $$
character_filters = ["to_lowercase"]
pre_tokenizer = "unicode_segmentation"
$$);
SELECT apply_text_analyzer('Cat dog cat', 'fts_words');
SELECT analyze_to_tsvector('Cat dog cat', 'fts_words');

CREATE TABLE fts_documents (id bigint PRIMARY KEY, passage text, terms tsvector);
INSERT INTO fts_documents VALUES
    (1, 'Cat dog cat', analyze_to_tsvector('Cat dog cat', 'fts_words'));
CREATE INDEX fts_documents_gin ON fts_documents USING gin (terms);
SELECT id FROM fts_documents WHERE terms @@ 'cat <-> dog'::tsquery;
```

`apply_text_analyzer` returns `text[]`; `analyze_to_tsvector` uses those strings as lexemes with positions. Analyzer definitions are catalog data, so these functions are `STABLE`. Store their output in an ordinary column and maintain it on writes; do not use them directly in an expression index or stored generated column.

```sql
CREATE FUNCTION fts_documents_set_terms() RETURNS trigger
LANGUAGE plpgsql AS $$
BEGIN
    NEW.terms := tokenizer_catalog.analyze_to_tsvector(NEW.passage, 'fts_words');
    RETURN NEW;
END;
$$;
CREATE TRIGGER fts_documents_terms
BEFORE INSERT OR UPDATE OF passage ON fts_documents
FOR EACH ROW EXECUTE FUNCTION fts_documents_set_terms();
```

The phrase example uses known analyzer lexemes. The extension does not supply a general `tsquery` builder: query terms must use the same analysis and representation as stored documents. Positions describe the final token sequence, not character offsets in the source text. Filters can change phrase distances. PostgreSQL limits still apply: long lexemes are skipped, positions above 16,383 are clamped, and each lexeme retains at most 256 positions.

### Chinese Text and Model Tokenization

Jieba is an analyzer pre-tokenizer. It can produce native Chinese lexemes without training a model:

```sql
SELECT create_text_analyzer('zh_words', $$
[pre_tokenizer.jieba]
mode = "search"
enable_hmm = true
$$);
SELECT analyze_to_tsvector('PostgreSQL全文检索', 'zh_words');

SELECT create_tokenizer('bert_tokens', $$
model = "bert_base_uncased"
$$);
SELECT tokenize('hello hello', 'bert_tokens');
SELECT tokenize_to_tsvector('hello hello', 'bert_tokens');
```

`tokenize` returns `integer[]`; `tokenize_to_tsvector` stores the model IDs as decimal lexemes. The two native vector functions preserve repeated output tokens without decoding the model or applying another PostgreSQL normalization step. A normal language query such as `plainto_tsquery` does not reproduce model IDs. Use matching vectors with the BM25 query constructor, or construct a native query from the matching lexemes.

### Function and Configuration Index

| Interface | Purpose |
| --- | --- |
| `create_text_analyzer`, `drop_text_analyzer` | Manage TOML pipelines of character filters, a pre-tokenizer, and token filters. |
| `create_tokenizer`, `drop_tokenizer` | Combine an analyzer with a vocabulary or subword model. |
| `create_custom_model_tokenizer_and_trigger` | Train a custom vocabulary, backfill integer token IDs, and create maintenance triggers; an optional seventh `filter` argument selects corpus rows. |
| `create_huggingface_model`, `create_lindera_model` | Import a Hugging Face tokenizer definition or configure a Lindera dictionary. |
| `create_stopwords`, `create_synonym` | Create named dictionaries for analyzer filters. |
| `add_preload_model`, `remove_preload_model`, `list_preload_models` | Manage or inspect model preloading; changes require authorized administrators and a restart. |
| `llmlingua2`, `bert_base_uncased`, `wiki_tocken`, `gemma2b` | Built-in model names. |

The custom-model helper still fills integer arrays, not native vectors. Its filter is a trusted SQL Boolean expression over source columns: only true rows participate, the vocabulary remains append-only, and rows leaving the filter keep their previous target value. Do not pass untrusted SQL. Keep a separate vector-maintenance step when using native full-text indexes.

### Upgrade, Permissions, and Backups

Install the 1.0.0 files and restart PostgreSQL before updating each database:

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '1.0.0';
```

Upgrade paths from 0.1.0 and 0.1.1 preserve catalog rows, custom-model IDs, and the six-argument helper. The upgrade refuses to proceed while expression indexes, generated columns, partition keys, or extended statistics depend on the old immutable analyzer function; replace these uses with maintained columns first. Committed definition changes now invalidate backend caches according to SQL snapshots, but never rewrite stored token arrays or vectors.

Lindera 6 changes dictionary token IDs. Rebuild external dictionaries with Lindera 6, regenerate affected stored tokens/vectors, and rebuild indexes. Native RPM/DEB builds do not embed Lindera dictionaries; provide the external dictionary yourself. Back up those files and the cluster's `pg_tokenizer/preload_models` directory separately from a logical dump.

`pg_database_owner` receives access to database-local catalogs and functions; application roles need explicit grants. Object ownership still controls removal. Cluster preload changes are nontransactional and restricted separately. Vocabulary writers introducing new tokens serialize per model; retry SQLSTATE `40001` under repeatable-read or serializable isolation. Definition-changing APIs disallow two-phase prepare.

An ordinary `pg_dump` now includes analyzer, tokenizer, model, stopword, and synonym definitions, together with custom vocabulary tables, sequences, and source triggers. The shipped `lucene_english`, `nltk_english`, and `iso_english` lists are reserved and recreated by extension installation; use another name for a custom list.
