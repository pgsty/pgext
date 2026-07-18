## 用法

来源：

- [上游 README](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/README.md)
- [扩展 control 文件](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/wdb_fdw.control)
- [SQL 安装脚本](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/sql/wdb_fdw.sql)
- [C 实现](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/src/wdb_fdw.c)

`wdb_fdw` `0.0.1` 版是用于已挂接 WhiteDB 共享内存数据库的历史外部数据包装器。外部服务器接受 `address` 和 `size` 选项，外部表把自身列映射到 WhiteDB 记录字段。实现声明支持扫描、插入、更新和删除回调。

### 历史示例

```sql
CREATE EXTENSION wdb_fdw;
CREATE SERVER white_local FOREIGN DATA WRAPPER wdb_fdw
  OPTIONS (address '1000', size '1000000');
CREATE USER MAPPING FOR CURRENT_USER SERVER white_local;
CREATE FOREIGN TABLE white_items (key text, value text) SERVER white_local;
SELECT * FROM white_items;
```

仓库已经归档，代码面向 2013 年前后的 PostgreSQL FDW API 和 WhiteDB。它只支持少量数值与文本转换，锁和错误语义并不完整，不能假定其崩溃安全性或事务行为与 PostgreSQL 存储等价。不要部署到当前生产服务器。若必须恢复历史数据，应固定匹配的旧工具链，并在隔离环境中使用可丢弃副本操作。
