

## 用法

> [pg_tle: PostgreSQL 可信语言扩展](https://github.com/aws/pg_tle)

`pg_tle` 允许你使用可信语言（SQL、PL/pgSQL、PL/v8、PL/Perl）创建和管理 PostgreSQL 扩展，无需访问文件系统或重启服务器。

在 `postgresql.conf` 中将 `pg_tle` 添加到 `shared_preload_libraries`：

```
shared_preload_libraries = 'pg_tle'
```

### 安装 TLE 扩展

```sql
SELECT pgtle.install_extension(
  'my_extension',      -- 扩展名称
  '1.0',               -- 版本
  'My custom extension', -- 描述
  $_pgtle_$
    CREATE FUNCTION my_func() RETURNS text AS $$
      SELECT 'hello from my_extension';
    $$ LANGUAGE SQL;
  $_pgtle_$
);
```

### 管理扩展版本

```sql
-- 添加升级路径
SELECT pgtle.install_update_path(
  'my_extension',  -- 扩展名称
  '1.0',           -- 源版本
  '1.1',           -- 目标版本
  $_pgtle_$
    CREATE OR REPLACE FUNCTION my_func() RETURNS text AS $$
      SELECT 'hello from my_extension v1.1';
    $$ LANGUAGE SQL;
  $_pgtle_$
);

-- 设置默认版本
SELECT pgtle.set_default_version('my_extension', '1.1');
```

### 使用 TLE 扩展

```sql
CREATE EXTENSION my_extension;
SELECT my_func();  -- 'hello from my_extension'
ALTER EXTENSION my_extension UPDATE TO '1.1';
```

### 删除 TLE 扩展

```sql
DROP EXTENSION my_extension;
SELECT pgtle.uninstall_extension('my_extension');
```

### 钩子与功能

注册自定义钩子（例如密码检查钩子）：

```sql
SELECT pgtle.register_feature('my_password_check', 'passcheck');
SELECT pgtle.unregister_feature('my_password_check', 'passcheck');
```

### 核心函数

| 函数 | 说明 |
|------|------|
| `pgtle.install_extension()` | 安装新的 TLE 扩展 |
| `pgtle.install_update_path()` | 添加版本间的升级路径 |
| `pgtle.set_default_version()` | 设置扩展的默认版本 |
| `pgtle.uninstall_extension()` | 删除 TLE 扩展 |
| `pgtle.register_feature()` | 注册功能钩子 |
| `pgtle.unregister_feature()` | 取消注册功能钩子 |
| `pgtle.available_extensions()` | 列出可用的 TLE 扩展 |
| `pgtle.available_extension_versions()` | 列出可用版本 |
