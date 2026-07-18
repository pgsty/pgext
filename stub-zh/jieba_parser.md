## 用法

来源：

- [上游 README](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/README.md)
- [1.0 版安装 SQL](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/jieba_parser--1.0.sql)
- [C++ 与嵌入式 Python 实现](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/jieba_parser.cc)
- [Python 2.7 构建配置](https://github.com/seanlong/jieba_parser/blob/47cd212cc0c4455c5783d64f79d89391b445a48c/Makefile)

jieba_parser 1.0 是一个 2014 年的全文搜索解析器，它嵌入 Python，并调用 Python Jieba 包完成中文分词与关键词提取。

### 解析器配置

只能在可丢弃后端中评估经过修复并固定依赖的构建：

```sql
CREATE EXTENSION jieba_parser;
CREATE TEXT SEARCH CONFIGURATION jiebacfg (PARSER = jiebaparser);
ALTER TEXT SEARCH CONFIGURATION jiebacfg
  ADD MAPPING FOR word WITH simple;
SELECT to_tsvector('jiebacfg', '我来到北京清华大学');
SELECT jieba_extract_tags('清华大学位于北京', 3);
```

文本搜索配置是独立数据库对象；删除或升级扩展时必须将其纳入处理。

### 注意事项

- Makefile 固定使用 Python 2.7 与仅限 Python 2 的 API。Python 2 已终止支持，该源码无法不加修改地针对受支持 Python 3 版本构建。
- 共享库在全局对象构造期间初始化 Python 并导入 Jieba。导入与调用失败由 C 断言而非 PostgreSQL 错误处理；发布构建可能继续使用空指针并使后端崩溃。
- 两个全局对象都会执行解释器生命周期调用。它们无法与其他使用 Python 的扩展安全协调同一解释器的嵌入或结束。
- jieba_rank 接受向量与查询，却始终返回零。它不是可用的排序函数。
- 分词与关键词结果取决于另行安装的 Jieba 包、词典与模型版本。应固定这些资产以保证索引可重复，并在有意变更后重建受影响的文本搜索索引。
- 上游没有许可证文件或当前 PostgreSQL 兼容证据。应使用受维护解析器，而不是按发布状态部署该实现。
