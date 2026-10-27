## Usage

Sources:

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-tokenizer)

`lakebase_tokenizer` adds whole-word processing and CJK morphological text-search dictionaries to managed Lakebase Postgres. Version 0.2.0 is listed for PostgreSQL 16–18.

### Tokenize and Search

First enable Lakebase Search in project settings. This is irreversible and restarts all project computes, dropping connections. Install the relocatable extension in a schema your database role can manage:

```sql
CREATE SCHEMA IF NOT EXISTS tokenizer_ext;
CREATE EXTENSION IF NOT EXISTS lakebase_tokenizer WITH SCHEMA tokenizer_ext;
SELECT to_tsvector('tokenizer_ext.tokenizer_chinese', '现代分布式文件系统架构');
```

`tokenizer_wholeword` supports normalization, stop words, synonyms and English stemming. `tokenizer_morph` supplies CJK analysis; built-in configurations include `tokenizer_chinese`, `tokenizer_japanese` and `tokenizer_korean`. Generated vectors work with PostgreSQL GIN or `lakebase_text` BM25 indexes; `lakebase_text` is optional. Match document and query configurations, and use `plainto_tsquery` for the mixed-mode Chinese configuration.

### Updates

`ALTER EXTENSION lakebase_tokenizer UPDATE` updates the extension; it does not regenerate stored vectors or rebuild indexes. The documented 0.1.x→0.2.0 update preserves tokenization output and existing configurations, so it needs neither operation. Follow release-specific requirements for later output changes. The provider guides do not establish source language, license or preload settings.
