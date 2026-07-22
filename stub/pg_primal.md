## Usage

Sources:

- [Official extension source](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/src/lib.rs)
- [Official extension manifest](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/Cargo.toml)
- [Official Primal Server README](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/README.md)

`pg_primal` is the small PostgreSQL helper extension used by Primal Server, a Nostr cache service. It provides Nostr identifier parsing and construction of Primal's media-cache URL; it does not install the server's large application schema or run the Nostr service itself.

### Core Workflow

```sql
CREATE EXTENSION pg_primal;

SELECT octet_length(nostr_parse(repeat('00', 32)));
-- 32

SELECT cdn_url(
  'https://example.org/media/photo.jpg',
  'large',
  false
);
```

### Function Index

- `nostr_parse(text) RETURNS bytea` asks the Nostr SDK to parse the input as a public key first and then as an event ID. A recognized identifier becomes its 32-byte representation; an unrecognized value returns `NULL`.
- `cdn_url(text, text, boolean) RETURNS text` builds a URL under Primal's `https://primal.b-cdn.net/media-cache` endpoint. The first character of the size argument becomes the `s` parameter, the animation flag becomes `a=1` or `a=0`, and the source URL is percent-encoded into `u`.

### Boundaries

The extension manifest advertises PostgreSQL `11` through `16` build features and version `0.0.0`; it has no preload requirement. Both functions are application-specific helpers. `cdn_url` unwraps all three arguments and reads the first size character, so pass non-NULL values and a non-empty size string. The Nostr cache database schema and websocket API are initialized and operated by Primal Server's separate setup workflow.
