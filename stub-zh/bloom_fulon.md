## 用法

来源：

- [官方扩展控制文件](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/bloom_fulon.control)
- [官方 Rust 实现](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/src/lib.rs)
- [官方 Rust 包清单](https://github.com/matheus-consoli/boom-index-pgrx/blob/7025b8186d770eee48ecff55f683c50355d98ff4/Cargo.toml)

`bloom_fulon` 是一个未完成的 `pgrx` 示例，用于注册 PostgreSQL 索引访问方法。版本 `0.0.0` 可用于学习访问方法回调，但它不是可用的 Bloom 索引，绝不能承载生产数据。

### 核心流程

创建扩展会注册 `bloom_fulon` 访问方法，以及面向整数等值比较的默认 `int4_ops` 操作符类：

```sql
CREATE EXTENSION bloom_fulon;

CREATE TABLE bloom_demo (id integer);
CREATE INDEX bloom_demo_idx ON bloom_demo USING bloom_fulon (id)
  WITH (size = 10);
```

`size` 关系选项默认值为 10，可接受 -10 到 100。它仅供演示；实现并不会用它构造真正的过滤器。

### 运维边界

构建回调报告零条索引元组，插入回调不会存储条目，清理回调也不执行维护。`amgettuple` 与 `amgetbitmap` 都未实现，因此该访问方法无法返回匹配行。其代价估算被固定为 `1e10`，以阻止规划器选用它。

控制文件要求超级用户安装，且扩展不可重定位。已核验清单为 PostgreSQL 11 到 16 定义了构建特性，但构建矩阵并不代表实现已经完整。只能在可丢弃的开发数据库中使用 `bloom_fulon` 进行访问方法实验。
