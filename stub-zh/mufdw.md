## 用法

来源：

- [Official README](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/README.md)
- [Extension control file](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/mufdw.control)
- [Installation SQL](https://github.com/pyhalov/mufdw/blob/5dabf4a2dde02e5d9a3fdcc1b59ea6f1790c4d96/mufdw--0.1.sql)

`mufdw` 是独立的教学外部数据包装器，通过 SPI 读取普通本地表，可用于学习 FDW 的规划与执行流程，并非远端数据库连接器。

### 基本用法

上游示例依次创建回环服务器、用户映射、本地表及匹配的外部表。

```sql
CREATE EXTENSION mufdw;
CREATE SERVER loopback FOREIGN DATA WRAPPER mufdw;
CREATE USER MAPPING FOR PUBLIC SERVER loopback;
CREATE TABLE players(id int PRIMARY KEY, nick text, score int);
CREATE FOREIGN TABLE f_players(id int, nick text, score int)
  SERVER loopback OPTIONS (table_name 'players', schema_name 'public');
SELECT * FROM f_players;
```

### 使用边界

`schema_name`、`table_name` 指定底层本地关系，应保持列定义兼容。源码定义 `mufdw_handler()`、`mufdw_validator(text[], oid)`，control 版本为 `0.1`。这是演示实现，未声明受支持大版本矩阵或生产级保证。不得利用外部表授权绕过底层数据原本的权限设计。
