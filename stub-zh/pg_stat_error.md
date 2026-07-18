## 用法

来源：

- [TencentDB for PostgreSQL 支持扩展矩阵](https://cloud.tencent.com/document/product/409/75121)

`pg_stat_error` 是 TencentDB for PostgreSQL 的服务商托管扩展。当前官方矩阵在 PostgreSQL 10 至 13 上列出扩展版本 1.0，在 PostgreSQL 14 至 18 上没有列出该扩展。

```sql
CREATE EXTENSION pg_stat_error;
```

只能在受支持的 TencentDB 实例上安装；创建前应根据实例引擎版本再次核对可用性。

### 服务商边界

腾讯云没有为该扩展发布源码仓库、control 文件、SQL 对象参考、许可证或可移植安装包。公开支持页面也没有说明可调用的函数或视图，因此不能仅凭名称推断错误统计 API。应以已安装版本暴露的对象和腾讯云支持说明为准；它不是自托管 PostgreSQL 组件。
