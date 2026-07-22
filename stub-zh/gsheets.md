## 用法

来源：

- [官方 README](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/README.md)
- [扩展控制文件](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/gsheets.control)

`gsheets` 允许 SQL 会话读写 Google Sheets。它适合小型集成、导入和导出流程，前提是数据库服务器能够访问 Google API，并且会话可以提供 OAuth 访问令牌。

### 核心流程

创建扩展，完成上游定义的认证流程，并为当前会话设置取得的令牌：

```sql
CREATE EXTENSION gsheets;

SELECT gsheets_auth();
SET gsheets.access_token = 'your_access_token';
```

`read_sheet` 是多态函数：调用方需要声明返回列。第一个参数可以是电子表格 ID 或 URL。

```sql
SELECT *
FROM read_sheet(
    '<spreadsheet_id_or_url>',
    sheet_name => 'Sheet1',
    header => true
) AS sheet(name text, age integer);
```

使用 `write_sheet` 写入行。不提供 `spreadsheet_id` 时，扩展会请求创建新电子表格；否则通过 JSON 选项指定已有电子表格和工作表。

```sql
SELECT write_sheet(
    person.*,
    '{"spreadsheet_id":"<spreadsheet_id>","sheet_name":"Sheet2","header":["name","age"]}'::jsonb
)
FROM person;
```

### 用户接口

- `gsheets_auth()`：启动上游文档描述的 OAuth 认证流程。
- `gsheets.access_token`：保存 Google 访问令牌的会话设置。
- `read_sheet(spreadsheet_id/url text, sheet_name text DEFAULT 'Sheet1', header boolean DEFAULT true)`：读取行；查询中需要声明预期的记录列。
- `write_sheet(data anyelement, options jsonb DEFAULT '{}')`：写入标量或复合行。选项包括 `spreadsheet_id`、`sheet_name` 和 `header`。

### 运维说明

该扩展链接 libcurl，并由 PostgreSQL 服务器进程发起外部网络请求。认证需要交互式 URL/令牌流程，因此无头服务器和令牌续期都需要专门规划。不要把 `gsheets.access_token` 写入日志或持久化的角色/数据库设置，并且只向允许访问目标工作表的角色授予扩展使用权。上游没有记录预加载要求。将 `write_sheet` 用于自动化作业前，应测试 API 错误、限流、部分写入和重试行为。
