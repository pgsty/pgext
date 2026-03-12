

## 用法

> [ip4r: 支持 GiST 索引的 IPv4/IPv6 地址与范围类型](https://github.com/RhodiumToad/ip4r)

`ip4r` 扩展提供了专用的 IPv4/IPv6 地址和范围数据类型，具有优秀的包含查询索引支持。

```sql
CREATE EXTENSION ip4r;
```

### 数据类型

| 类型 | 描述 |
|------|------|
| `ip4` | 单个 IPv4 地址（32 位） |
| `ip6` | 单个 IPv6 地址（双 64 位） |
| `ip4r` | IPv4 地址范围 |
| `ip6r` | IPv6 地址范围 |
| `ipaddress` | 混合 IPv4/IPv6 地址 |
| `iprange` | 混合 IPv4/IPv6 范围 |

### 地址输入

```sql
SELECT '192.168.1.1'::ip4;
SELECT '2001:db8::1'::ip6;
SELECT '10.0.0.0/24'::ip4r;                   -- CIDR 表示法
SELECT '192.168.1.100-192.168.1.200'::ip4r;   -- 显式范围
```

### 地址运算符

- **比较**：`=`、`<>`、`<`、`>`、`<=`、`>=`
- **算术**：`+`、`-`（与整数运算）
- **位运算**：`&`（AND）、`|`（OR）、`#`（XOR）、`~`（NOT）

### 地址函数

```sql
SELECT family('192.168.1.1'::ipaddress);       -- 4
SELECT ip4_netmask(24);                         -- 255.255.255.0
```

### 范围运算符

| 运算符 | 描述 |
|--------|------|
| `>>=` | 包含或等于 |
| `>>` | 严格包含 |
| `<<=` | 被包含或等于 |
| `<<` | 严格被包含 |
| `&&` | 重叠 |

### 范围函数

```sql
SELECT lower('10.0.0.0/24'::ip4r);           -- 10.0.0.0
SELECT upper('10.0.0.0/24'::ip4r);           -- 10.0.0.255
SELECT is_cidr('10.0.0.0/24'::ip4r);         -- true
SELECT cidr_split('10.0.0.0-10.0.0.5'::ip4r); -- 分解为 CIDR
SELECT @ '10.0.0.0/24'::ip4r;                 -- 近似大小
```

### 索引

```sql
-- 用于包含查询的 GiST 索引
CREATE INDEX idx ON ipranges USING gist (range);

-- 查找包含特定 IP 的范围
SELECT * FROM ipranges WHERE range >>= '10.0.1.15'::ip4;

-- 查找最精确匹配
SELECT * FROM ipranges
WHERE range >>= '10.0.1.15'::ip4
ORDER BY @ range LIMIT 1;
```
