## 用法

来源：

- [官方 README](https://github.com/sabino/openapi_fdw/blob/v0.4.1/README.md)
- [扩展控制文件](https://github.com/sabino/openapi_fdw/blob/v0.4.1/openapi_fdw.control)
- [pgrx 清单](https://github.com/sabino/openapi_fdw/blob/v0.4.1/Cargo.toml)

`openapi_fdw` 把 OpenAPI 3.0/3.1 JSON 服务映射为实时 PostgreSQL 外部表。它从公开契约导入类型化列，也可以在 `attrs jsonb` 中保留完整源对象。

### 启用

0.4.1 版本为 PostgreSQL 14–18 发布构件。安装与大版本准确匹配的原生库后，创建仅限超级用户、可迁移的扩展：

```sql
CREATE EXTENSION openapi_fdw;
```

无需预加载或重启。可选 control-plane 服务用于发现并应用定义，配置完成后不再是运行依赖。

### 导入 API

创建指向可信 OpenAPI 文档的外部服务器，再把选定操作导入 schema。

```sql
CREATE SERVER pokeapi
FOREIGN DATA WRAPPER openapi_fdw
OPTIONS (
  spec_url 'https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml'
);

CREATE SCHEMA poke;

IMPORT FOREIGN SCHEMA api
  LIMIT TO (pokemon_list)
  FROM SERVER pokeapi
  INTO poke
  OPTIONS (methods 'GET', include_attrs 'true');

SELECT name, attrs ->> 'url'
FROM poke.pokemon_list
LIMIT 5;
```

每次扫描都会发出有界的实时 HTTP 请求；需要本地快照时应显式物化结果。带路径参数的端点只有在等值条件绑定全部路径变量后才会发起请求。

### 认证与写入

优先使用 `bearer_token_env`、`api_key_env` 与 `headers_env` 等环境变量选项，让 PostgreSQL 服务进程解析凭据，避免把明文 secret 写入 catalog option。只有显式配置 insert/update/delete 端点后，外部表才可写。

远端 HTTP 写入不参与 PostgreSQL 事务：后续回滚无法撤销已接受的请求，多行语句也可能部分成功。0.4.1 每行发送一次请求，不会自动重试 POST/PATCH，也不支持 `RETURNING`；幂等与补偿逻辑必须在远端 API 边界设计。
