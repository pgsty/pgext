## 用法

来源：

- [官方扩展 SQL](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa--1.0.sql)
- [官方 C 实现](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.c)
- [官方扩展控制文件](https://github.com/akorotkov/pg_aa/blob/4ae0a182dfd66e65e18ef8faf460973c8706ac44/pg_aa.control)

`pg_aa` 在 PostgreSQL 内把以 `bytea` 保存的 PNG 图像转换成文本画。它提供一个由 aalib 支持的单色 ASCII 渲染器，以及一个由 libcaca 支持的 UTF-8 渲染器；这是演示扩展，并非图像存储或文档处理系统。

### 核心流程

把原始 PNG 字节存入表中或作为查询参数绑定，然后选择输出宽度：

```sql
CREATE EXTENSION pg_aa;

CREATE TABLE images (
  id bigint PRIMARY KEY,
  png bytea NOT NULL
);

SELECT aa_out(png, 80)
FROM images
WHERE id = 1;

SELECT caca_out(png, 80)
FROM images
WHERE id = 1;
```

输出高度根据原始宽高比和请求宽度推导，并对终端字符比例做了调整。

### 函数索引

- `aa_out(bytea, int4)` 使用 aalib 返回由换行分隔、与区域设置无关的 ASCII。
- `caca_out(bytea, int4)` 返回 libcaca 生成的 UTF-8 文本画。

两个函数都声明为不可变且严格。输入必须是可解码的 PNG；其他格式会报错。

### 运维边界

服务端二进制必须链接兼容版本的 libgd、aalib 和 libcaca。解码和栅格处理在数据库后端中同步执行，因此应限制图像大小与输出宽度，防止 CPU 或内存消耗过大。本次复核的源码来自 2015 年，没有 README、兼容矩阵或尺寸保护；把函数开放给不可信调用者前，应测试损坏 PNG、零/负宽度、极端宽高比、Unicode 客户端编码、库升级和并行调用。
