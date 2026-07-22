## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/README.md)
- [1.0 版本 SQL 声明](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index--1.0.sql)
- [数组实现](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.c)
- [扩展 control 文件](https://github.com/rlichtenwalter/pg_array_multi_index/blob/d315364e110db6ec50e822e2eb5957d988d5a5a6/pg_array_multi_index.control)

`pg_array_multi_index` 提供 `array_multi_index(anyarray, integer[])`，在一次 C 调用中选择多个不连续的数组元素，并返回元素类型相同的数组。当应用已经持有位置列表，希望避免为多个独立下标反复扫描可变长数组时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION pg_array_multi_index;

SELECT array_multi_index(
  ARRAY['zero', 'one', 'two', 'three', 'four'],
  ARRAY[2, 5, 2, NULL, 99]
);
-- {one,four,one,NULL,NULL}
```

位置按数组扁平存储顺序的一基偏移解释。结果保持请求顺序，重复位置会重复值，NULL 或越界位置会生成 NULL 结果元素。空索引数组返回空数组。由于 SQL 函数是 strict，值数组或索引数组为 NULL 时不会调用函数。

### 形状与兼容性边界

无论源数组是否有自定义下界或多个维度，结果总是一维且下界为 1。源数组的下界元数据会被忽略：位置 1 表示解构后的第一个物理元素，不一定是源数组声明的下标 1。如果原始维度或下界具有业务语义，应使用内置下标逻辑。

1.0 版本使用旧式 `anyarray` 多态签名和 PostgreSQL 内部数组 API。仓库没有发行标签、兼容矩阵或近期维护记录。依赖它之前，应针对确切服务器大版本构建与测试，并覆盖按值、按引用、含 NULL、toasted、多维及自定义下界数组。
