

## 用法

> [pg_summarize](https://github.com/HexaCluster/pg_summarize)：使用 LLM 进行文本摘要，基于 pgrx 构建。
> 来源：[README.md](https://raw.githubusercontent.com/HexaCluster/pg_summarize/main/README.md)

`pg_summarize` 是一个用 Rust（使用 `pgrx`）编写的 PostgreSQL 扩展，集成了 OpenAI API。它包含一个基础的 "Hello, pg_summarize!" 函数和一个使用 OpenAI 模型进行文本摘要的 `summarize` 函数。


--------

### 快速开始

```sql
CREATE EXTENSION pg_summarize;

-- 测试 hello 函数
SELECT hello_pg_summarize();
--  hello_pg_summarize
-- ----------------------
--  Hello, pg_summarize
```


--------

### 配置

该扩展从 PostgreSQL 设置中获取配置。使用 `summarize` 函数前需设置以下参数：

```sql
-- 设置 OpenAI API 密钥（必需）
ALTER SYSTEM SET pg_summarizer.api_key = 'your_openai_api_key';

-- 可选设置模型（默认：gpt-3.5-turbo）
ALTER SYSTEM SET pg_summarizer.model = 'gpt-3.5-turbo';

-- 或在会话级别设置提示词
SET pg_summarizer.prompt = 'Your custom prompt here';

-- 如果在 SYSTEM 级别设置，需重新加载配置
SELECT pg_reload_conf();
```


--------

### Summarize 函数

`summarize` 函数接受文本输入，发送到 OpenAI API，并返回摘要：

```sql
-- 摘要文本输入
SELECT summarize('<This is the text to be summarized.>');

-- 从现有数据创建摘要表
CREATE TABLE blogs_summary AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;

-- 使用不同模型
SET pg_summarizer.model = 'gpt-4o';
CREATE TABLE blogs_summary_4o AS
  SELECT blog_url, summarize(blogs_text)
  FROM hexacluster_blogs;
```


--------

### 工作原理

- **配置获取**：`summarize` 函数通过 `current_setting()` 从 PostgreSQL 获取设置（API 密钥、模型、提示词）。未找到设置时使用默认值。
- **默认提示词**：内置提示词指示 AI 从 `<text>` 标签中摘要文本，重点简洁地捕获最重要的信息。
- **API 调用**：该函数向 OpenAI 聊天补全端点发送 POST 请求，包含配置的模型和提示词，返回摘要内容。
