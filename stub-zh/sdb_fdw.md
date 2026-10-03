## 用法

来源：

- [Official README](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/README.md)
- [Extension control file](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/sdb_fdw.control)
- [Installation SQL](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/sdb_fdw--1.0.sql)
- [Build configuration](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/Makefile)

`sdb_fdw` 将 SequoiaDB 集合映射为 PostgreSQL 外部表。官方连接器文档面向历史版本的 SequoiaSQL/PostgreSQL 9.3.4 环境，尚无当前原生 PostgreSQL 的兼容性证据。

### 基本用法

安装匹配的连接器和 SequoiaDB 客户端库后，管理员可创建包装器并映射已有集合。

```sql
CREATE EXTENSION sdb_fdw;
CREATE SERVER sdb_server FOREIGN DATA WRAPPER sdb_fdw
  OPTIONS (address 'localhost', service '11810');
CREATE FOREIGN TABLE records (a integer, b integer, c text)
  SERVER sdb_server OPTIONS (collectionspace 'cs', collection 'cl');
SELECT * FROM records;
```

### 映射与边界

服务器选项指定协调节点地址和服务。表选项 `collectionspace`、`collection` 指定远端集合。带引号的点分列名可映射文档嵌套字段，兼容的 PostgreSQL 数组可映射数组值。适配器及其加密、客户端依赖在 SequoiaDB 源码树中构建，并非独立的现代 PGXS 软件包。部署时须按实际 SequoiaDB 版本核查凭据与写入语义；这里不作跨数据库原子提交保证。
