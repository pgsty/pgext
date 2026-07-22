## Usage

Sources:

- [Extension SQL](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/sql/chess_index.sql)
- [C implementation](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/src/chess_index.c)
- [Regression examples](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/test/sql/setup.sql)
- [Build file](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/Makefile)
- [Extension control file](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/chess_index.control)

`chess_index` defines compact PostgreSQL types for chess positions, pieces, and squares, with comparison and hash/B-tree operator classes intended to make positions indexable. The project is an unfinished pre-release implementation; use it only for research with expendable data, not for correctness-sensitive chess storage or indexes.

### Core Workflow

The `board` input is a shortened four-field FEN value: piece placement, active side, castling rights, and en-passant target. Halfmove and fullmove counters are not part of the stored representation.

```sql
CREATE EXTENSION chess_index;

CREATE TABLE positions (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    position board NOT NULL
);

INSERT INTO positions (position)
VALUES ('rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq -');

SELECT position, pcount(position), side(position), pretty(position)
FROM positions;
```

The SQL declares hash and B-tree operator classes for `board`, `square`, and `pindex`, so indexes can be created syntactically:

```sql
CREATE INDEX positions_board_hash ON positions USING hash (position);
```

Do not rely on such an index until the implementation defects below are fixed and verified.

### Important Objects

- `board`: compact position value with `pcount(board)`, `side(board)`, and `pieces(board, side)` accessors.
- `side`, `piece`, `pindex`, `square`, and `piecesquare`: compact chess component types.
- `cfile`, `rank`, and `diagonal`: square-derived coordinate types.
- `pretty(board, boolean, boolean)`: format a position as text or Unicode pieces.
- `invert(board)`: swap piece case and active side in the textual representation.
- Hash and B-tree operators/operator classes support equality and ordering for selected types.

### Correctness and Packaging Caveats

At the reviewed commit, the Makefile names `sql/chess_index--0.0.1.sql`, but the repository contains only `sql/chess_index.sql`; an unpatched source install can therefore fail before `CREATE EXTENSION` is usable.

More importantly, the C comparator uses the size of a pointer rather than the size of the stored `board` structure when comparing values. Equality and ordering can ignore relevant position bytes, while hashing uses a different representation. This can produce wrong comparisons and unsafe index semantics. The code also keeps the original FEN in a fixed 100-byte array and has a boundary check that permits an input at the array's nominal limit. Treat both parsing and indexing as untrusted experimental code, keep backups, and validate any locally patched build with adversarial FEN values before use.
