## 用法

来源：[README](https://github.com/HexaCluster/plpgsql_wrap/blob/v1.0/README.md)、[v1.0 release](https://github.com/HexaCluster/plpgsql_wrap/releases/tag/v1.0)、[control file](https://github.com/HexaCluster/plpgsql_wrap/blob/v1.0/plpgsql_wrap.control)

`plpgsql_wrap` 为 PostgreSQL 提供 Oracle WRAP 风格的 procedural language。使用 `LANGUAGE plpgsql_wrap` 编写的函数会先按 PL/pgSQL 校验，然后以 `PLPGSQLWRAP:1:<hex>` 形式加密存储在 `pg_proc.prosrc` 中。

### 带 Key 安装

使用 32-byte AES-256-GCM key 构建扩展：

```bash
export WRAP_KEY_HEX=$(openssl rand -hex 32)
make WRAP_KEY_HEX=$WRAP_KEY_HEX
sudo make install
```

备份该 key。只有正确的 compiled key 可用时，wrapped functions 才能安全 unwrap 或 restore。

在每个需要该 language 的数据库中安装扩展：

```sql
CREATE EXTENSION plpgsql_wrap; -- requires plpgsql
```

### 创建 Wrapped Functions

使用普通 PL/pgSQL 语法，只是 language name 不同：

```sql
CREATE OR REPLACE FUNCTION public.calculate_bonus(emp_id int, yr int)
RETURNS numeric
LANGUAGE plpgsql_wrap
AS $$
DECLARE
  v_salary numeric;
BEGIN
  SELECT salary INTO v_salary FROM employees WHERE id = emp_id;
  RETURN v_salary * 0.15;
END;
$$;
```

存储的函数体不可读：

```sql
SELECT substring(prosrc, 1, 32) AS wrapped_code
FROM pg_proc
WHERE proname = 'calculate_bonus';
```

### Dump、Restore 与 Unwrap

`pg_dump` 会输出加密后的 `PLPGSQLWRAP:1:` blob。在具有相同 compiled key 的服务器上 restore 可以正常工作。不同 key 会保留 blob，但调用时如果 validator/authentication path 无法认证它，就会失败。

superuser 知道 key 时，可以永久 unwrap 一个函数：

```sql
SELECT plpgsql_wrap.unwrap_procedure(
  'myhexkey',
  'public',
  'calculate_bonus',
  'emp_id int, yr int'
);
```

### 注意事项

- 版本 1.0 支持 PostgreSQL 14-18。
- Control file 要求 `plpgsql`，并且需要 superuser installation。
- 这可以防止随意查看源码和 dumps，但 compiled key 是关键 secret。应相应保护 package artifacts 和 build logs。
- 语法会在加密前校验，因此普通 PL/pgSQL syntax errors 会在写入 encrypted storage 前中止 `CREATE FUNCTION`。
