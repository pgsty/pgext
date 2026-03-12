

## 用法

> [aux_mysql: MySQL 补充扩展](https://github.com/HaloTech-Co-Ltd/openHalo)

`aux_mysql` 扩展是 openHalo 项目的一部分，为 PostgreSQL 提供 MySQL 兼容函数和特性。它使 PostgreSQL 能够理解 MySQL SQL 方言和通信协议。

### 启用

```sql
CREATE EXTENSION aux_mysql CASCADE;
```

### 概述

与 openHalo 的 MySQL 兼容模式配合使用时，该扩展允许：

- 通过 MySQL 线协议进行 MySQL 客户端连接（3306 端口）
- MySQL 兼容的 SQL 语法解析
- MySQL 兼容的函数和操作符

### MySQL 兼容模式

在 `postgresql.conf` 中配置：

```ini
database_compat_mode = 'mysql'      -- 启用 MySQL 模式
mysql.listener_on = true            -- 启用 MySQL 协议监听
mysql.port = 3306                   -- MySQL 协议端口
```

启用后，MySQL 客户端可以直接连接：

```bash
mysql -P 3306 -h 127.0.0.1
```

### 关键特性

- MySQL 兼容 SQL 方言支持
- MySQL 线协议兼容性（TDS）
- MySQL 风格认证（`mysql_native_password`）
- PostgreSQL 中可用常见 MySQL 函数和操作符

### 注意事项

- 该扩展设计为作为 openHalo 发行版的一部分使用
- 标准 PostgreSQL 连接在 MySQL 协议连接旁边继续工作
- 并非所有 MySQL 特性都受支持；专注于常用功能
