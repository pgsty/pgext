## 用法

来源：

- [上游 README](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/README.md)
- [扩展 control 文件](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/wdb_fdw.control)
- [SQL 安装脚本](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/sql/wdb_fdw.sql)
- [C 实现](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/src/wdb_fdw.c)

`wdb_fdw` `0.0.1` 版是用于已挂接 WhiteDB 共享内存数据库的历史外部数据包装器。外部服务器接受 `address` 和 `size` 选项，外部表把自身列映射到 WhiteDB 记录字段。处理器只接入扫描、插入和删除回调；尽管 README 声称支持 UPDATE，源码中的更新回调实际已被注释。

### 历史示例

```sql
CREATE EXTENSION wdb_fdw;
CREATE SERVER white_local FOREIGN DATA WRAPPER wdb_fdw
  OPTIONS (address '1000', size '1000000');
CREATE USER MAPPING FOR CURRENT_USER SERVER white_local;
CREATE FOREIGN TABLE white_items (key text, value text) SERVER white_local;
SELECT * FROM white_items;
```

即使使用匹配的旧工具链，写入路径也不安全。写锁的获取与释放均被注释；删除回调还会忽略 PostgreSQL 选中的目标行，直接取得并删除第一条 WhiteDB 记录。检查历史数据时应把该包装器视为只读。

仓库已经归档，代码面向 2013 年前后的 PostgreSQL FDW API 和 WhiteDB。它只支持少量数值与文本转换，错误和事务语义并不完整，不能假定其崩溃安全性或事务行为与 PostgreSQL 存储等价。不要部署到当前生产服务器。若必须恢复历史数据，应固定匹配的旧工具链，并在隔离环境中使用可丢弃副本操作。
