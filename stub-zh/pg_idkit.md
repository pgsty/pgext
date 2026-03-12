

## 用法

> [pg_idkit: 生成各类新型/小众通用唯一标识符的多功能工具](https://github.com/VADOSWARE/pg_idkit)

```sql
CREATE EXTENSION pg_idkit;
SELECT idkit_uuidv7_generate();
```

### 可用函数

| 方法 | 函数 | 描述 |
|---|---|---|
| UUID v6 | `idkit_uuidv6_generate()` | UUID v6（RFC 4122） |
| | `idkit_uuidv6_generate_uuid()` | UUID v6，返回原生 UUID 类型 |
| | `idkit_uuidv6_extract_timestamptz(TEXT)` | 从 UUID v6 中提取时间戳 |
| UUID v7 | `idkit_uuidv7_generate()` | UUID v7（RFC 4122） |
| | `idkit_uuidv7_generate_uuid()` | UUID v7，返回原生 UUID 类型 |
| | `idkit_uuidv7_extract_timestamptz(TEXT)` | 从 UUID v7 中提取时间戳 |
| NanoID | `idkit_nanoid_generate()` | NanoID |
| | `idkit_nanoid_custom_generate_text()` | 自定义长度和字母表的 NanoID |
| KSUID | `idkit_ksuid_generate()` | K-可排序 UID |
| | `idkit_ksuid_extract_timestamptz(TEXT)` | 从 KSUID 中提取时间戳 |
| | `idkit_ksuidms_generate()` | 毫秒精度的 KSUID |
| | `idkit_ksuidms_extract_timestamptz(TEXT)` | 从 KSUID-ms 中提取时间戳 |
| ULID | `idkit_ulid_generate()` | 通用唯一字典序可排序标识符 |
| | `idkit_ulid_extract_timestamptz(TEXT)` | 从 ULID 中提取时间戳 |
| Timeflake | `idkit_timeflake_generate()` | Snowflake + Instagram ID + Firebase PushID |
| | `idkit_timeflake_extract_timestamptz(TEXT)` | 从 Timeflake 中提取时间戳 |
| PushID | `idkit_pushid_generate()` | Google Firebase PushID |
| XID | `idkit_xid_generate()` | XID |
| | `idkit_xid_extract_timestamptz(TEXT)` | 从 XID 中提取时间戳 |
| CUID | `idkit_cuid_generate()` | CUID（已弃用） |
| | `idkit_cuid_extract_timestamptz(TEXT)` | 从 CUID 中提取时间戳 |
| CUID2 | `idkit_cuid2_generate()` | CUID2 |
| | `idkit_cuid2_generate_with_len(length)` | 自定义长度的 CUID2 |
| TypeID | `idkit_typeid_generate(TEXT)` | 带前缀和 UUIDv7 的 TypeID |
| | `idkit_typeid_generate_text(TEXT)` | TypeID，返回文本 |
| | `idkit_typeid_from_uuid_v7(TEXT, TEXT)` | 从给定的 UUID v7 创建 TypeID |
| | `idkit_typeid_extract_timestamptz(TEXT)` | 从 TypeID 中提取时间戳 |

### 示例

```sql
-- 生成不同类型的 ID
SELECT idkit_uuidv7_generate();           -- 018c106f-9304-79bb-b5be-4483b92b036c
SELECT idkit_nanoid_generate();            -- A8jFA0r3NC6FdalR4LEJ0
SELECT idkit_ksuid_generate();             -- 2HMQIBkTJmEN11JI7tvSTMwfYI3
SELECT idkit_ulid_generate();              -- 01HPYV2X17GM5SQP22M3DVFZY3
SELECT idkit_cuid2_generate();             -- clrjx3bwh0000fj3x4c2y1z0s

-- 提取时间戳
SELECT idkit_uuidv7_extract_timestamptz('018c106f-9304-79bb-b5be-4483b92b036c');
```
