## 用法

来源：

- [已复核 commit 的 pg_sentence_transformer README](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/README.md)
- [版本 0.1 的安装 SQL](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer--0.1.sql)
- [后台工作进程实现](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.c)
- [Python 模型实现](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer/model.py)
- [扩展 control 文件](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.control)

`pg_sentence_transformer` 是一个原型后台工作进程，在 PostgreSQL 进程内运行 Hugging Face sentence-transformer 模型。它监视一张配置好的文本表，把插入或发生变化的行加入队列，再把 `vector(384)` 嵌入写入相邻表。它省去了外部推理服务，但也把 Python、模型内存、下载和推理失败带入数据库服务器边界。

它依赖 `plpython3u`、`vector`、扩展的 Python 软件包及依赖项，并且必须通过 `shared_preload_libraries` 加载。

### 配置与启用

所有扩展设置都是仅超级用户可修改的 `PGC_POSTMASTER` 参数，因此修改后必须重启服务器。

```conf
shared_preload_libraries = 'pg_sentence_transformer'
sentence_transformer.venv_path = '/opt/pg_sentence_transformer/venv'
sentence_transformer.database_name = 'appdb'
sentence_transformer.schema_name = 'public'
sentence_transformer.src_table_name = 'posts'
sentence_transformer.src_column = 'body'
sentence_transformer.id_column = 'id'
```

先把 Python 环境安装在 PostgreSQL 操作系统用户可读的位置，重启服务器，再到配置的数据库中创建扩展。

```sql
CREATE EXTENSION pg_sentence_transformer CASCADE;
```

工作进程启动时，`sentence_transformer.prepare` 会创建类似 `public.posts_embeddings` 的存储、`ivfflat` 索引、入队与删除触发器，以及初始队列行。`sentence_transformer.process` 使用 `FOR UPDATE SKIP LOCKED` 取得一个队列行，计算并向上插入嵌入，再删除队列项。

### 固定模型与原型语义

Python 模块硬编码 `all-MiniLM-L6-v2`，其输出保存为 `vector(384)`。`sentence_transformer.prepare` 会下载 NLTK 的 `punkt` 和 `punkt_tab` 数据；模型初始化在缓存缺失时也可能下载模型文件。在网络受限或生产服务器上启用工作进程前，应预置并验证所有产物。

处理器把一条源文本拆成多个句子，再把每个句子都向上插入同一个 `ref_id`。因此后面的句子会覆盖前面的嵌入；版本 `0.1` 没有实现文档池化，也没有为每个句子保留一行。

### 运维注意事项

- 预加载会为一个配置好的数据库、schema、表、源列和整数 ID 列启动一个工作进程。一组服务器配置不能管理任意多张表。
- control 文件声称扩展可重定位，但 SQL 硬编码了 `sentence_transformer` schema，应把它视为固定 schema。
- schema、表和列设置以动态 SQL 字符串插值，而不是作为标识符引用。应只使用简单可信的标识符，并在开放源表写入前验证生成对象。
- 模型推理与非可信 Python 在 PostgreSQL 工作进程中运行。应测量常驻内存、CPU 争用、重启循环、队列增长、Python ABI 兼容性、模型缓存所有权和故障恢复。
- 工作进程可能在扩展 SQL 存在前启动，并配置为失败后重启。应在受控维护窗口中编排预加载、重启、依赖创建和扩展创建，并在之后确认工作进程恢复健康。
- 更新文本列会入队，删除会移除嵌入，但只修改 ID 列不在触发器更新条件内。依赖引用新鲜度前，必须测试应用的所有变更路径。
