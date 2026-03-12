

## 用法

> [informix_fdw: 访问 Informix 的外部数据包装器](https://github.com/credativ/informix_fdw)

### 创建服务器

```sql
CREATE EXTENSION informix_fdw;

CREATE SERVER informix_server
  FOREIGN DATA WRAPPER informix_fdw
  OPTIONS (
    informixserver 'informix_hostname',
    informixdir '/opt/informix/csdk'
  );
```

**服务器选项：** `informixserver`（必填，sqlhosts 文件中的 Informix 服务器标识符）、`informixdir`（必填，Informix Client SDK 路径）、`disable_predicate_pushdown`（禁用 WHERE 下推）、`gl_datetime`（自定义日期时间格式，默认 `%iY-%m-%d %H:%M:%S`）、`gl_date`（自定义日期格式，默认 `%iY-%m-%d`）。

### 创建用户映射

```sql
CREATE USER MAPPING FOR CURRENT_USER
  SERVER informix_server
  OPTIONS (username 'informix', password 'secret');
```

### 创建外部表

```sql
CREATE FOREIGN TABLE remote_table (
  id bigint NOT NULL,
  name varchar(255),
  amount numeric(10,2),
  created timestamp
)
SERVER informix_server
OPTIONS (
  database 'mydb',
  table 'remote_table',
  client_locale 'en_US.utf8',
  db_locale 'en_US.819'
);
```

**表选项：** `database`（必填，远程数据库名）、`table` 或 `query`（二选一）、`client_locale`（必填，须与 PostgreSQL 服务器编码匹配）、`db_locale`（必填，须与 Informix 区域设置匹配）、`disable_rowid`（使用可更新游标替代 ROWID）、`enable_blobs`（表包含 BLOB 时需包含）。

使用 `query` 替代 `table` 实现类视图访问：

```sql
CREATE FOREIGN TABLE remote_view (
  id bigint,
  total numeric(10,2)
)
SERVER informix_server
OPTIONS (
  database 'mydb',
  query 'SELECT id, SUM(amount) AS total FROM orders GROUP BY id',
  client_locale 'en_US.utf8',
  db_locale 'en_US.819'
);
```

### CRUD 操作

```sql
SELECT * FROM remote_table WHERE id > 100;
INSERT INTO remote_table (id, name, amount) VALUES (1, 'test', 99.99);
UPDATE remote_table SET amount = 100.00 WHERE id = 1;
DELETE FROM remote_table WHERE id = 1;
```

### 支持的数据类型

查询：BOOLEAN、DATE、DATETIME、INTERVAL、SMALLINT、INTEGER、BIGINT、SERIAL、VARCHAR、CHARACTER、TEXT、NUMERIC、MONEY。

DML 操作：SERIAL、INTEGER、BIGINT、INTERVAL、TEXT、BYTEA、VARCHAR、NVARCHAR、TIMESTAMP、DATE、NUMERIC、MONEY。
