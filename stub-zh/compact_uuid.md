## 用法

来源：

- [README](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/README.md)
- [Control file / 控制文件](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/compact_uuid.control)
- [Cargo.toml](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/Cargo.toml)
- [src/lib.rs](https://github.com/mikkelam/compact-uuid/blob/f81ba87320fb42fcfa6b8f1c07ee905bb2e2f8d0/src/lib.rs)

`compact_uuid` 以 16 字节存储 UUID，并将其显示为规范的 22 字符 Base64url 文本。审核的 0.1.1 扩展支持 PostgreSQL 18。

### 核心工作流

```sql
CREATE EXTENSION compact_uuid;
SELECT '550e8400-e29b-41d4-a716-446655440000'::compact_uuid;
SELECT 'VQ6EAOKbQdSnFkRmVUQAAA'::compact_uuid::uuid;
CREATE TABLE entities (id compact_uuid PRIMARY KEY, label text);
INSERT INTO entities VALUES
  ('550e8400-e29b-41d4-a716-446655440000', 'example');
SELECT * FROM entities
WHERE id = '550e8400-e29b-41d4-a716-446655440000'::uuid;
```

### 类型兼容性

该类型接受普通 UUID 文本和规范的紧凑文本，但始终输出紧凑表示。它通过混合类型运算符和运算符族与 `uuid` 的比较、哈希行为互操作，支持索引、主键、外键、数组和二进制 COPY。

较短的显示形式不会把 UUID 存储压缩到 16 字节以下。需要统一 SQL 类型的上下文仍须显式转换；上游有意不提供双向隐式转换，因此混合类型的 `JOIN USING` 可能需要改写为显式连接条件。

### 运行

由超级用户安装该可重定位扩展，无需预加载。目前兼容性边界是 PostgreSQL 18；修改应用 UUID 列类型前，应验证客户端编解码和备份恢复行为。
