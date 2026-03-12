

## 用法

> [pgjwt: PostgreSQL 的 JSON Web Token 支持](https://github.com/michelp/pgjwt)

需要 `pgcrypto` 扩展。

### 签发令牌

```sql
SELECT sign(
    '{"sub":"1234567890","name":"John Doe","admin":true}',
    'secret'
);
```

返回一个 JWT 字符串，例如：`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOi...`

### 验证令牌

```sql
SELECT * FROM verify(
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ',
    'secret'
);
```

返回结果：

| header | payload | valid |
|--------|---------|-------|
| `{"alg":"HS256","typ":"JWT"}` | `{"sub":"1234567890","name":"John Doe","admin":true}` | `t` |

### 算法选择

`sign()` 和 `verify()` 接受可选的第三个参数用于指定算法：`'HS256'`（默认）、`'HS384'` 或 `'HS512'`。

```sql
SELECT sign('{"data":"value"}', 'secret', 'HS384');
```
