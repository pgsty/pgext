## 用法

来源：

- [已复核 commit 的 pg_sentence_transformer README](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/README.md)
- [已复核 commit 的 pg_sentence_transformer.control](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer.control)
- [版本 0.1 的安装 SQL](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer--0.1.sql)
- [后台工作进程实现](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/src/pg_sentence_transformer.c)
- [Python 模型实现](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pg_sentence_transformer/model.py)
- [Python 依赖元数据](https://github.com/serpent7776/pg_sentence_transformer/blob/02e32457e51813fe9d7c4895f0a7785c49e69735/pyproject.toml)

`pg_sentence_transformer` 通过后台工作进程，把一张已配置表中的文本转换为句子嵌入。它必须经由 `shared_preload_libraries` 加载，并依赖 `plpython3u`、`vector`、Python 虚拟环境及模型所需的 Python 软件包。

安装脚本会创建 `sentence_transformer` 模式、队列、源表触发器，以及 `prepare()` 和 `process()` 过程。已复核实现会加载 `all-MiniLM-L6-v2`，写入 `vector(384)` 值，并创建 `ivfflat` 索引。

### 配置

```conf
shared_preload_libraries = 'pg_sentence_transformer'
sentence_transformer.venv_path = '/opt/pg_sentence_transformer/venv'
sentence_transformer.database_name = 'appdb'
sentence_transformer.schema_name = 'public'
sentence_transformer.src_table_name = 'posts'
sentence_transformer.src_column = 'body'
sentence_transformer.id_column = 'id'
```

设置这些 postmaster 参数后重启 PostgreSQL，再连接到配置的数据库：

```sql
CREATE EXTENSION pg_sentence_transformer CASCADE;
```

应先把 Python 依赖安装到指定环境中。在网络受限环境部署前，还要准备好 `all-MiniLM-L6-v2` 模型和 NLTK 分词数据。

### 注意事项

- `prepare()` 会下载 NLTK 的 `punkt` 和 `punkt_tab` 数据集；模型文件没有缓存时，模型初始化也可能触发下载。应在普通查询流量之外完成并验证这些准备工作。
- 当前 `process()` 循环会把同一源值的每个句子都按相同 `ref_id` 执行 upsert；多句文本中，后面的句子会覆盖前面的嵌入。应把这视为原型行为，而不是文档级池化。
- 控制文件声明版本 `0.1` 且称扩展可重定位，但 SQL 会创建固定的 `sentence_transformer` 模式。工作进程每次服务器启动也只面向一个已配置数据库和表；必须在确切 PostgreSQL 构建上测试模式位置、重启行为、权限、Python ABI 和模型内存占用。
