## 用法

来源：

- [1.0 版本 README](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/README.md)
- [1.0 版本 SQL 声明](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity--1.0.sql)
- [PostgreSQL 与 ICU 包装层](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity_pg.c)
- [ICU 构建依赖](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/Makefile)
- [扩展 control 文件](https://github.com/urbic/postgresql-similarity/blob/ddfb6c2114240ef3421e0e91cf7a40455f15e155/similarity.control)

`similarity` 1.0 使用 fstrcmp 风格的近似匹配算法比较两个字符串，返回 `double precision` 归一化分数：完全不同为 0，完全相同为 1。它适用于模糊候选比较，不是语言学搜索或带索引的最近邻检索。

### 核心工作流

```sql
CREATE EXTENSION similarity;

SELECT similarity('similarity', 'distinction');

SELECT name, similarity(name, 'postgresql') AS score
FROM projects
ORDER BY score DESC
LIMIT 10;
```

扩展只创建两个函数：

- `similarity(text,text)` 执行完整比较，等价于零截止值。
- `similarity(text,text,double precision)` 接收最小分数；一旦候选无法达到该截止值，就可以提前停止。

### 截止值语义

只有三参数函数达到请求的截止值时，结果才是精确值。如果返回值低于截止值，上游明确说明该数值无效，但可以确定它低于阈值。应将截止形式用于是否通过的剪枝优化，或与当前最佳分数搭配；需要对所有返回数值排名时，应使用 `similarity(text,text)`。

两个函数都声明为 `STRICT` 和 `IMMUTABLE`。它们不定义运算符、索引操作符类、选择率估算器或搜索索引。除非其他谓词先缩小候选集，对每行调用函数仍然是 CPU 与内存扫描。

### 编码、依赖与兼容性边界

C 包装层通过 ICU 的 UTF-8 转换函数处理 PostgreSQL 文本，并链接 `libicu`。应使用 `UTF8` 数据库，并测试多语言、组合字符、空串和超长输入。实现无法健壮地报告 ICU 转换错误，相似度也基于字符序列算法，而不是区域排序规则或语义语言规则。

1.0 没有后续发行系列或当前 PostgreSQL 兼容矩阵。应保持 ICU 运行时与二进制兼容，并在准确的 PostgreSQL 大版本上验证构建与回归结果。近似字符串比较的计算量与内存会随两个字符串增长，因此应限制不可信调用者的输入长度。
