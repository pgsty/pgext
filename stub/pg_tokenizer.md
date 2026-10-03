## Usage

Sources:

- [0.1.1 control](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/pg_tokenizer.control)
- [Installation and preload](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/01-installation.md)
- [Tokenizer usage](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/04-usage.md)
- [Reference and n-grams](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/00-reference.md)
- [Models](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/06-model.md)
- [Cache limitations](https://github.com/supervc-stack/pg_tokenizer.rs/blob/0.1.1/docs/07-limitation.md)

`pg_tokenizer` converts text to token IDs for search applications, commonly with VectorChord-BM25. A tokenizer combines text analysis with a vocabulary model. It requires shared preloading and installs its SQL objects in the fixed `tokenizer_catalog` schema.

### Enable and Tokenize

Add the library to the existing preload list and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_tokenizer'
```

```sql
CREATE EXTENSION pg_tokenizer;
SET search_path = public, tokenizer_catalog;

SELECT create_tokenizer('english', $$
model = "llmlingua2"
$$);
SELECT tokenize('PostgreSQL full text search', 'english');
```

`tokenize(text, text)` returns integer token IDs, not a relevance score or one row per token. With the BM25 companion installed, the returned array can be converted to its sparse vector type.

### Analyze Chinese Text

Jieba is a **pre-tokenizer**, not a built-in vocabulary model named jieba. Build a text analyzer and a custom model for the corpus:

```sql
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    passage text,
    token_ids integer[]
);
SELECT create_text_analyzer('chinese', $$
[pre_tokenizer.jieba]
$$);
SELECT create_custom_model_tokenizer_and_trigger(
    tokenizer_name => 'zh_tokenizer',
    model_name => 'zh_model',
    text_analyzer_name => 'chinese',
    table_name => 'documents',
    source_column => 'passage',
    target_column => 'token_ids'
);
INSERT INTO documents(passage) VALUES ('PostgreSQL全文检索');
SELECT tokenize('数据库', 'zh_tokenizer');
```

The helper learns a vocabulary from the source column and creates a trigger to maintain token IDs. Use the same tokenizer/model for documents and queries. Japanese uses an explicitly created Lindera model with a configured dictionary; a bare model name from another tokenizer API is not interchangeable.

### Object and Configuration Index

- `create_tokenizer`, `drop_tokenizer`, `tokenize`: manage and execute tokenizers.
- `create_text_analyzer`, `apply_text_analyzer`: run character filters, pre-tokenization, and token filters.
- `create_custom_model_tokenizer_and_trigger`, `create_lindera_model`, `create_huggingface_model`: create corpus or imported models.
- `create_stopwords`, `create_synonym`: manage dictionaries.
- Built-in models include `llmlingua2`, `bert_base_uncased`, `wiki_tocken`, and `gemma2b`.
- Version 0.1.1 adds the `ngram` token filter; `min_gram` and `max_gram` range from 1 to 255, while `preserve_original` defaults to false. Configurations use TOML.

### Upgrade and Cache Boundaries

```sql
ALTER EXTENSION pg_tokenizer UPDATE TO '0.1.1';
```

After replacing a preloaded library, restart PostgreSQL before updating database objects. The 0.1.0-to-0.1.1 migration adds no SQL objects; behavior changes are in the library. Analyzer, model, and tokenizer objects are cached per connection, and this cache does not follow transaction isolation or rollback. Reconnect or use the relevant drop helper to clear an object retained after rollback. The extension is not relocatable after creation. It supplies tokenization; ranking and search indexes belong to its consumers.
