## 用法

来源：

- [官方 README](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/README.md)
- [Rust 实现](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/src/lib.rs)
- [Cargo 兼容特性](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/Cargo.toml)
- [扩展 control 文件](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/plid.control)

`plid` 是受 ULID 启发、带一至三个字母应用前缀的 128 位标识符类型。它存储 16 位前缀、48 位毫秒 Unix 时间戳和 64 位随机数，并打印成小写前缀、下划线与 23 个 Crockford Base32 字符。

### 预加载与创建

单调生成使用 postmaster 共享内存。把共享库加入 `shared_preload_libraries`，重启 PostgreSQL 后再创建扩展。

```conf
shared_preload_libraries = 'plid'
```

```sql
CREATE EXTENSION plid;
```

### 核心流程

为每类实体使用稳定前缀，并在列默认值中生成 ID。前缀只接受 `a-z` 中的一至三个 ASCII 字母。

```sql
CREATE TABLE app_user (
  id plid PRIMARY KEY DEFAULT gen_plid_monotonic('usr'),
  name text NOT NULL
);

INSERT INTO app_user(name) VALUES ('Ada') RETURNING id;
```

可以提取内嵌时间戳，也可以构造随机位全为一的边界 ID，用于时间范围谓词。

```sql
SELECT id, plid_to_timestamptz(id), plid_to_prefix(id)
FROM app_user;

SELECT *
FROM app_user
WHERE id > timestamptz_to_plid('2026-01-01 00:00:00+00', 'usr');
```

### 重要对象

- `plid`：固定 16 字节类型，带比较运算符和 B-tree 运算符类。
- `gen_plid(text)`：生成随机位，不做单调调整。
- `gen_plid_monotonic(text)`：同一毫秒内生成 ID 时递增随机部分。
- `plid_to_timestamptz(plid)`、`plid_to_timestamp(plid)`：提取内嵌时间。
- `timestamptz_to_plid(timestamptz, text)`：创建随机位全为一的边界值。
- `plid_to_prefix(plid)`：提取前缀。

### 排序与作用域

前缀位是最高有效位。全局排序会先按前缀分组，再比较时间戳，因此只有谓词使用相同前缀时，时间范围查询才有意义。单调状态属于单个 PostgreSQL postmaster；重启后会重置，也不会在主备、分片或独立服务器之间协调。它不提供全局序列保证。

文本输入不区分大小写，输出会把前缀规范为小写、Base32 部分规范为大写。时间戳必须不早于 Unix epoch，并能装入 48 位毫秒字段。源码已列出 B-tree 支持，但把 hash 索引支持标为未完成。扩展版本为 `0.0.0`，安装需要超级用户；作者说明它是学习练习，自己未曾用于生产。即使 Cargo 清单提供 PostgreSQL 13–18 构建特性，也应把这些成熟度信号视为部署边界。
