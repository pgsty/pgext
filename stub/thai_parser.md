## Usage

Sources:

- [Official upstream README](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/README.md)
- [Official extension SQL](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/sql/thai_parser--1.0.0.sql)
- [Official extension control file](https://github.com/zdk/pg-search-thai/blob/e943164a3640a65c5714028447189f9a3a7b4f82/thai_parser/thai_parser.control)

`thai_parser` supplies Thai word segmentation for PostgreSQL full-text search, where spaces do not reliably mark word boundaries. It creates the parser `thai_parser` and the ready-to-use configuration `thaicfg`; the extension requires `libthai` and an UTF-8 database.

### Core Workflow

Create the extension, derive a stored search vector, and index it with GIN:

```sql
CREATE EXTENSION thai_parser;

CREATE TABLE articles (
    id bigserial PRIMARY KEY,
    body text NOT NULL,
    search tsvector GENERATED ALWAYS AS
        (to_tsvector('thaicfg', body)) STORED
);

CREATE INDEX articles_search_idx ON articles USING GIN (search);

INSERT INTO articles (body) VALUES
    ('ส้มตำกับข้าวเหนียว'),
    ('ต้มยำกุ้ง');

SELECT id, body
FROM articles, to_tsquery('thaicfg', 'ส้มตำ') AS query
WHERE search @@ query;
```

Use a bound parameter and an appropriate query constructor for application input. `to_tsquery` expects full text-search syntax; use `plainto_tsquery` when users should not supply operators.

### Important Objects

- `thai_parser` is both the extension name and the installed text-search parser.
- `thaicfg` is the default text-search configuration created by the extension.
- `ts_parse` exposes tokenizer output for diagnosis. Token IDs 97, 98, and 99 represent Thai words, ASCII words, and spaces respectively.
- The default `thaicfg` mapping indexes Thai-word tokens only. Add and test an `asciiword` mapping in a separate text-search configuration when mixed Thai/English content must be searchable.

### Prerequisites and Caveats

The server build needs the `libthai` development/runtime library used by the parser, and the database encoding must be UTF-8. These are host-level requirements, so a database dump alone is insufficient for restore onto a new server.

The upstream repository also contains optional Hunspell dictionary assets, but they are not required for parser-only use. Thai segmentation and dictionary choices affect indexed lexemes; after changing the extension, dictionary, or mapping, rebuild affected text-search indexes and regression-test representative spelling, compound words, mixed-language text, ranking, and transactional search behavior.
