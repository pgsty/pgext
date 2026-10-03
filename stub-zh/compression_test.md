## 用法

来源：

- [Official README](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/README)
- [Control 1.0](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test.control)
- [SQL 1.0](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test--1.0.sql)
- [Implementation](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/compress_test/compression_test.c)

`compression_test` 1.0 提供实验 pglz 压缩和读取关系原始页面的后端工具，属于 Michael Paquier 插件集合中的开发者模块。模块 README 声明兼容 PostgreSQL 9.5 及以后版本，但实现依赖服务器内部接口，构建时仍需核对目标大版本。

### 基本用法

由超级用户安装。使用独立 schema 可以避免通用函数名与其他对象冲突：

```sql
CREATE SCHEMA compression_lab;
CREATE EXTENSION compression_test WITH SCHEMA compression_lab;

WITH input AS (
  SELECT convert_to(repeat('abc', 100), 'UTF8') AS original
), packed AS (
  SELECT original, compression_lab.compress_data(original) AS compressed
  FROM input
)
SELECT compression_lab.bytea_size(original) AS original_bytes,
       compression_lab.bytea_size(compressed) AS compressed_bytes
FROM packed;
```

这个例子只比较字节数。当算法无法有效压缩时，函数会直接返回原始字节，且没有标记来区分这种情况。应保留原始数据与长度，不能将每个输出都视为可解压的数据帧。

### 函数索引

- `compress_data(bytea)` 使用内置的始终尝试压缩策略。七参数重载还接受最小与最大输入大小、最小压缩率、首次成功阈值以及两个匹配搜索控制值，适用于底层实验。
- `decompress_data(bytea, smallint)` 的第二个参数是原始未压缩长度，声明类型将其限制为正的有符号 16 位整数。非法数据或未压缩输入可能失败；它不是自描述归档格式。
- `bytea_size(bytea)` 返回不含 varlena 头部的有效载荷字节数。
- `get_raw_page(oid, int4, bool)` 返回页面字节和空洞偏移，函数内部检查超级用户权限。最后一个参数为 false 时移除页面空闲空间，为 true 时保留整页。不要依靠后一个分支清除未使用字节：当前实现是在复制返回页面之后才清零本地缓冲区。

### 使用边界

原始页面暴露的是物理存储内容，不是经过 MVCC 可见性过滤的查询结果。仅使用受控测试数据，并保留 schema 的访问限制。这些工具不会开启表压缩，也不定义稳定、可移植的持久化格式，不需要后台工作进程或预加载。
