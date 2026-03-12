


## Usage

> [url_encode: URL encoding and decoding functions for PostgreSQL](https://github.com/okbob/url_encode)

### Functions

#### `url_encode(text) returns text`

Percent-encode a string for use in URLs:

```sql
SELECT url_encode('Hello World');
-- Hello%20World

SELECT url_encode('Ahoj Svetе');
-- Ahoj%20Sv%C4%9Bte
```

#### `url_decode(text) returns text`

Decode a percent-encoded string:

```sql
SELECT url_decode('Hello%20World');
-- Hello World

SELECT url_decode('Ahoj%20Sv%C4%9Bte');
-- Ahoj Svetе
```

#### `uri_encode(text) returns text`

Encode a full URI (preserves scheme, slashes, etc.):

```sql
SELECT uri_encode('http://hu.wikipedia.org/wiki/Sao_Paulo');
-- http://hu.wikipedia.org/wiki/S%C3%A3o_Paulo
```

#### `uri_decode(text) returns text`

Decode an encoded URI:

```sql
SELECT uri_decode('http://hu.wikipedia.org/wiki/S%C3%A3o_Paulo');
-- http://hu.wikipedia.org/wiki/Sao_Paulo
```
