## Usage

Sources:

- [Official v0.1.0 README](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/Readme.md)
- [v0.1.0 release notes](https://github.com/huangjimmy/pg_cjk_parser/releases/tag/v0.1.0)
- [v0.1.0 control file](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser.control)
- [v0.1.0 SQL functions](https://github.com/huangjimmy/pg_cjk_parser/blob/v0.1.0/pg_cjk_parser--0.1.0.sql)

`pg_cjk_parser` is a PostgreSQL full-text-search parser derived from the built-in parser. In a UTF-8 database it keeps the default behavior for non-CJK text while emitting overlapping 2-gram tokens for Chinese, Japanese, and Korean text. The extension installs parser support functions; you create the text-search parser and configuration that use them.

### Core Workflow

```sql
CREATE EXTENSION pg_cjk_parser;

CREATE TEXT SEARCH PARSER public.pg_cjk_parser (
    START = prsd2_cjk_start,
    GETTOKEN = prsd2_cjk_nexttoken,
    END = prsd2_cjk_end,
    LEXTYPES = prsd2_cjk_lextype,
    HEADLINE = prsd2_cjk_headline
);

CREATE TEXT SEARCH CONFIGURATION public.config_2_gram_cjk (
    PARSER = public.pg_cjk_parser
);

SELECT alias, description, token
FROM ts_debug(
    'public.config_2_gram_cjk',
    'PostgreSQL 全文検索和中文检索'
);
```

Use the configuration explicitly in generated `tsvector` columns and queries, or set it as the session default:

```sql
SET default_text_search_config = 'public.config_2_gram_cjk';

SELECT to_tsvector('public.config_2_gram_cjk', '日本語全文検索');
```

### Important Objects

- `prsd2_cjk_start`, `prsd2_cjk_nexttoken`, `prsd2_cjk_end`, `prsd2_cjk_lextype`, and `prsd2_cjk_headline`: support functions used by `CREATE TEXT SEARCH PARSER`.
- `cjk_zht2zhs(text)`: converts mapped Traditional Chinese characters to Simplified Chinese while leaving other characters unchanged.
- Parser token type `cjk`: emits overlapping CJK bigrams; CJK punctuation is emitted as a unigram.

```sql
SELECT cjk_zht2zhs('漢語');
-- 汉语
```

### Version Notes and Caveats

- Version `0.1.0` fixes incorrect `cjk_zht2zhs` scanning across mixed-width UTF-8 characters and corrects handling of four-byte CJK code points.
- Upstream supports PostgreSQL 11 through 18 at this release.
- The database must use UTF-8 for CJK bigram behavior. With another encoding, the parser behaves like the PostgreSQL default parser.
- Creating a text-search parser requires elevated privileges. Decide mappings, dictionaries, stop words, and ranking separately; the example configuration defines only the parser.
