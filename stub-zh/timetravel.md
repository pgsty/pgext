## 用法

来源：

- [TencentDB for PostgreSQL 支持扩展矩阵](https://cloud.tencent.com/document/product/409/75121)

`timetravel` 是 TencentDB for PostgreSQL 的服务商托管扩展。当前官方矩阵仅在 PostgreSQL 10 和 11 上列出扩展版本 1.0。

```sql
CREATE EXTENSION timetravel;
```

只能在引擎版本明确列出该扩展的 TencentDB 实例上创建；新建或升级实例前应检查当前矩阵。

### 服务商边界

腾讯云没有为该组件发布公开源码、control 文件、SQL 对象参考、许可证、依赖列表或可移植包。特别是，支持矩阵不能证明它实现了历史 PostgreSQL `timetravel` contrib module 或任何同名第三方 API。应检查已安装对象并使用服务商支持文档，不能根据名称臆测函数。
