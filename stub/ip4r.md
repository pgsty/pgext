

## Usage

> [ip4r: IPv4/IPv6 address and range types with GiST indexing](https://github.com/RhodiumToad/ip4r)

The `ip4r` extension provides specialized data types for IPv4/IPv6 addresses and ranges with superior indexing for containment queries.

```sql
CREATE EXTENSION ip4r;
```

### Data Types

| Type | Description |
|------|-------------|
| `ip4` | Single IPv4 address (32-bit) |
| `ip6` | Single IPv6 address (dual 64-bit) |
| `ip4r` | IPv4 address range |
| `ip6r` | IPv6 address range |
| `ipaddress` | Mixed IPv4/IPv6 address |
| `iprange` | Mixed IPv4/IPv6 range |

### Address Input

```sql
SELECT '192.168.1.1'::ip4;
SELECT '2001:db8::1'::ip6;
SELECT '10.0.0.0/24'::ip4r;                   -- CIDR notation
SELECT '192.168.1.100-192.168.1.200'::ip4r;   -- explicit range
```

### Address Operators

- **Comparison**: `=`, `<>`, `<`, `>`, `<=`, `>=`
- **Arithmetic**: `+`, `-` with integers
- **Bitwise**: `&` (AND), `|` (OR), `#` (XOR), `~` (NOT)

### Address Functions

```sql
SELECT family('192.168.1.1'::ipaddress);       -- 4
SELECT ip4_netmask(24);                         -- 255.255.255.0
```

### Range Operators

| Operator | Description |
|----------|-------------|
| `>>=` | Contains or equal |
| `>>` | Strictly contains |
| `<<=` | Contained in or equal |
| `<<` | Strictly contained in |
| `&&` | Overlaps |

### Range Functions

```sql
SELECT lower('10.0.0.0/24'::ip4r);           -- 10.0.0.0
SELECT upper('10.0.0.0/24'::ip4r);           -- 10.0.0.255
SELECT is_cidr('10.0.0.0/24'::ip4r);         -- true
SELECT cidr_split('10.0.0.0-10.0.0.5'::ip4r); -- decompose to CIDRs
SELECT @ '10.0.0.0/24'::ip4r;                 -- approximate size
```

### Indexing

```sql
-- GiST index for containment queries
CREATE INDEX idx ON ipranges USING gist (range);

-- Find ranges containing a specific IP
SELECT * FROM ipranges WHERE range >>= '10.0.1.15'::ip4;

-- Find most specific match
SELECT * FROM ipranges
WHERE range >>= '10.0.1.15'::ip4
ORDER BY @ range LIMIT 1;
```
