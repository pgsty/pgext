

## 用法

> [hashlib: PostgreSQL 的稳定哈希函数库](https://github.com/markokr/pghashlib)

提供实现不会随 PostgreSQL 版本变化的稳定哈希函数。

### 字符串哈希（32 位）

```sql
SELECT hash_string('hello', 'crc32');
SELECT hash_string('hello', 'murmur3');
```

使用可选的初始值：

```sql
SELECT hash_string('hello', 'crc32', 42);
```

### 字符串哈希（64 位）

```sql
SELECT hash64_string('hello', 'city64');
SELECT hash64_string('hello', 'siphash24');
SELECT hash64_string('hello', 'lookup3');
```

### 字符串哈希（128 位）

```sql
SELECT hash128_string('hello', 'md5');
SELECT hash128_string('hello', 'city128');
SELECT hash128_string('hello', 'spooky');
```

### 整数哈希

```sql
SELECT hash_int4(42);        -- 32 位整数的 32 位哈希
SELECT hash_int8(42::bigint); -- 64 位整数的 64 位哈希
```

### 可用算法

| 算法 | CPU 无关 | 位数 | 描述 |
|------|----------|------|------|
| `crc32` | 是 | 32 | CRC32 |
| `murmur3` | 否 | 32 | MurmurHash v3 |
| `md5` | 是 | 128 | MD5 |
| `city64` | 否 | 64 | CityHash64 |
| `city128` | 否 | 128 | CityHash128 |
| `siphash24` | 是 | 64 | SipHash-2-4 |
| `spooky` | 否 | 128 | SpookyHash |
| `lookup2` | 否 | 64 | Jenkins lookup2 |
| `lookup3` | 否 | 64 | Jenkins lookup3 CPU 原生字节序 |
| `lookup3be` | 是 | 64 | Jenkins lookup3 大端字节序 |
| `lookup3le` | 是 | 64 | Jenkins lookup3 小端字节序 |
| `pgsql84` | 否 | 64 | Postgres 8.4+ 中修改的 lookup3 |

整数算法：`wang32`、`wang32mult`、`jenkins`（32 位）；`wang64`、`wang64to32`（64 位）。所有整数算法都是可逆的（1:1 映射），适合在唯一 ID 上创建随机排序。
