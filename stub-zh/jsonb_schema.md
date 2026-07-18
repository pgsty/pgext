## 用法

来源：

- [已复核 commit 的 jsonb_schema README](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/README.md)
- [已复核 commit 的 jsonb_schema 1.0 安装 SQL](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/jsonb_schema--1.0.sql)
- [已复核 commit 的 jsonb_schema 回归 SQL](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/sql/test.sql)

1.0 版 `jsonb_schema` 把重复的 JSONB 结构与数据值分开存储。`jsonb_pack` 将去重后的二进制模式写入 `jsonb_schemes`，并返回包含模式 ID 与编码数据的两元素 JSONB 数组；`jsonb_unpack` 则还原原始值。

```sql
CREATE EXTENSION jsonb_schema;

CREATE TABLE packed_documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload jsonb NOT NULL
);

INSERT INTO packed_documents (payload)
VALUES (jsonb_pack('{"kind":"event","value":42}'::jsonb));

SELECT id, jsonb_unpack(payload)
FROM packed_documents;

SELECT count(*) FROM jsonb_schemes;
```

### 注意事项

- 打包值是扩展专用的数组表示，不能像原对象一样透明查询；应用需要先解包，再执行常规 JSONB 访问。
- 新模式出现时，`jsonb_pack` 会执行插入。即使安装 SQL 将函数标记为 parallel safe，也应把它视为写操作。
- 函数以未限定模式名的 `jsonb_schemes` 查表。如果扩展安装模式不在当前搜索路径中，调用可能失败或解析到非预期关系；应固定受控搜索路径并测试所选扩展模式。
- 必须把 `jsonb_schemes` 与所有打包列一起保留和备份；引用的模式行丢失后，打包值将无法还原。
