

## 用法

> [typeid: PostgreSQL 的 TypeID 支持——带前缀的类型安全、可排序 UUID](https://github.com/blitss/typeid-postgres)

TypeID 是 UUIDv7 的扩展，带有类型前缀，内部存储为前缀加二进制 UUID。

```sql
CREATE EXTENSION typeid;
```

### 函数

| 函数 | 返回类型 | 描述 |
|---|---|---|
| `typeid_generate(prefix TEXT)` | `typeid` | 使用给定前缀生成新的 TypeID |
| `typeid_generate_nil()` | `typeid` | 生成空前缀的 TypeID |
| `typeid_is_valid(input TEXT)` | `BOOLEAN` | 检查 TypeID 字符串是否有效 |
| `typeid_prefix(typeid)` | `TEXT` | 从 TypeID 中提取前缀 |
| `typeid_to_uuid(typeid)` | `UUID` | 将 TypeID 转换为 UUID |
| `uuid_to_typeid(prefix TEXT, uuid UUID)` | `typeid` | 将 UUID 转换为 TypeID |
| `typeid_uuid_generate_v7()` | `UUID` | 生成 UUID v7 |
| `typeid_has_prefix(typeid, prefix TEXT)` | `BOOLEAN` | 检查 TypeID 是否具有指定前缀 |
| `typeid_is_nil_prefix(typeid)` | `BOOLEAN` | 检查 TypeID 是否具有空前缀 |
| `typeid_generate_batch(prefix TEXT, count INTEGER)` | `SETOF typeid` | 批量生成 TypeID |

### 运算符

- `<`, `<=`, `=`, `>=`, `>`, `<>` 用于比较 TypeID
- `@>` 用于检查 TypeID 是否具有特定前缀（例如 `id @> 'user'`）
- B-tree 运算符类用于索引

### 示例

```sql
-- 创建使用 TypeID 主键的表
CREATE TABLE users (
  id typeid DEFAULT typeid_generate('user') NOT NULL PRIMARY KEY,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- 插入数据
INSERT INTO users (id) SELECT typeid_generate('user') FROM generate_series(1, 100);

-- 提取前缀
SELECT typeid_prefix(id) FROM users LIMIT 1;  -- 'user'

-- 使用运算符检查前缀
SELECT * FROM users WHERE id @> 'user';

-- 转换为 UUID
SELECT typeid_to_uuid(id) FROM users LIMIT 1;
```
