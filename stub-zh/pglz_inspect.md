## 用法

来源：

- [pgrx-examples/pglz_inspect/pglz_inspect.control](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/pglz_inspect.control)
- [pgrx-examples/pglz_inspect/README.md](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/README.md)
- [pgrx-examples/pglz_inspect/Cargo.toml](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/Cargo.toml)
- [pgrx-examples/pglz_inspect/src/lib.rs](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/pgrx-examples/pglz_inspect/src/lib.rs)
- [LICENSE](https://github.com/pgcentralfoundation/pgrx/blob/fc91c63ebad11784647b50ee7e265c1fd9c9924f/LICENSE)

`pglz_inspect` 是 PGRX 中可用的示例扩展，可测量单个值或采样列的 PGLZ 压缩效果。其未发布的示例清单版本为 0.0.0。

### 核心用法

```sql
CREATE EXTENSION pglz_inspect;
SELECT * FROM pglz_size(convert_to(repeat('abc', 100), 'UTF8'));
CREATE TEMP TABLE compression_sample AS SELECT repeat('abc', 100) AS payload;
SELECT * FROM pglz_analyze_column('compression_sample'::regclass, 'payload', 1000, 'default');
SELECT pglz_recommend('compression_sample'::regclass, 'payload', 1000);
```

### 运行边界

`pglz_size` 返回原始／压缩字节数、比例及是否接受压缩；`pglz_analyze_column` 给出采样统计，`pglz_ratio_histogram` 对比例分组，`pglz_recommend` 返回建议。列采样使用 ORDER BY random() 查询，可能扫描和排序大量数据；估算取决于样本和规划器行数统计。函数遵循调用者的表访问权限，不会修改列存储。控制文件不限定超级用户安装，也未标记为可信；上游未要求预加载。README 说明支持 PostgreSQL 13–18，清单额外定义 PG19 构建选项，但这不等同于已验证支持。
