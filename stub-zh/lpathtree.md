## 用法

来源：

- [官方 README](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/README.md)
- [官方扩展 SQL](https://github.com/mediait/lpathtree/blob/6b1fd0acd1d4a698d9f4b04d59eb4e144a1eda15/lpathtree--1.0.sql)

`lpathtree` 是 PostgreSQL ltree 模块的一个分支，使用 `/` 分隔标签。它适合需要斜杠语法，或标签中需要 `{`、`}`、`!`、`*` 等字符的物化路径；代价是移除了原版的 `ltxtquery` 功能。

### 核心流程

使用 `lpathtree` 类型存储路径，并为祖先、后代和路径查询操作符建立 GiST 索引。

```sql
CREATE EXTENSION lpathtree;

CREATE TABLE taxonomy (path lpathtree PRIMARY KEY);
CREATE INDEX taxonomy_path_gist ON taxonomy USING gist (path);

INSERT INTO taxonomy(path) VALUES
  ('Top/Science'),
  ('Top/Science/Astronomy'),
  ('Top/Hobbies');

SELECT path
FROM taxonomy
WHERE 'Top/Science'::lpathtree @> path;

SELECT path
FROM taxonomy
WHERE path ~ 'Top/*/Astronomy'::lpathquery;
```

### 主要对象

- `lpathtree` 与 `lpathquery` 分别是路径和模式类型。
- `@>` 与 `<@` 测试祖先和后代关系；`~` 匹配 `lpathquery`；`||` 拼接路径。
- `nlevel`、`subpath`、`index` 与 `lca` 分别提供深度、切片、子路径搜索和最低公共祖先操作。
- 扩展为标量路径和路径数组提供 B-tree 排序及 GiST 操作符类。

### 兼容边界

这份代码分支自较旧的 PostgreSQL ltree 实现。斜杠分隔的值不能与点号分隔的 `ltree` 值直接互换，而且有意不提供 `ltxtquery`。从 `ltree` 迁移时应同时重写数据与查询，并在生产使用前针对准确的 PostgreSQL 大版本验证这个旧 C 扩展。
