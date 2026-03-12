

## Usage

> [pagevis: ASCII visualization of PostgreSQL database pages](https://github.com/hollobon/pagevis)

pagevis provides an ASCII-graphical representation of the contents of a PostgreSQL database page. It requires the `pageinspect` extension.

### Function

```sql
-- show_page(relation, block_number, line_width)
SELECT show_page('my_table', 0, 64);
```

### Output Characters

| Character | Meaning |
|-----------|---------|
| `P` | Page header |
| `U` | Unused line pointer |
| `N` | Normal line pointer |
| `R` | Redirect line pointer |
| `D` | Dead line pointer |
| (space) | Free space (the "hole") |
| `[n]` | Tuple number (overlaid on tuple header) |
| `H` | Tuple header |
| `#` | Tuple data |
| `-` | Tuple invisible to current transaction (dead) |
| `.` | Inter-tuple padding |

### Example

```sql
SELECT show_page('pvtest', 0, 64);
                             show_page
------------------------------------------------------------------
 PPPPPPPPPPPPPPPPPPPPPPPP001N002N003N004N005N006N007N008N009N010N
 011N012N013N014N015N016N017N018N019N020N...

 [226]HHHHHHHHHHHHHHHHHHH####....[225]HHHHHHHHHHHHHHHHHHH####....
 [224]HHHHHHHHHHHHHHHHHHH####....[223]HHHHHHHHHHHHHHHHHHH####....
```

The page header (`P`) and line pointers (`N`) appear at the start, free space in the middle, and tuples grow from the end of the page backward. The page is full when no more space remains between line pointers and tuples.

Requires superuser access (uses `pageinspect`'s `get_raw_page`).
