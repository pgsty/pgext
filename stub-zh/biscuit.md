
## 用法

> [README](https://github.com/CrystallineCore/Biscuit) | [Docs](https://biscuit.readthedocs.io/)

`biscuit` 是 PostgreSQL 的一种索引访问方法，专为 `LIKE` 和 `ILIKE` 模式匹配优化，也支持多列检索。上游将其定位为一种确定性的位图索引，可避免基于 trigram 的搜索常见的误命中复查开销。

### 快速上手

创建扩展，并在一个或多个文本列上建立 Biscuit 索引：

```sql
CREATE EXTENSION biscuit;

CREATE INDEX idx_users_name ON users USING biscuit(name);

CREATE INDEX idx_products_search
ON products USING biscuit(name, description, category);
```

带通配符的常见查询同样可以使用该索引：

```sql
SELECT * FROM users WHERE name LIKE '%john%';
SELECT * FROM users WHERE name NOT LIKE 'a%b%c';
SELECT COUNT(*) FROM users WHERE name LIKE '%test%';

SELECT *
FROM products
WHERE name LIKE '%widget%'
  AND description LIKE '%blue%'
  AND category LIKE 'electronics%'
LIMIT 10;
```

## 索引行为

Biscuit 为每个字符串维护位图位置索引，能够同时匹配正向和反向字符位置。上游设计强调：

- 正向索引，用于匹配字符在精确位置上的出现
- 反向索引，用于按字符串末尾倒数位置匹配字符
- `ILIKE` 的大小写不敏感变体
- 用于快速长度过滤的精确长度位图和最小长度位图

对于 `LIKE 'abc%def'` 这类模式，Biscuit 可以把前缀位图、后缀位图以及最小长度过滤合并起来，从而在不执行 heap 复查的情况下得到精确结果。

### 模式类型

上游文档对常见模式给出了优化路径：

- 精确匹配，例如 `'abc'`
- 前缀匹配，例如 `'abc%'`
- 后缀匹配，例如 `'%xyz'`
- 子串匹配，例如 `'%abc%'`
- 多列谓词，Biscuit 会按估计选择性重排谓词顺序

## 性能说明

上游 README 强调了纯位图求值及多项执行优化，包括：

- 中间位图为空时提前结束
- 对稀疏和稠密数据直接使用 roaring bitmap
- 后缀谓词使用反向位置查找
- 对 TID 做排序，以提高 heap 访问局部性
- 对聚合查询和 `LIMIT` 的特殊处理

项目 README 还给出了一个 100 万行测试表的基准方案，用来比较 Biscuit 索引与 trigram 方案。

## 需求

当前上游 README 列出的源码构建要求包括：

- PostgreSQL 16 或更高版本
- 标准构建工具，如 `gcc`、`make` 和 `pg_config`
- 可选的 CRoaring，用于提升性能

该项目通过 [PGXN](https://pgxn.org/dist/biscuit/) 发布包，并在 Read the Docs 上维护独立文档站。
