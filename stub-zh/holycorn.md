## 用法

来源：

- [官方 README](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/README.md)
- [官方扩展 SQL](https://github.com/franckverrot/holycorn/blob/6cdfa71caffa28ba3ca720f253c4d5e112d5fa1d/holycorn--1.0.sql)

`holycorn` 是基于沙箱化 mruby VM 的外部数据包装器宿主。它既可以通过内置包装器读取 Redis，也可以执行为外部表提供的服务端 Ruby 生产器。

### 核心流程

通过 `wrapper_class` 选择内置 Redis 包装器；其余表选项会传给 Ruby 包装器构造函数。

```sql
CREATE EXTENSION holycorn;
CREATE SERVER holycorn_server FOREIGN DATA WRAPPER holycorn;

CREATE FOREIGN TABLE redis_table (key text, value text)
  SERVER holycorn_server
  OPTIONS (
    wrapper_class 'HolycornRedis',
    host '127.0.0.1',
    port '6379',
    db '0'
  );

SELECT * FROM redis_table;
```

批量定义时，`IMPORT FOREIGN SCHEMA` 会调用包装器的导入回调。上游 Redis 示例要求远端模式名为 `holycorn_schema`，并建议使用 `prefix` 避免本地名称冲突。

```sql
CREATE SCHEMA holycorn_tables;
IMPORT FOREIGN SCHEMA holycorn_schema
  FROM SERVER holycorn_server
  INTO holycorn_tables
  OPTIONS (
    wrapper_class 'HolycornRedis',
    host '127.0.0.1',
    port '6379',
    db '0',
    prefix 'holycorn_'
  );
```

### 包装器契约

- `wrapper_class` 指定已经编译进 mruby 运行时的 Ruby 对象；内置类为 `HolycornRedis`。
- `wrapper_path` 则指向服务端 Ruby 文件。其对象必须接受带一个环境参数的 `.new`，并提供无参数的 `each`。
- 生产器按外部表列顺序返回行值。Ruby 标量会按上游文档映射为 PostgreSQL 布尔值、整数、浮点数、文本或时间戳。

### 限制与安全

本次复核的上游要求 PostgreSQL `9.4+`，且只记录了读取能力：`SELECT` 可用，但不支持 `INSERT`。自定义 `wrapper_path` 涉及数据库服务端的代码与文件访问，因此只能由可信管理员控制。这是一个较旧且范围有限的 FDW 实现；采用前应在准确的 PostgreSQL 构建上验证类型转换与 Redis 故障行为。
