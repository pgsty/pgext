## Usage

Sources:

- [Official douban_fdw README](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/README.md)
- [Extension SQL](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_fdw--1.0.sql)
- [Douban ranking client source](https://github.com/xiaowing/douban_fdw/blob/8576d03007de63e22a2a114d6617756dca4412c5/douban_rank.go)

`douban_fdw` is a read-only foreign data wrapper that exposes two historical Douban movie-ranking endpoints as foreign tables. It was written in Go with cgo for old PostgreSQL releases and depends on the legacy, unauthenticated Douban API, so verify that the remote endpoint still works before using it even for experiments.

### Create a ranking table

The foreign server has no options. Each foreign table requires `rank_name`, whose implemented values are `top250` and `us_box`:

```sql
CREATE EXTENSION douban_fdw;
CREATE SERVER douban_api FOREIGN DATA WRAPPER douban_fdw;

CREATE FOREIGN TABLE douban_top250 (
    id           bigint,
    title        text,
    originname   text,
    rating       real,
    year         text,
    genres       text,
    directors    text,
    casts        text,
    collectcount integer,
    subtype      text,
    url          text
)
SERVER douban_api
OPTIONS (rank_name 'top250');

SELECT id, title, rating
FROM douban_top250;
```

Recognized columns are `casts`, `collectcount`, `directors`, `genres`, `id`, `originname`, `rating`, `subtype`, `title`, `url`, and `year`. The source explicitly rejects the misspelled `imags` column. Columns not represented in the upstream JSON response become null or empty according to the wrapper's conversion logic.

### Operational boundaries

Every scan downloads the complete selected ranking through a hard-coded plain-HTTP `/v2/movie/` endpoint. There is no authentication, timeout, local cache, predicate pushdown, or write support. The upstream README also documents a historical limit of 40 API calls per IP address per hour and requires a UTF-8 database. Treat those limits as historical observations rather than a current service contract.

Because both the API surface and the PostgreSQL compatibility notes are old, pin the source, test network failures and response-shape changes, and do not use this wrapper as a durable production integration. Prefer a maintained, authenticated ingestion job if the data is operationally important.
