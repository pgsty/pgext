## 用法

来源：

- [openai.control](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/openai.control)
- [README.md](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/README.md)
- [openai--1.0.sql](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/openai--1.0.sql)
- [Makefile](https://github.com/pramsey/pgsql-openai/blob/83eedc5982d5e525e1cb7a0b31315fc75bd05382/Makefile)

`openai` 通过兼容 OpenAI 的端点提供模型列表、提示、嵌入和图像分析 SQL 函数，依赖 `http` 扩展。所核对项目为纯 SQL 实现，尽管控制文件仍包含库路径字段。

### 核心用法

```sql
CREATE EXTENSION http;
CREATE EXTENSION openai;
SET openai.api_uri = 'http://127.0.0.1:11434/v1/';
SET openai.api_key = 'none';
SELECT * FROM openai.models();
```

### 运行边界

安装需要特权，没有预加载要求。远程调用前需在会话中配置端点、API 密钥与模型。示例使用已运行的本地端点，并不会安装模型服务。`openai.models`、`openai.prompt`、`openai.vector` 和 `openai.image` 会发起外部请求，数据会离开数据库，延迟及错误影响当前查询。应保护凭据、限制 HTTP 超时并校验模型结果。未找到明确的项目许可证。
