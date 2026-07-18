## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/README.md)
- [1.0 版本 SQL 对象](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus--1.0.sql)
- [C 实现](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus.c)
- [扩展 control 文件](https://github.com/harukat/pg_pageinspect_plus/blob/5e7970fe57b80a909b0d6d65b1de52bf4eb5156f/pg_pageinspect_plus.control)

`pg_pageinspect_plus` 为 `pageinspect` 增加辅助函数，将原始元组属性解码为常用 PostgreSQL 类型。它安装在 `pg_catalog`，当前上游文档记录支持 PostgreSQL 10–18。

```sql
CREATE EXTENSION pageinspect;
CREATE EXTENSION pg_pageinspect_plus;
SELECT tuple_data_parse(
         'public.sample'::regclass,
         tuple_data_split('public.sample'::regclass,
                          t_data, t_infomask, t_infomask2, t_bits)
       )
FROM heap_page_items(get_raw_page('public.sample', 0));
```

其他辅助函数可把原始 `bytea` 转换为 timestamp、timestamptz、interval、integer、bigint、boolean、text、float4 与 float8。它们暴露的是物理表示，并非可移植序列化格式。

页面检查绕过正常行可见性与类型访问路径，可能暴露已删除或敏感值。执行权应仅授予可信管理员。结果依赖准确的表描述符、PostgreSQL 构建、字节序、对齐与页面状态；调查损坏时应在副本上操作，绝不能依据解码结果写入原始页面。
