## Usage

Sources:

- [Official repository README](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/README.md)
- [Version 1.0.0 extension SQL](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds--1.0.0.sql)
- [Extension control file](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds.control)

`newsfeeds` publishes the database schema and query helpers for a specific news-gathering application. It stores feed scraping rules and collected headlines, builds a weighted full-text vector, returns pending feeds to an external crawler, and exposes headline search as rows or JSON. It does not fetch web pages by itself.

### Core Workflow

The intended setup is to create the extension in the application's schema, register a feed, and let an external worker consume `pending_feeds()` and populate `headlines`:

```sql
CREATE SCHEMA aggregator;
CREATE EXTENSION newsfeeds WITH SCHEMA aggregator;

INSERT INTO aggregator.newsfeeds
    (url, entries, title_selector, link_selector, feedname)
VALUES
    ('https://example.com/news', '.entry', '.title', 'a', 'Example News');

SELECT aggregator.pending_feeds();
SELECT * FROM aggregator.headlines('7 days', '', 0.1, 100, 0);
SELECT aggregator.headlines_as_json('7 days', '', 0.1, 100, 0);
```

### Object Index

- `newsfeeds` stores source URLs, CSS selectors, refresh intervals, and crawler state.
- `headlines` stores collected metadata, content, labels, archive data, and a weighted `tsvector`; it has a GIN index on `fts`.
- `clean_query(text)` rewrites the application's search syntax.
- `pending_feeds()` returns crawler work as JSON.
- `headlines(...)` and `headlines_as_json(...)` browse or search collected items.
- `fts()` is the trigger function intended to maintain the weighted search vector.

### Operational Notes

The published `1.0.0` SQL is application-specific and incomplete as a standalone extension. It contains a questionable column-level `FOREIGN KEY (newsfeed)` declaration, references `title`, `description`, `rank_modifier`, and `reify_url` objects not created by the same script, and hard-codes the `aggregator` schema in helpers despite the control file declaring the extension relocatable. Validate and patch the exact SQL before attempting a fresh installation.

No upstream crawler implementation or supported deployment contract is included. Treat selector execution, URL fetching, content sanitization, retention, and network security as responsibilities of the external worker. The repository provides only a one-line README and no release or compatibility guidance.
