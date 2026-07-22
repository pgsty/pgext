## 用法

来源：

- [官方扩展控制文件](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy.control)
- [官方扩展 SQL](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/advcopy--0.0.2.sql)
- [官方 README](https://github.com/satya10x/advcopy/blob/d7dfcb6fafca0ea0c4ad7b67fa06df662fec6fbc/README.md)

`advcopy` 0.0.2 是一个实验性的 PL/Python 扩展，用于在 PostgreSQL 服务器、兼容 Amazon S3 的对象存储和可通过 `scp` 访问的主机之间移动 CSV 数据。所有文件、网络和 shell 操作都发生在数据库服务器上，而不是 SQL 客户端上。

### 核心流程

需要安装非受信语言 `plpython3u`，并确保 PostgreSQL 实际使用的 Python 解释器能够导入 Python `boto3` 模块。默认只有超级用户能够创建或执行非受信 PL/Python 函数。

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION advcopy;

CREATE TABLE public.import_target (id bigint, payload text);

SELECT advcopy.import_from_s3(
    'public.import_target',
    's3://example-bucket/incoming/data.csv'
);
```

`import_from_s3` 会把对象下载到数据库主机，然后使用服务器端 `COPY` 导入指定表。AWS 凭据与区域配置由 PostgreSQL 服务器进程环境中的 `boto3` 解析。

### 面向用户的函数

- `advcopy.import_from_s3(table_name text, endpoint_url text)` 下载一个 S3 对象并按 CSV 导入。
- `advcopy.import_to_s3(query text, file_name text, endpoint_url text)` 设计用于导出查询结果并上传到 S3。
- `advcopy.export_to_ip(query text, file_name text, ip text, folder text)` 设计用于导出查询结果并从数据库主机调用 `scp`。

### 运维边界

发布的 0.0.2 SQL 存在实质缺陷。`advcopy.import_to_s3` 把 URL 解析模块赋给了错误的 Python 变量，随后引用未定义的 `urlparse`，因此文档中的 S3 导出路径按原样会失败。`advcopy.export_to_ip` 声明了 `RETURNS int`，但没有返回语句。依赖这两个函数前必须修补并测试上游代码。

两个 S3 函数都使用 `/tmp` 下可预测的文件；`advcopy.import_from_s3` 固定使用 `/tmp/s3_data.csv`。并发调用可能互相覆盖，失败调用可能遗留敏感数据，而且函数不会清理临时文件。表名与查询文本会被插入 SQL，`scp` 目标也由调用参数拼接。不要把这些函数授权给不受信角色。

该扩展没有定义 CSV 选项、事务安全的对象写入、重试、分片上传、凭据管理或逐行错误处理。应将其视为已停止维护的概念验证代码；生产数据搬运应优先采用持续维护的客户端工具。
