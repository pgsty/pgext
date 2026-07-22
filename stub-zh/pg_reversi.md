## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/README.md)
- [1.0 版本游戏实现](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/pg_reversi--1.0.sql)
- [示例对局](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/sample/demo.sql)

`pg_reversi` 使用数据库级全局表实现一局黑白棋。创建扩展时会删除并重建 `turn`、拼错名称的棋盘表 `boad`、`history` 和 `num_seq`，随后由辅助函数修改这份共享状态。

```sql
CREATE EXTENSION pg_reversi;
SELECT black(3, 2);
SELECT white(2, 2);
SELECT get_turn_boad_status();
TABLE boad;
```

`black(x,y)` 和 `white(x,y)` 落子，方向辅助函数判定翻转，pass 函数推进回合，`history` 记录落子。该实现是教学用 PL/pgSQL 程序，并非可复用的游戏模型：它没有对局标识、玩家身份、行级隔离、授权层或并发协议。

安装会破坏目标模式中同名对象，而且所有会话争用同一棋盘。所复核代码已废弃，包含手写方向逻辑和未完成注释。只能在可丢弃模式使用，应撤销不可信角色的写权限；若应用数据库中已有 `turn`、`boad`、`history` 或 `num_seq`，不要安装该扩展。
