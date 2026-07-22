## 用法

来源：

- [官方 README](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/README.md)
- [1.0.0 扩展 SQL](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/sql/shared_ispell--1.0.0.sql)
- [扩展 control 文件](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/shared_ispell.control)

`shared_ispell` 把 Ispell 词典保存在固定共享内存段中，使 PostgreSQL 会话能够复用同一份已加载数据，而不必让每个后端分别解析和保存私有副本。它适合大型 Ispell 词典；Snowball 词典不会从中获益。

### 核心流程

选择共享内存预算，预加载模块并重启 PostgreSQL，然后创建扩展。

```conf
shared_preload_libraries = 'shared_ispell'
shared_ispell.max_size = '32MB'
```

```sql
CREATE EXTENSION shared_ispell;

CREATE TEXT SEARCH DICTIONARY czech_shared (
  TEMPLATE = shared_ispell,
  DictFile = czech,
  AffFile = czech,
  StopWords = czech
);

SELECT ts_lexize('czech_shared', 'automobile');
```

词典、词缀与停用词文件遵循 PostgreSQL 文本搜索文件放置规则。词典第一次被使用时才会载入共享内存段。

### 管理对象

- `shared_ispell_mem_used()` 与 `shared_ispell_mem_available()`：报告共享内存用量。
- `shared_ispell_dicts()` 与 `shared_ispell_stoplists()`：列出已加载词典和停用词表。
- `shared_ispell_reset()`：使已加载数据失效，让会话重新加载变更后的文件。
- `shared_ispell` 文本搜索模板：用于创建自定义词典。

共享内存分配不能在运行时增长。容量必须覆盖所有词典并留出余量；修改 `shared_ispell.max_size` 需要重启。上游警告，需要完整正则表达式的词缀不受支持，并会在词典初始化时失败。该项目已无人维护，因此生产使用前必须针对实际 PostgreSQL 构建验证兼容性。
