## 用法

来源：

- [pg_steadytext/README.md](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/README.md)
- [pg_steadytext/sql/pg_steadytext--2025.11.25.sql](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/sql/pg_steadytext--2025.11.25.sql)
- [pg_steadytext/Makefile](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/Makefile)
- [pg_steadytext/pg_steadytext.control](https://github.com/julep-ai/steadytext/blob/e4d877620a6801de9705e362eda8d3a4f80cc5d1/pg_steadytext/pg_steadytext.control)

`pg_steadytext` 通过 PL/Python 提供 SteadyText 文本生成和嵌入能力，并附带 SQL 缓存、异步队列与提示模板。扩展版本为 2025.11.25，与 Python 库版本分别管理。

### 核心用法

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION vector;
CREATE EXTENSION pg_steadytext;
SELECT steadytext_generate('Write a haiku about PostgreSQL');
SELECT steadytext_embed('PostgreSQL is a database');
SELECT * FROM steadytext_cache_stats();
```

### 运行边界

上游要求 PostgreSQL 14+、Python 3.10+、`plpython3u` 和 `vector`。Python 依赖必须安装到 PostgreSQL 链接的解释器中，而不仅是终端使用的解释器。虽然 control 设置为 superuser=false，创建非受信语言函数仍需要超级用户。对象固定在 `public`；Makefile 安装 SQL 和 Python 文件，而非原生共享库。对已有数据使用前应检查安装 SQL，其中会以 CASCADE 删除 `steadytext_cache`。生成过程可能加载大型模型或连接已配置的守护进程。远程模型属于显式启用的不安全模式，可能将提示内容发送到数据库外。更高层的输入限制由应用负责。异步处理需要独立运行工作进程，可选的 `pg_cron` 可调度缓存淘汰。上游未要求共享预加载。
