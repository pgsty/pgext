## 用法

来源：

- [官方 README](https://github.com/za-arthur/pg_textparser/blob/master/README.md)
- [扩展控制文件](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser.control)
- [1.0 版扩展 SQL](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser--1.0.sql)
- [1.0 版 C 实现](https://github.com/za-arthur/pg_textparser/blob/master/pg_textparser.c)
- [1.0 版扫描器源码](https://github.com/za-arthur/pg_textparser/blob/master/textscan.l)

pg_textparser 是一个已废弃且未完成的 PostgreSQL 全文搜索解析器原型。目录中的 1.0 版源码没有提供可用解析器，应视为开发参考源码，而不是可部署扩展。

### 上游状态

扩展 SQL 声明了解析器回调函数，但解析器定义引用了 SQL 没有创建的标题回调，也没有把已声明的词法类型回调放入解析器定义。因此，按原样创建扩展并不完整。

启动、取下一标记、结束、词法类型和标题等 C 回调全部返回空值，没有维护解析器状态或输出标记。独立的扫描器原型只识别数字，并返回类似结束标记的值，而且没有集成到回调中。

### 预期验证界面

维护者补全 SQL 声明与回调实现后，应首先使用 PostgreSQL 的解析器检查函数作为验收门槛：

```sql
SELECT * FROM ts_token_type('textparser');
SELECT * FROM ts_parse('textparser', 'alpha 123 beta');
```

真正的实现必须返回稳定的标记类型目录，在不发生死循环的情况下消费输入，保持正确的字节长度和多字节边界，并产生可映射到文本搜索配置的标记。标题行为也需要独立测试。

### 部署边界

不要创建依赖此源码快照的生产索引或搜索配置。编译成功本身不能证明解析器正确；即使本地修改让 SQL 可以安装，仍需针对分词、空输入、空白、标点、编码、长值和标题生成进行回归测试。

控制文件把扩展标记为可重定位，也没有声明预加载或重启要求，但这不会改变其未完成的生命周期状态。在权威维护版本提供一致的 SQL 界面、有效回调和测试之前，应保持隔离。
