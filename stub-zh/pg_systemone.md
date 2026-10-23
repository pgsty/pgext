## 用法

来源：

- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/README.md)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/pg_systemone.control)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/pg_systemone--1.0.sql)
- [Official source](https://github.com/determinate/pg_systemone/blob/bdccd600eb458eb30018db6ebc6d1ee6ae828c78/src/pg_systemone.c)

`pg_systemone` 通过 C 语言 SQL 函数调用 SystemOne 决策模型。它向独立 HTTP 服务发送请求，并不在 PostgreSQL 中内嵌模型。上游文档列出 PostgreSQL 14–17，运行需要 libcurl 和兼容的服务端。

### 核心用法

```sql
CREATE EXTENSION pg_systemone;
SET systemone.api_url = 'http://127.0.0.1:9931/v1/systemone';
SELECT systemone_prob('The shipment arrived damaged.', 'Does this need customer support?');
```

`systemone_choice(state, instructions, criteria)` 返回分类结果；`systemone_prob(state, question)` 返回概率；`systemone_score(state, question, levels)` 对有序标签评分；`systemone_match_score(job_desc, resume)` 比较职位描述与简历。这些都是模型输出，并非确定事实。

### 配置与权限

`systemone.api_key` 提供可选的 bearer token。`systemone.batch_size` 目前只是为将来批处理预留的提示，不保证请求会合并发送。旧的 `jev_*` 函数与 `jev.*` 配置是兼容别名；这是独立的 C 扩展，与基于 Python 的 Jev 扩展不同，但两者的 SQL 名称可能冲突。

安装需要超级用户；文档没有要求预加载。扩展可迁移到其他模式。HTTP 调用在后端同步执行，会传输传入的文本。端点与凭据配置允许用户修改，安装脚本也未撤销函数的公共执行权限：向不可信角色开放前，应限制函数权限和网络出口。此修订中的 README 授权徽章没有对应的许可文件，目录因此保留许可未核实状态。
