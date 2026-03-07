


## 用法

- [介绍博客](https://www.cybertec-postgresql.com/en/bringing-etcd-to-the-database-with-rust-and-pgrx/)
- [Github 仓库](https://github.com/cybertec-postgresql/etcd_fdw)


-----------

## 快速上手

### 1. 启用扩展

```sql
CREATE EXTENSION etcd_fdw;
```

### 2. 创建外部数据包装器

```sql
CREATE FOREIGN DATA WRAPPER etcd_fdw
  HANDLER etcd_fdw_handler
  VALIDATOR etcd_fdw_validator;
```

### 3. 创建服务器

```sql
-- 基本连接
CREATE SERVER etcd_plain
  FOREIGN DATA WRAPPER etcd_fdw
  OPTIONS (connstr '127.0.0.1:2379');

-- 生产环境 etcd，启用认证和 SSL
CREATE SERVER etcd FOREIGN DATA WRAPPER etcd_fdw OPTIONS (
    connstr '127.0.0.1:2379',
    username 'root',
    password 'Etcd.Root',
    ssl_ca '/pg/cert/ca.crt',
    ssl_cert '/pg/cert/server.crt',
    ssl_key '/pg/cert/server.key'
);
```

### 4. 创建外部表

```sql
-- 基本表，映射所有键
CREATE FOREIGN TABLE etcd_kv (key TEXT, value TEXT) SERVER etcd OPTIONS (rowid_column 'key');

-- 带前缀过滤的表（仅包含以 '/config/' 开头的键）
CREATE FOREIGN TABLE etcd_config (key TEXT, value TEXT)
  SERVER etcd OPTIONS (rowid_column 'key', prefix '/config/');
```

### 5. 查询数据

```sql
-- 读取所有键
SELECT * FROM etcd_kv;

-- 按键模式过滤（支持下推）
SELECT * FROM etcd_kv WHERE key LIKE '/app/%';

-- 范围查询
SELECT * FROM etcd_kv WHERE key >= '/a' AND key < '/b';

-- 插入新键
INSERT INTO etcd_kv (key, value) VALUES ('/mykey', 'myvalue');

-- 删除键
DELETE FROM etcd_kv WHERE key = '/mykey';
```

### 6. 与 etcd 实时同步

在 PostgreSQL 外部所做的变更会立即可见：

```bash
# 通过 etcdctl 插入
etcdctl put '/config/db_pool_size' '20'
```

```sql
-- 在 PostgreSQL 中即时可见
SELECT * FROM etcd_config;
     key               | value
-----------------------+-------
 /config/db_pool_size  | 20
(1 row)
```

-----------

## 参考

### 服务器选项

| 选项              | 必填     | 说明                                   |
|-------------------|----------|----------------------------------------|
| `connstr`         | 是       | etcd 端点地址（例如 `127.0.0.1:2379`） |
| `username`        | 否       | 认证用户名                             |
| `password`        | 否       | 认证密码                               |
| `ssl_ca`          | 否       | CA 证书文件路径                        |
| `ssl_cert`        | 否       | 客户端证书文件路径                     |
| `ssl_key`         | 否       | 客户端私钥文件路径                     |
| `ssl_servername`  | 否       | TLS 验证使用的域名                     |
| `connect_timeout` | 否       | 连接超时时间（默认：`10s`）            |
| `request_timeout` | 否       | 请求超时时间（默认：`30s`）            |

### 外部表选项

| 选项           | 默认值   | 说明                                     |
|----------------|----------|------------------------------------------|
| `rowid_column` | 必填     | 用作唯一行标识符的列                     |
| `prefix`       | 无       | 仅包含具有此前缀的键                     |
| `keys_only`    | `false`  | 仅获取键，跳过值                         |
| `revision`     | `0`      | 从指定的 etcd 修订版本读取               |
| `key`          | `\0`     | 范围扫描的起始键                         |
| `range_end`    | 无       | 范围扫描的终止键（不包含）               |
| `consistency`  | `l`      | `l`（线性一致性）或 `s`（串行化）        |

### 查询下推

以下操作会下推至 etcd 执行以提升性能：

- **WHERE**：`=`、`>=`、`>`、`<=`、`<`、`BETWEEN`、`LIKE 'prefix%'`
- **ORDER BY**：远程排序
- **LIMIT/OFFSET**：远程分页

### 限制

- 不支持对键列执行 `UPDATE` 操作。替代方案：先 `INSERT` 新键，再 `DELETE` 旧键。
- 需要 etcd v3 API。
