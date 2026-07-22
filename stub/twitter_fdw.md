## Usage

Sources:

- [Pinned official README](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/README.md)
- [Pinned extension SQL](https://github.com/umitanuki/twitter_fdw/blob/9433a19b1ee7abb1ac08c4471ece2238785d677b/twitter_fdw--1.1.0.sql)

`twitter_fdw` is a historical read-only foreign data wrapper for Twitter's unauthenticated Search API. It maps search results to a predefined `public.twitter` foreign table. The endpoint and response format used by version 1.1.0 have been retired, so this extension is documentation of an old integration rather than a working interface to today's X API.

### Historical Workflow

Creating the extension also creates the wrapper, a server, a current-user mapping, and the foreign table:

```sql
CREATE EXTENSION twitter_fdw;

SELECT from_user, created_at, text
FROM public.twitter
WHERE q = '#postgresql';
```

The virtual `q` column is recognized when used with equality and is percent-encoded into the remote search query. It is an input parameter rather than a property returned for each message.

### Installed Objects

- Functions: `twitter_fdw_handler()` and `twitter_fdw_validator(text[], oid)`.
- Foreign data wrapper: `twitter_fdw`.
- Foreign server: `twitter_service`.
- User mapping: for the role that executes `CREATE EXTENSION`.
- Foreign table: `public.twitter` with `id`, `text`, `from_user`, `from_user_id`, `to_user`, `to_user_id`, `iso_language_code`, `source`, `profile_image_url`, `created_at`, and virtual `q` columns.

### Compatibility and Security Boundaries

The pinned C code calls `http://search.twitter.com/search.json` with libcurl and expects the legacy JSON search fields. That unauthenticated HTTP endpoint no longer represents a supported Twitter/X API. There are no OAuth options, API keys, bearer tokens, TLS endpoint configuration, or mappings for modern response objects and rate-limit rules.

The project is abandoned and targets an early PostgreSQL FDW API. It depends on system libcurl and bundles an old libjson copy. Even if adapted to another endpoint, audit its HTTP transport, parser, memory handling, and PostgreSQL API compatibility before loading native code.

Do not deploy it for current X access and do not treat a successful `CREATE EXTENSION` as proof that remote queries work. Use a maintained API client or build a new FDW against the current authenticated API contract. It does not require preload or restart.
