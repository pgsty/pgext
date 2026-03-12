

## Usage

> [acl: access control list data type for PostgreSQL](https://github.com/arkhipov/acl)

The `acl` extension provides Access Control List types for row-level security based on application users, without requiring separate database accounts.

```sql
CREATE EXTENSION acl;
```

### Data Types

- **`ace`**: Standard role-based ACE using PostgreSQL OIDs
- **`ace_int4`**: ACE with 32-bit integer identifiers
- **`ace_int8`**: ACE with 64-bit integer identifiers
- **`ace_uuid`**: ACE with UUID identifiers

ACLs are stored as PostgreSQL arrays of ACE types (e.g., `ace[]`).

### ACE Format

```
[type]/[flags]/[who]=[mask]
```

- **Type**: `a` (allow) or `d` (deny)
- **Flags**: `i` (inherit only), `o` (object inherit), `c` (container inherit), `p` (no propagate), `h` (inherited)
- **Who**: Role name, OID, integer, UUID, or `""` (everyone)
- **Permissions**: `r` (read), `w` (write), `d` (delete), `c` (read ACL), `s` (write ACL), plus 16 custom permissions (0-F)

### Checking Permissions

```sql
-- Check current user's access
SELECT acl_check_access(acl_column, 'rw', false) FROM my_table;

-- Check specific role
SELECT acl_check_access(acl_column, 'r', 'username'::name, false);

-- Check custom int4 roles
SELECT acl_check_access(acl_column, 'rw', ARRAY[1001, 1002]::int4[], false);
```

### ACL Inheritance

```sql
-- Compute child ACL from parent
SELECT acl_merge(parent_acl, child_acl, true, true);
```

### Row-Level Security Example

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
