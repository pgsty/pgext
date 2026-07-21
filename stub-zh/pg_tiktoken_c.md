## 用法

来源：

- [pg_tiktoken_c 包装版本的 README](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/README.md)
- [1.1 版本 SQL API](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/sql/pg_tiktoken_c--1.1.sql)
- [扩展控制文件](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/pg_tiktoken_c.control)
- [捆绑词汇数据](https://github.com/relytcloud/pg_tiktoken_c/tree/fa2957b6ece483322f4c4ce0c374b3b77e22b892/data)

pg_tiktoken_c 在 PostgreSQL 内部实现了与 OpenAI 兼容的 tiktoken 编码。使用它来计算或在存储文本附近生成令牌计数，并在嵌入或模型请求之前将文本分割为以令牌边界分隔的片段。

### 创建扩展

    CREATE EXTENSION pg_tiktoken_c;

实现依赖于构建时 PCRE2 10.30 或更高版本。它不需要 shared_preload_libraries；词汇数据按后端加载并缓存，当使用编码时。

### 编码和计数

    SELECT tiktoken_encode('cl100k_base', 'PostgreSQL 搜索');
    SELECT tiktoken_count('cl100k_base', 'PostgreSQL 搜索');

tiktoken_encode 返回一个 bigint 数组的令牌标识符。tiktoken_count 返回令牌计数，无需调用者保留令牌数组。

捆绑的选择器包括 cl100k_base、o200k_base、r50k_base、p50k_base 和 p50k_edit 以及由项目文档说明的别名。选择目标模型所需的编码而不是假设所有模型共享词汇表。

### 分割文本

以数组形式返回片段：

    SELECT chunk_text(
      '长文档文本',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

或每行返回一个片段：

    SELECT *
    FROM chunk_text_table(
      '长文档文本',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

chunk_text_table 返回 chunk_index、chunk 和 token_count。chunk_index 是零基索引。重叠会在相邻片段之间重复边界令牌，并且必须小于片段大小。

### 函数索引

- tiktoken_encode(selector, text) 返回 bigint[] 令牌标识符。
- tiktoken_count(selector, text) 返回 bigint 令牌计数。
- chunk_text(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) 返回 text[]。
- chunk_text_table(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) 每个片段返回一行，包含其索引和令牌计数。

SQL 函数声明为不可变且并行安全。因此，当所选词汇文件在每个服务器上一致部署时，它们可以用于生成表达式或并行计划。

### 运营注意事项

- 词元化是模型编码特定的。确认应用程序中的编码名称以及目标模型当前上下文限制。
- 计算或分割大文本会消耗后端 CPU 和内存；批量处理大型语料库并监控查询延迟。
- 后端本地缓存避免重复解析，但在涉及多个词汇表的会话中增加了内存使用量。
- 上游 README 的兼容性列表可能滞后于打包。直接测试具体的 pg_tiktoken_c 构建版本与目标 PostgreSQL 主要版本之间的兼容性，而不是从不同的二进制文件推断支持情况。
