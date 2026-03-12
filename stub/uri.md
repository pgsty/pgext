

## Usage

> [uri: URI data type with validation and component extraction](https://github.com/petere/pguri)

The `uri` extension provides a data type for storing URIs with syntax validation per RFC 3986, component extraction functions, and human-friendly sorting.

```sql
CREATE EXTENSION uri;

CREATE TABLE links (
    id   int PRIMARY KEY,
    link uri
);

INSERT INTO links VALUES (1, 'https://github.com/petere/pguri');
```

### Component Extraction Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `uri_scheme(uri)` | `text` | Scheme (http, ftp, mailto) |
| `uri_userinfo(uri)` | `text` | User info; NULL if absent |
| `uri_host(uri)` | `text` | Hostname or IP address |
| `uri_host_inet(uri)` | `inet` | IP host as inet; NULL if not IP |
| `uri_port(uri)` | `integer` | Port number; NULL if unspecified |
| `uri_path(uri)` | `text` | Path component (never NULL) |
| `uri_path_array(uri)` | `text[]` | Path split by `/` |
| `uri_query(uri)` | `text` | Query string; NULL if absent |
| `uri_fragment(uri)` | `text` | Fragment; NULL if absent |

### Utility Functions

```sql
-- Normalize URI per RFC 3986
SELECT uri_normalize('HTTP://Example.COM/foo/../bar');

-- Percent-encode text
SELECT uri_escape('hello world', true, false);  -- hello+world

-- Decode percent-encoded text
SELECT uri_unescape('hello+world', true, false);  -- hello world
```

### Example

```sql
SELECT uri_scheme(link), uri_host(link), uri_path(link)
FROM links
WHERE uri_host(link) = 'github.com';
```
