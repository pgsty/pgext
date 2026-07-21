## 用法

来源：

- [PL/Proxy 2.12.0 README](https://github.com/plproxy/plproxy/blob/v2.12.0/README.md)
- [PL/Proxy语言语法](https://github.com/plproxy/plproxy/blob/v2.12.0/doc/syntax.md)
- [PL/Proxy集群配置](https://github.com/plproxy/plproxy/blob/v2.12.0/doc/config.md)
- [PL/Proxy 2.12.0 发行版](https://github.com/plproxy/plproxy/releases/tag/v2.12.0)

PL/Proxy 是一个 PostgreSQL 的过程语言处理程序，它允许在不同的 PostgreSQL 数据库之间进行远程过程调用，并可选地实现分片。

### 创建扩展

```sql
CREATE EXTENSION plproxy;
```

### 语言语句

PL/Proxy 函数使用四种类型的语句：

**集群选择** —— 连接到一个预配置的集群：

```sql
CREATE FUNCTION get_user(i_id int) RETURNS SETOF users AS $$
    CLUSTER 'mycluster';
    RUN ON i_id;
$$ LANGUAGE plproxy;
```

**直接连接** —— 使用连接字符串：

```sql
CREATE FUNCTION get_config(key text) RETURNS text AS $$
    CONNECT 'host=remotehost dbname=config';
    SELECT val FROM config WHERE key = $1;
$$ LANGUAGE plproxy;
```

### 执行模式

**RUN ON hash** —— 基于哈希值路由到特定分区：

```sql
CREATE FUNCTION get_user_settings(i_username text) RETURNS SETOF user_settings AS $$
    RUN ON namehash(i_username);
$$ LANGUAGE plproxy;
```

**RUN ON ALL** —— 并行执行在所有数据库上：

```sql
CREATE FUNCTION get_all_counts() RETURNS SETOF record AS $$
    RUN ON ALL;
    SELECT count(*) FROM users;
$$ LANGUAGE plproxy;
```

**RUN ON ANY** —— 随机选择一个服务器：

```sql
CREATE FUNCTION get_random_quote() RETURNS text AS $$
    RUN ON ANY;
    SELECT quote FROM quotes ORDER BY random() LIMIT 1;
$$ LANGUAGE plproxy;
```

### 集群配置

集群通过 SQL/MED（外部数据管理）进行配置：

```sql
CREATE SERVER mycluster FOREIGN DATA WRAPPER plproxy
    OPTIONS (
        connection_lifetime '1800',
        p0 'host=node0 dbname=mydb',
        p1 'host=node1 dbname=mydb',
        p2 'host=node2 dbname=mydb',
        p3 'host=node3 dbname=mydb'
    );

CREATE USER MAPPING FOR CURRENT_USER
    SERVER mycluster
    OPTIONS (user 'proxy_user', password 'secret');
```

### 注意事项

- PL/Proxy 路由函数调用，而不是任意的跨数据库事务。设计远程函数时要确保其具有重试安全性，并明确事务边界。
- 集群定义和用户映射可能会暴露连接细节；保护系统目录访问并优先使用受限的远程角色。
- 2.12.0 版本修复了 `SELECT` 中引用标识符解析问题、`plproxy_fdw_validator` 中的空指针问题、Windows 构建问题以及与 PostgreSQL 19 的兼容性。现有 SQL 对象不需要新的使用模式，但从旧版本扩展升级的数据应运行相应的 `ALTER EXTENSION plproxy UPDATE` 路径。
