## 用法

来源：

- [Pinned official README](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/README.md)
- [Pinned extension SQL and game interface](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/sql/pg2podg.sql)

`pg2podg` 是面向双人、信息公开、确定性游戏的 SQL/PLpgSQL 搜索引擎。它提供可复用的游戏树和走法评估流程，但不实现具体游戏；必须由另一个扩展或模式先提供所需的 `game` 与 `move` 接口。

### 所需游戏接口

创建 `pg2podg` 前，数据库必须定义：

- 类型 `game` 和 `move`。
- 运算符 `## game`、`%% game`、`game ^ move`，分别用于显示、稳定文本标识和执行走法。
- 函数 `new_game()`、`score(game)`、`gain(game, game)`、`valid_moves(game)`、`parse_move(text, game)`。
- 终端 UI 使用的单参数显示运算符 `# game`。

上游示例用 `pgchess` 提供这套契约。应先创建具体游戏扩展：

```sql
CREATE EXTENSION pgchess;
CREATE EXTENSION pg2podg;

SELECT ui_loop(
  iter := 10,
  time_target := interval '1 second',
  depth_target := 2,
  restart := true,
  regress := true
);
```

### 核心对象

- 持久搜索状态：`games` 存储局面和父走法；`gains` 缓存评估。
- 工作状态：非日志表 `choices`、`status`，以及选择排序/显示视图。
- 搜索：`research_games`、`compute_gains`、`evaluate_choices`、`apply_best_choice`。
- 用户界面：`ui_cpu_moves`、`ui_human_moves`、`ui_loop`、`status_display_vt100`、`display_game_end`。

`research_games` 将局面至少扩展到指定深度，并在第一层之后应用可选的时间目标。`evaluate_choices` 对已探索的锥体评分；`apply_best_choice` 推进当前局面。

### 数据与升级注意事项

安装时，扩展会向 `games` 插入初始局面。`games` 和 `gains` 是扩展拥有的持久表；删除扩展会删除其中内容。上游 0.1.3 README 称升级需要删除、重新安装并重建扩展，而且不会保留这些表，因此应先导出有价值的状态。

该代码面向 PostgreSQL 9.1 时代的接口，游戏契约也只隐含在 SQL 源码中。安装前应验证候选游戏实现及其运算符/函数，并使用有界的时间和深度值，因为游戏树可能快速增长。无需预加载或重启。
