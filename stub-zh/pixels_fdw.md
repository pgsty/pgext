## 用法

来源：

- [Official README](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/README.md)
- [Extension control file](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/pixels_fdw.control)
- [Installation SQL](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/pixels_fdw--1.0.sql)
- [Build configuration](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/Makefile)

`pixels_fdw` 通过 PostgreSQL 外部表读取 Pixels 列式文件。这一研究性质的集成使用 Pixels C++ 读取器和数据库服务器本地文件路径。上游明确要求预加载并重启服务器。

### 启用

安装匹配的库依赖后，将该库追加到已有预加载列表并重启。

```conf
shared_preload_libraries = 'pixels_fdw'
```

### 基本用法

外部表的列必须与输入文件匹配。路径和过滤表达式使用上游读取器的语法。

```sql
CREATE EXTENSION pixels_fdw;
CREATE SERVER pixels_server FOREIGN DATA WRAPPER pixels_fdw;
CREATE FOREIGN TABLE example (id int, name varchar, birthday date, score decimal(15,2))
  SERVER pixels_server
  OPTIONS (filename '|/srv/pixels/data|', filters 'id > 1 & score < 90');
SELECT * FROM example;
```

### 依赖与限制

源码构建链接 Pixels C++ 组件、Protocol Buffers 和 C++ 运行库。`filename` 选择文件，`filters` 指定读取器侧过滤。control 声明版本为 `1.0` 且可迁移，但没有 PostgreSQL 大版本矩阵。应限制创建基于文件的外部表的权限，数据库服务账户必须能够读取这些文件。上游未记录写入流程。
