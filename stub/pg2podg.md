## Usage

Sources:

- [Pinned official README](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/README.md)
- [Pinned extension SQL and game interface](https://github.com/gciolli/pg2podg/blob/039dedf753e8c4aeb05564b4bf9082d80aa8b417/sql/pg2podg.sql)

`pg2podg` is a SQL/PLpgSQL search engine for two-player, open-information, deterministic games. It supplies a reusable game-tree and move-evaluation workflow, but not a game implementation: another extension or schema must first provide the required `game` and `move` interface.

### Required Game Interface

Before creating `pg2podg`, the database must define:

- Types `game` and `move`.
- Operators `## game`, `%% game`, and `game ^ move` for display, stable text identity, and applying a move.
- Functions `new_game()`, `score(game)`, `gain(game, game)`, `valid_moves(game)`, and `parse_move(text, game)`.
- The single-argument display operator `# game`, used by the terminal UI.

The upstream example uses `pgchess` to supply this contract. Create the game-specific extension first:

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

### Core Objects

- Persistent search state: `games` stores positions and parent moves; `gains` caches evaluations.
- Working state: unlogged `choices` and `status` tables, plus choice-ordering/display views.
- Search: `research_games`, `compute_gains`, `evaluate_choices`, and `apply_best_choice`.
- User interface: `ui_cpu_moves`, `ui_human_moves`, `ui_loop`, `status_display_vt100`, and `display_game_end`.

`research_games` expands positions to a minimum depth and, after the first level, an optional time target. `evaluate_choices` scores the explored cone; `apply_best_choice` advances the current position.

### Data and Upgrade Caveats

The extension inserts an initial position into `games` when installed. `games` and `gains` are persistent extension-owned tables; dropping the extension removes their contents. The upstream 0.1.3 README says upgrades require dropping, reinstalling, and recreating the extension and do not preserve those tables, so export any valuable state first.

This code targets PostgreSQL 9.1-era interfaces and its game contract is only implicit in the SQL source. Validate a candidate game implementation and its operators/functions before installation, and use bounded time/depth values because the game tree can grow rapidly. No preload or restart is required.
