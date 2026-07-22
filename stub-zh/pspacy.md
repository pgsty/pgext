## 用法

来源：

- [扩展 SQL](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy--1.0.sql)
- [Python 实现](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/pspacy.py)
- [固定的 Python 依赖](https://github.com/mikeizbicki/chajda/blob/a9912d87b6f9bf8dc6a92ace5395ed5f97b3df51/requirements.txt)

`pspacy` 1.0 是 chajda 仓库中的 PL/Python 扩展，把 spaCy 分词和词形还原转换为 PostgreSQL 全文检索值。它可以原型验证多语言检索规范化，但会在数据库后端执行老旧 Python 依赖栈，必须严格控制权限与资源。

### 核心流程

在 PostgreSQL 服务器环境中安装 `plpython3u` 以及准确的 Python 与原生依赖，再创建扩展并测试目标语言：

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pspacy;

SELECT spacy_lemmatize('en', 'The striped bats were hanging on their feet');
SELECT spacy_tsvector('en', 'The striped bats were hanging on their feet');
SELECT spacy_tsquery('en', 'striped bats');
```

每个后端首次调用时会通过 `pg_config` 获取 PostgreSQL 扩展目录，导入已安装的 Python 模块，并缓存在 PL/Python 会话状态中。语言流水线延迟加载；不支持的语言代码会退回多语言 `xx` 流水线。

### 函数索引

- `spacy_load()` 初始化会话模块缓存。
- `spacy_lemmatize(...)` 控制小写转换、特殊字符删除、停用词删除与位置输出。
- `spacy_lemmatize_query(...)` 用 AND 操作符连接词元，构造查询。
- `spacy_tsvector(lang, text)` 返回 `tsvector`。
- `spacy_tsquery(lang, text)` 把生成的查询文本交给 `simple` 文本搜索配置。

每个词元会截断到 500 个字符。部分处理错误会返回 NULL，因此必须明确监控索引失败。

### 安全与维护风险

`plpython3u` 是不受信语言，且若不修改权限，函数默认可由 PUBLIC 执行；应撤销权限，只授予受控检索角色。加载模型会消耗大量后端内存与 CPU。所有函数都声明为不可变且并行安全，但初始化会运行子进程和导入操作，结果还依赖外部 Python 包与语言数据；未经修正并验证这些易变性和并行标签，不要把函数用于表达式索引或存储生成列。固定依赖集合与自定义 spaCy 分支来自 2020 年前后生态，任何现代部署都应先隔离，并制定迁移计划。
