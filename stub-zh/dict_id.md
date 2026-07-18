## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/README.md)
- [扩展 control 文件](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/dict_id.control)
- [词典实现](https://github.com/mohangk/dict_id/blob/d331b60e4bba5a5fedd79060f56d0ed88331955e/dict_id.c)

`dict_id` 是一个 alpha 阶段的 PostgreSQL 印尼语全文检索词干词典，底层使用 cSastrawi 词干提取器。上游明确说明它只适合参与开发者使用，并且文档针对 PostgreSQL 9.4.4。

```sql
CREATE EXTENSION dict_id;
SELECT ts_lexize('dict_id', 'bertapa');
```

示例返回词干 `tapa`。要用于搜索，应创建独立文本搜索配置，并在代表性印尼语语料验证后只映射选定 token 类型；词干规则会直接改变索引词元与查询命中。

应把它视为历史实验代码。已复核材料没有当前的语言质量、线程安全、内存安全或 PostgreSQL 兼容性声明。应与持续维护的印尼语搜索管线比较输出，测量误报与漏报，并在词典行为变化后重建所有依赖的全文索引。不要静默替换生产词典，否则新旧索引项可能出现语义不一致。
