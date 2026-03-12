

## 用法

> [acl: PostgreSQL 的访问控制列表数据类型](https://github.com/arkhipov/acl)

`acl` 扩展提供了访问控制列表类型，用于基于应用用户的行级安全控制，无需为每个用户创建独立的数据库账户。

```sql
CREATE EXTENSION acl;
```

### 数据类型

- **`ace`**：基于 PostgreSQL OID 的标准角色 ACE
- **`ace_int4`**：使用 32 位整数标识符的 ACE
- **`ace_int8`**：使用 64 位整数标识符的 ACE
- **`ace_uuid`**：使用 UUID 标识符的 ACE

ACL 以 ACE 类型的 PostgreSQL 数组形式存储（例如 `ace[]`）。

### ACE 格式

```
[type]/[flags]/[who]=[mask]
```

- **Type（类型）**：`a`（允许）或 `d`（拒绝）
- **Flags（标志）**：`i`（仅继承）、`o`（对象继承）、`c`（容器继承）、`p`（不传播）、`h`（已继承）
- **Who（对象）**：角色名、OID、整数、UUID 或 `""`（所有人）
- **Permissions（权限）**：`r`（读取）、`w`（写入）、`d`（删除）、`c`（读取 ACL）、`s`（写入 ACL），以及 16 个自定义权限（0-F）

### 权限检查

```sql
-- 检查当前用户的访问权限
SELECT acl_check_access(acl_column, 'rw', false) FROM my_table;

-- 检查指定角色
SELECT acl_check_access(acl_column, 'r', 'username'::name, false);

-- 检查自定义 int4 角色
SELECT acl_check_access(acl_column, 'rw', ARRAY[1001, 1002]::int4[], false);
```

### ACL 继承

```sql
-- 从父级计算子级 ACL
SELECT acl_merge(parent_acl, child_acl, true, true);
```

### 行级安全示例

```sql
CREATE TABLE file_system (
    id   int PRIMARY KEY,
    name text,
    acl  ace[]
);

ALTER TABLE file_system ENABLE ROW LEVEL SECURITY;

CREATE POLICY read_policy ON file_system FOR SELECT TO PUBLIC
    USING (acl_check_access(acl, 'r', false) = 'r');

CREATE POLICY write_policy ON file_system FOR UPDATE TO PUBLIC
    USING (acl_check_access(acl, 'w', false) = 'w');
```
