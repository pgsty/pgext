## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/marcmunro/pgbitmap/blob/58c425172ef564ab032c4f3e73da19a01759be8b/README.md)
- [0.9.5 版本源码发行包](https://api.pgxn.org/src/pgbitmap/pgbitmap-0.9.5/)
- [生成的 API 文档](https://marcmunro.github.io/pgbitmap/docs/html/index.html)

`pgbitmap` 定义紧凑、非稀疏的 `bitmap` 类型，用于表示非负整数位置集合。它支持增删/测试位、最小最大值、并集、交集、差集、数组转换、迭代、比较和聚合。

```sql
CREATE EXTENSION pgbitmap;
SELECT to_bitmap(ARRAY[1, 2, 1000000]) AS members;
SELECT bitmap(1) + 2 + 3 AS members;
SELECT * FROM bits(to_bitmap(ARRAY[2, 5, 8]));
```

该类型用于避免保存高位之前的所有零。文本表示紧凑，适合往返转换，不面向人工阅读。应验证负数和极端索引、空值、聚合、排序以及跨架构转储恢复。

上游把 0.9.5 称为 beta，而目录标记项目已放弃。它部分用于权限集合，但 bitmap 本身不是授权系统：行访问与修改权限仍需独立执行。应固定二进制和文本格式，测试宽或碎片化集合的 TOAST 与内存行为；表示变化后必须重建依赖数据。重视可移植性时，应优先使用持续维护的核心或应用集合表示。
