

## 用法

> [shacrypt: PostgreSQL 的 SHA-crypt 密码哈希](https://github.com/dverite/postgres-shacrypt)

按照 [SHA-crypt 规范](https://www.akkadia.org/drepper/SHA-crypt.txt) 生成 SHA256-CRYPT 和 SHA512-CRYPT 密码哈希。

### 函数

#### `sha256_crypt(key text, salt text) RETURNS text`

```sql
SELECT sha256_crypt('clearpassword', 'somesalt');
-- $5$somesalt$l3SlbI688JBlRS9RWFC1EwZLNJqfQKcrF3yhcbc7ffA
```

使用自定义轮数：

```sql
SELECT sha256_crypt('clearpassword', '$5$rounds=10000$somesalt');
-- $5$rounds=10000$somesalt$OekH6Tu7EOJIAvxKJ4Ko4bG0DxgO83gZODJLTTjXJi5
```

#### `sha512_crypt(key text, salt text) RETURNS text`

```sql
SELECT sha512_crypt('clearpassword', 'somesalt');
-- $6$somesalt$dDcgWMHOtvHI6qT/Khi3uaaxXN6v4N9bnOeWFl/Y6K3pzxi/...
```

### 盐值格式

- 简单盐值：`'somesalt'`
- 带算法前缀：`'$5$somesalt'`（SHA-256）或 `'$6$somesalt'`（SHA-512）
- 带自定义轮数：`'$5$rounds=10000$somesalt'`
