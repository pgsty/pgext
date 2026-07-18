## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/README.md)
- [编译后的扩展 SQL](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/tsp_semantic_headline--1.0.sql)
- [扩展 control 文件](https://github.com/thevermeer/pg_ts_semantic_headline/blob/c34c0fc2848ebb90590d20026f02eeb2db09d190/tsp_semantic_headline.control)

`tsp_semantic_headline` 是面向短语的全文高亮纯 SQL 实验。它提供 `TS_SEMANTIC_HEADLINE`、预处理查询/向量辅助函数，以及可利用对齐的预计算 `TSVECTOR` 和 `text[]` 列加速召回的 `TS_FAST_HEADLINE`。

```sql
CREATE EXTENSION unaccent;
CREATE EXTENSION tsp_semantic_headline;
SELECT TS_SEMANTIC_HEADLINE(
  'english',
  'phrase matches are highlighted as complete phrases',
  TO_TSPQUERY('english', 'phrase<->match')
);
```

索引文本、召回数组与查询必须使用相同预处理，否则 token 位置会偏离，高亮也会错误。快速路径用额外存储列和更新工作换取更低运行成本。性能声明应使用自己的语言、语料、查询形状和选项复现。

上游称项目仍在开发，并记录多个未实现选项。其特殊字符处理会插入分隔符，还明确警告基于空格的假设会改变或破坏多种语言的含义。高亮结果只应作为展示文本，必须针对目标 UI 转义，绝不能把不可信 HTML 标记为安全。应测试词干、否定、短语距离、标点、Unicode、XSS、长文档和内容更新后的偏差。
