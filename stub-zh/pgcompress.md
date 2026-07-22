## 用法

来源：

- [官方 README](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/README.md)
- [扩展 SQL API](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/pgcompress--1.0.sql)
- [压缩实现](https://github.com/chet0xhenry/pgcompress/blob/8264107a26c07b6415f75c4a8eb84f9d27e77ea3/pgcompress.c)

`pgcompress` 1.0 使用 zlib、原始 DEFLATE、gzip 或 Brotli，把 `text`、`bytea`、`json` 和 `jsonb` 压缩为 `bytea`，并提供配套解码函数。它是已经停止维护的旧 C 扩展；只应在应用确实需要这些外部线格式时使用，而不能替代 PostgreSQL TOAST 压缩。

### 核心流程

通用接口默认使用 zlib。提供算法代码后，解码时必须使用同一代码：

```sql
CREATE EXTENSION pgcompress;

SELECT decode_compressed(encode_compressed('hello PostgreSQL'));

SELECT decode_compressed(
  encode_compressed('hello PostgreSQL', 2, 6),
  2
);

SELECT decode_compressed_jsonb(
  encode_compressed('{"ok":true}'::jsonb, 3, 5),
  3
);
```

算法代码 `0` 表示 zlib，`1` 表示原始 zlib/DEFLATE，`2` 表示 gzip，`3` 表示 Brotli。zlib/gzip 级别范围为 0–9，Brotli 为 0–11。超出范围的类型会退回 zlib，无效级别会退回库默认值，因此存储前应验证参数，而不能依赖静默回退。

### 函数族

`encode_compressed` 和 `decode_compressed*` 通过代码选择算法。显式函数族包括 `encode_zlib`/`decode_zlib*`、`encode_zlib_raw`/`decode_zlib_raw*`、`encode_gzip`/`decode_gzip*` 以及 `encode_br`/`decode_br*`。`_bytea`、`_json`、`_jsonb` 等后缀选择解码后的 PostgreSQL 返回类型；编码结果始终为 `bytea`。

### 运维边界

模块必须链接 zlib 和 Brotli。压缩会在调用后端中同步运行，消耗 CPU 与内存；之后 PostgreSQL 还可能独立对结果进行 TOAST 压缩。应随每个载荷记录算法，因为压缩字节本身没有 PostgreSQL 级类型契约，错误解码器会报错或得到不可用数据。

不要在缺少应用大小限制时解压不可信、无界载荷：高压缩率输入可能膨胀到耗尽后端内存。仓库未确认对当前 PostgreSQL 大版本的兼容性，并使用服务器内部 C API，因此应针对精确目标版本重新构建和测试后再依赖已存数据。若目标只是透明降低数据库存储量，应优先使用内置列压缩。
