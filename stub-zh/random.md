

## 用法

> [random: PostgreSQL 的可重现随机数据生成器](https://github.com/tvondra/random)

提供为各种数据类型生成随机值的函数，通过种子控制输出的可重现性。

```sql
CREATE EXTENSION random;
```

### 函数

所有函数接受 `seed`（用于可重现性）和 `nvalues`（不同值的数量）参数。

| 函数 | 说明 |
|---|---|
| `random_string(seed, nvalues, min_length, max_length)` | 随机 ASCII 字符串 |
| `random_bytea(seed, nvalues, min_length, max_length)` | 随机 bytea |
| `random_int(seed, nvalues, min_value, max_value)` | 随机 32 位整数 |
| `random_bigint(seed, nvalues, min_value, max_value)` | 随机 64 位整数 |
| `random_real(seed, nvalues, min_value, max_value)` | 随机 32 位浮点数 |
| `random_double_precision(seed, nvalues, min_value, max_value)` | 随机 64 位浮点数 |
| `random_inet(seed, nvalues)` | 随机 INET 地址（/32 掩码） |
| `random_cnet(seed, nvalues)` | 随机 CIDR，掩码为 8/16/24/32 |
| `random_cnet2(seed, nvalues)` | 随机 CIDR，每个掩码长度均等分配 |
| `random_macaddr(seed, nvalues)` | 随机 6 字节 MAC 地址 |
| `random_macaddr8(seed, nvalues)` | 随机 8 字节 MAC 地址 |

### 示例

```sql
-- 生成可重现的随机整数
SELECT random_int(42, 100, 1, 1000) FROM generate_series(1, 10);

-- 长度为 5-10 的随机字符串
SELECT random_string(42, 1000, 5, 10) FROM generate_series(1, 5);

-- 随机 IP 地址
SELECT random_inet(42, 256) FROM generate_series(1, 5);
```
