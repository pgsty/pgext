

## Usage

> [PGroonga Documentation](https://pgroonga.github.io/) | [GitHub: pgroonga/pgroonga](https://github.com/pgroonga/pgroonga)

`pgroonga_database` is a sub-extension of the [PGroonga](https://pgroonga.github.io/) project. It provides database management functionality for PGroonga, which makes PostgreSQL a fast full text search platform for all languages.

PGroonga is a comprehensive full text search solution that uses [Groonga](https://groonga.org/) as a backend. It supports all languages including CJK (Chinese, Japanese, Korean) out of the box, and provides rich features such as:

- Fast full text search with all language support
- Rich query syntax (query language, script syntax)
- JSON search
- HTML/XML tag aware highlighting
- Similar search
- Synonym expansion
- Autocomplete
- Query log analysis

The PGroonga documentation is extensive and spans hundreds of pages. For detailed usage, API reference, operators, functions, and tuning guides, please refer to the official documentation site:

- [PGroonga Official Documentation](https://pgroonga.github.io/)
- [Getting Started](https://pgroonga.github.io/install/)
- [Tutorial](https://pgroonga.github.io/tutorial/)
- [How-to Guides](https://pgroonga.github.io/how-to/)
- [Reference](https://pgroonga.github.io/reference/)

## Quick Start

```sql
CREATE EXTENSION pgroonga_database;
CREATE EXTENSION pgroonga;

-- Create a table with text content
CREATE TABLE memos (
  id integer,
  content text
);

-- Create a PGroonga index
CREATE INDEX pgroonga_content_index ON memos USING pgroonga (content);

-- Insert data
INSERT INTO memos VALUES (1, 'PostgreSQL is a relational database management system.');
INSERT INTO memos VALUES (2, 'Groonga is a fast full text search engine that supports all languages.');
INSERT INTO memos VALUES (3, 'PGroonga is a PostgreSQL extension that uses Groonga as its backend.');

-- Full text search
SELECT * FROM memos WHERE content &@~ 'PostgreSQL OR Groonga';
```
