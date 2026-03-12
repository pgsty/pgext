

## 用法

> [url_encode: PostgreSQL 的 URL 编码与解码函数](https://github.com/okbob/url_encode)

### 函数

#### `url_encode(text) returns text`

对字符串进行百分号编码，用于 URL 中：

```sql
SELECT url_encode('Hello World');
-- Hello%20World

SELECT url_encode('Ahoj Svetе');
-- Ahoj%20Sv%C4%9Bte
```

#### `url_decode(text) returns text`

解码百分号编码的字符串：

```sql
SELECT url_decode('Hello%20World');
-- Hello World

SELECT url_decode('Ahoj%20Sv%C4%9Bte');
-- Ahoj Svetе
```

#### `uri_encode(text) returns text`

编码完整的 URI（保留协议、斜杠等）：

```sql
SELECT uri_encode('http://hu.wikipedia.org/wiki/Sao_Paulo');
-- http://hu.wikipedia.org/wiki/S%C3%A3o_Paulo
```

#### `uri_decode(text) returns text`

解码已编码的 URI：

```sql
SELECT uri_decode('http://hu.wikipedia.org/wiki/S%C3%A3o_Paulo');
-- http://hu.wikipedia.org/wiki/Sao_Paulo
```
