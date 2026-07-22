## 用法

来源：

- [官方扩展 SQL](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer--1.0.sql)
- [官方 C 实现](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer.c)
- [官方树形格式化实现](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer_utils.c)

`jsonb_explorer` 是一个小型 C 扩展，可将 JSONB 文档渲染为缩进文本树。版本 `1.0` 只公开两个函数，其中一个存在源码级返回类型缺陷；仅应在可丢弃的检查会话中使用可工作的树形渲染函数。

### 核心流程

安装扩展，并使用 `jsonb_tree(jsonb)` 渲染文档：

```sql
CREATE EXTENSION jsonb_explorer;

SELECT jsonb_tree(
  '{"service":{"name":"api","ports":[8080,8081]}}'::jsonb
);
```

结果是供人阅读的展示文本，其中包含树形字符、键、标量值与数组大小说明，不应作为稳定的机器解析格式。

### 重要对象

- `jsonb_tree(jsonb)` 以 `text` 返回格式化树；该函数声明为 strict，因此拒绝 NULL 输入。
- `jsonb_paths(jsonb)` 声明返回 `text[][]`，但 C 实现构造并返回的是单个 `text` datum。该 ABI/类型不匹配使函数不应被调用。

### 运维注意事项

除非维护中的下游构建同时修正了 SQL 声明与实现，否则不要调用 `jsonb_paths(jsonb)`。仓库没有实质性 API 文档或兼容矩阵，最后一次复核的提交来自 2018 年。格式化器还使用 PostgreSQL 内部 JSONB 迭代器 API，因此二进制必须针对确切的 PostgreSQL 大版本重新构建并测试。应用逻辑应优先使用内置 JSONB 操作符、`jsonb_each`、`jsonb_array_elements` 或仍在维护的客户端格式化器。
