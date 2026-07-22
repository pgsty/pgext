## 用法

来源：

- [官方 database.dev 软件包页面](https://database.dev/langchain/embedding_search)

`embedding_search` 是 database.dev 软件包 `langchain@embedding_search` 的规范目录名称。它安装 LangChain 向量存储示例所需的 Supabase/PostgreSQL 对象，包括由 `vector` 支撑的文档存储和相似度匹配。registry 最新软件包发布为 1.1.1，但公开的扩展安装示例和默认扩展版本仍为 1.1.0。

### 安装软件包

用 `dbdev.install` 安装带命名空间的软件包，显式创建依赖，并在 `CREATE EXTENSION` 中使用带命名空间的扩展名，而不是 `embedding_search`：

```sql
SELECT dbdev.install('langchain@embedding_search');
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION "langchain@embedding_search"
  SCHEMA public
  VERSION '1.1.0';
```

文档所示流程不会自动解析依赖，因此数据库中必须已经可以使用 `vector`。

### 核心流程

软件包会创建与所链接 LangChain/Supabase 集成相符的 `documents` 表和 `match_documents` 函数。客户端为文档生成嵌入，将内容与元数据存入 Supabase，再调用匹配函数进行向量相似度搜索；示例还展示了元数据过滤。

应使用选定软件包发布实际安装的表和函数签名，不要假定其他发布的示例完全一致。软件包页面所示 LangChain 导入路径来自较早的 JavaScript 客户端世代，因此客户端代码应与应用使用的 LangChain 版本对应。

### 安全与运维

示例使用 Supabase service-role key。该密钥与嵌入服务凭据必须保存在服务端密钥存储中，不得写入浏览器代码或 SQL 文件。对 `documents` 和 `match_documents` 应用行级安全与最小权限授权，明确谁可以写入嵌入或读取元数据，并以生产规模数据测试索引选择、嵌入维度、过滤选择性和迁移行为。
