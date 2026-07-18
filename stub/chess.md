## Usage

Sources:

- [Extension control file](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess.control)
- [SQL installation script](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess--1.0.sql)
- [C implementation](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess.c)

`chess` version `1.0` defines `san` and `fen` types for chess game/move text and board positions. It also supplies B-tree comparison operators for `san` plus helpers that extract opening moves, reconstruct a board after a move number, or test whether a game contains an opening or board.

### Example

```sql
CREATE EXTENSION chess;
SELECT getFirstMoves('1.e4 e5 2.Nf3 Nc6'::san, 2);
SELECT getBoard('1.e4 e5 2.Nf3 Nc6'::san, 4);
```

Treat this as historical demonstration code, not a production parser. The repository has almost no user documentation or compatibility statement, and the C storage/input/output implementation embeds pointers in fixed-length values and performs unsafe fixed-size copies. Malformed or simply long input can corrupt memory or crash a backend. Do not expose it to untrusted input; isolate any evaluation and prefer a maintained chess parser with a safe, self-contained storage format.
