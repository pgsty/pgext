## 用法

来源：

- [IMCS 官方用户指南](https://pgxn.org/dist/imcs/user_guide.html)
- [固定提交的扩展 control 文件](https://github.com/knizhnik/imcs/blob/940964df4f3d03ea67ebe868691a17b41c89218b/imcs.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/imcs/)

`imcs` 为分析型时序负载提供内存列式存储。普通表仍是持久数据源；选定列会装入共享内存，生成的时序访问器利用向量化操作、压缩和可选的并行执行。

应在服务器启动时加载模块，并在修改配置后重启 PostgreSQL：

```conf
shared_preload_libraries = '$libdir/imcs'
```

对于时间戳键为 `Day`、序列标识为 `Symbol` 的 `Quote` 表，指南给出以下流程：

```sql
CREATE EXTENSION imcs;
SELECT cs_create('Quote', 'Day', 'Symbol');
SELECT Quote_load();
SELECT cs_max(Close) FROM Quote_get('IBM');
```

`cs_create()` 会为所选表生成 load、append、delete、trigger 和 accessor 对象。从表加载时会按时间戳排序；增量或触发器驱动的插入必须保持时间戳递增，否则会报告 out-of-order 错误。该存储不能直接表示 SQL `NULL`，支持的元素类型也以定长值为主；使用变长字符串前应复核 dictionary mapping。

服务器重启后必须重新加载共享内存内容，除非使用文档中的 autoload 模式。加载数据前应调整 `imcs.shmem_size` 及相关限制，并测试 autoload 的首次访问延迟。固定提交的 control 报告版本 `1.1`，而 PGXN 当前列出的最新分发版是 `0.1.7`；应把它们视为两条不同版本线，并核实准确源码、许可证、服务器兼容性和升级路径。
