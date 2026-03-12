

## 用法

> [plproxy: 以过程语言实现的数据库分区](https://github.com/plproxy/plproxy)

PL/Proxy 是一个 PostgreSQL 过程语言处理器，实现 PostgreSQL 数据库之间的远程过程调用，支持可选的分片。

### 创建扩展

```sql
CREATE EXTENSION plproxy;
```

### 语言语句

PL/Proxy 函数使用四种类型的语句：

**集群选择** -- 连接到预配置的集群：

```sql
CREATE FUNCTION get_user(i_id int) RETURNS SETOF users AS $$
    CLUSTER 'mycluster';
    RUN ON i_id;
$$ LANGUAGE plproxy;
```

**直接连接** -- 使用连接字符串：

```sql
CREATE FUNCTION get_config(key text) RETURNS text AS $$
    CONNECT 'host=remotehost dbname=config';
    SELECT val FROM config WHERE key = $1;
$$ LANGUAGE plproxy;
```

### 执行模式

**RUN ON hash** -- 基于哈希路由到特定分区：

```sql
CREATE FUNCTION get_user_settings(i_username text) RETURNS SETOF user_settings AS $$
    RUN ON namehash(i_username);
$$ LANGUAGE plproxy;
```

**RUN ON ALL** -- 在所有数据库上并行执行：

```sql
CREATE FUNCTION get_all_counts() RETURNS SETOF record AS $$
    RUN ON ALL;
    SELECT count(*) FROM users;
$$ LANGUAGE plproxy;
```

**RUN ON ANY** -- 随机选择一个服务器：

```sql
CREATE FUNCTION get_random_quote() RETURNS text AS $$
    RUN ON ANY;
    SELECT quote FROM quotes ORDER BY random() LIMIT 1;
$$ LANGUAGE plproxy;
```

### 集群配置

集群通过 SQL/MED（外部数据管理）配置：

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
