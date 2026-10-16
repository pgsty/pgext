## 用法

来源：

- [README.md](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/README.md)
- [client/extension/pokemon--0.1.0.sql](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/pokemon--0.1.0.sql)
- [client/extension/Cargo.toml](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/Cargo.toml)
- [LICENSE](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/LICENSE)
- [client/extension/pokemon.control](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/pokemon.control)

`pokemon` 0.1.0 通过外部表暴露模拟器状态，是 PostgresPlaysPokemon 的 PostgreSQL 客户端。模拟器本身不在 PostgreSQL 内运行。

### 核心用法

```sql
CREATE EXTENSION pokemon;
IMPORT FOREIGN SCHEMA red FROM SERVER red INTO public;
SELECT * FROM party;
```

### 运行边界

官方环境测试 PostgreSQL 10，需要单独运行基于 PyBoy 的服务器、准备 ROM 数据，并能通过 Unix socket 使用 gRPC。创建扩展会注册 `pokemon_fdw()`、封装器 `pokemon` 和服务器 `red`。导入模式 `red` 后可使用 `party`、`inventory`、`story`：party 与 story 支持 SELECT、UPDATE，inventory 还支持 INSERT。写入会修改模拟器状态。默认仅超级用户可安装；这个历史集成没有证明对当前 PostgreSQL 大版本或可配置远程连接的支持。
