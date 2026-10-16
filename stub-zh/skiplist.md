## 用法

来源：

- [README.md](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/README.md)
- [extension/skiplist--1.0.sql](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/extension/skiplist--1.0.sql)
- [db/init_skiplist.sql](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/db/init_skiplist.sql)
- [extension/skiplist.control](https://github.com/pdmauricio/Knoo/blob/d6182b865c7b4af9f7bd6c3dca096b321c5e8978/extension/skiplist.control)

`skiplist` 1.0 是 Knoo PostgreSQL 18 课程项目中的 C 索引扩展，是具有已知正确性与持久性限制的教学实现。

### 核心用法

```sql
CREATE EXTENSION skiplist;
CREATE TABLE skiplist_example (k integer);
INSERT INTO skiplist_example VALUES (1), (2);
CREATE INDEX skiplist_example_idx ON skiplist_example USING skiplist (k);
SELECT skiplist_dump('skiplist_example_idx'::regclass);
```

### 运行边界

仅用于键值唯一且可丢弃的 `int4` 演示数据。上游明确说明只支持单列等值查询，重复键最多返回一条结果，不支持范围与排序，每个键占用一个 8 KB 页面，且最多只有四层。实现缺少自身 WAL、完整并发处理和专用 VACUUM。SQL 操作符类虽然声明了范围操作符，但不代表范围扫描已经实现。`skiplist_dump(regclass)` 用于输出结构以便调试。默认仅超级用户可安装，仓库未声明许可证。
