
## 用法

> [pgextwlist: PostgreSQL 扩展白名单](https://github.com/dimitri/pgextwlist)

`pgextwlist` 实现扩展白名单机制：只有明确允许的扩展才能安装，白名单中的扩展即使由非超级用户请求，也会以超级用户权限安装。

### 配置

添加到 `postgresql.conf`：

```ini
local_preload_libraries = 'pgextwlist'
extwlist.extensions = 'hstore,cube,pg_stat_statements'
```

或按角色设置：

```sql
ALTER ROLE adminuser SET extwlist.extensions = 'pg_stat_statements, postgis';
```

| 参数 | 描述 |
|-----------|-------------|
| `extwlist.extensions` | 逗号分隔的白名单扩展列表 |
| `extwlist.custom_path` | 自定义前/后脚本的文件系统路径 |

### 行为

非超级用户可以安装白名单扩展：

```sql
-- 允许，hstore 在白名单中
CREATE EXTENSION hstore;

-- 阻止，不在白名单中
CREATE EXTENSION earthdistance;
-- ERROR: extension "earthdistance" is not whitelisted
```

对于白名单扩展，`CREATE EXTENSION`、`DROP EXTENSION`、`ALTER EXTENSION ... UPDATE` 和 `COMMENT ON EXTENSION` 会以超级用户身份运行。

### 自定义脚本

将脚本放在 `${extwlist.custom_path}/extname/` 目录：

| 脚本 | 时机 |
|--------|------|
| `before--1.0.sql` | 安装 1.0 版本之前 |
| `before-create.sql` | CREATE 之前，作为回退脚本 |
| `after--1.0.sql` | 安装 1.0 版本之后 |
| `after-create.sql` | CREATE 之后，作为回退脚本 |
| `before-update.sql` / `after-update.sql` | ALTER EXTENSION UPDATE 前后 |
| `before-drop.sql` / `after-drop.sql` | DROP EXTENSION 前后 |

自定义脚本支持模板变量：`@extschema@`、`@current_user@`、`@database_owner@`。
