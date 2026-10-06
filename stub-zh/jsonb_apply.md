## 用法

来源：

- [README.md](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/README.md)
- [sql/jsonb_apply--0.1.0.sql](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/sql/jsonb_apply--0.1.0.sql)
- [jsonb_apply.control](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/jsonb_apply.control)

`jsonb_apply` 调用指定的 PostgreSQL 文本函数，递归转换 JSONB 中的字符串，保留文档结构与非字符串值。

### 核心用法

```sql
CREATE EXTENSION jsonb_apply;
SELECT jsonb_apply('{"name":"Alice","tags":["hello"]}'::jsonb, 'upper');
SELECT jsonb_apply('{"message":"hello"}'::jsonb, 'replace', 'hello', 'bye');
```

### 运行边界

参数为 JSONB 文档、函数名称和可选的可变参数。被调用函数须以 text 作为首个参数并返回 text，其他参数类型参与重载解析。这里会调用 SQL 函数，而非固定转换列表，因此应使用经过审查的函数名，并控制模式可见性。

安装 C 库后由超级用户创建扩展；上游未要求预加载或重启。所核对的源码未声明 PostgreSQL 主版本支持矩阵或许可证，目录不推断这些信息。
