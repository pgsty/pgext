## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/README.md)
- [Version 1.0 game implementation](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/pg_reversi--1.0.sql)
- [Example game](https://github.com/nuko-yokohama/pg_reversi/blob/2afdbb8941d1c59173c711d906638c6e57a76aa6/sample/demo.sql)

`pg_reversi` implements a single Reversi game in database-global tables. Creating the extension drops and recreates `turn`, the misspelled board table `boad`, `history`, and `num_seq`; helper functions then mutate that shared state.

```sql
CREATE EXTENSION pg_reversi;
SELECT black(3, 2);
SELECT white(2, 2);
SELECT get_turn_boad_status();
TABLE boad;
```

`black(x,y)` and `white(x,y)` place stones, directional helpers determine flips, pass functions advance turns, and `history` records moves. The implementation is an educational PL/pgSQL program, not a reusable game model: there is no game identifier, player identity, row-level isolation, authorization layer, or concurrency protocol.

Installation is destructive to same-named objects in the target schema, and every session competes over one board. The reviewed code is abandoned and contains hand-written directional logic with unfinished comments. Use only in a disposable schema, revoke writes from untrusted roles, and do not install it in an application database with objects named `turn`, `boad`, `history`, or `num_seq`.
