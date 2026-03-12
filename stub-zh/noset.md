

## 用法

> [noset: 阻止用户通过 SET/RESET 更改会话参数](https://gitlab.com/ongresinc/extensions/noset)

`noset` 是一个可加载模块，阻止特定用户使用 `SET` 或 `RESET` 命令更改会话参数。

```sql
CREATE EXTENSION noset;
```

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'noset'
```

### GUC 参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `noset.enabled` | `false` | 为该角色启用 SET/RESET 阻止 |
| `noset.parameters` | `*` | 要阻止的参数（逗号分隔，`*` = 全部） |

### 设置按用户限制

```sql
-- 阻止用户的所有 SET/RESET
ALTER USER appuser SET noset.enabled = true;

-- 仅阻止特定参数
ALTER USER appuser SET noset.enabled = true;
ALTER USER appuser SET noset.parameters = 'work_mem,jit';
```

### 示例

```sql
-- 作为 appuser：
SET work_mem = '1GB';
-- ERROR: permission denied to set/reset parameter 'set work_mem = '1GB';'

SET maintenance_work_mem = '1GB';
-- SET（允许，不在阻止列表中）
```

### 查找受限用户

```sql
SELECT usename, useconfig FROM pg_user
WHERE useconfig IS NOT NULL
  AND array['noset.enabled=on'] <@ useconfig;
```

### 注意事项

- 不适用于超级用户
- 该扩展撤销了 PUBLIC 对 `set_config` 函数的访问权限
