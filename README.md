# sqlite mmap leak repro

This repo contains a test case that will cause [cznic/sqlite](https://gitlab.com/cznic/sqlite) to leak mmap allocations.

This testcase downloads a database of cities from around the world, and
attempts to look up the latitude/longitude of a selection.


The leak seems to occur when the program runs queries that return no
rows. There are some entries in the test data labelled `does not exist`
which will cause `QueryRow()` to return `sql.ErrNoRows` to be returned.

When these entries are commented out, `make pmap` will remain at a
consistent number of entries, roughly ~18 on my machine. If they're
uncommented, then the number of mmap allocations gradually grows.

## running the repro script

You can run the script using:

```
make debug
```

which will compile the testcase and start debugging it using
[delve](https://github.com/go-delve/delve). Delve pause execution after
compilation, and give you the opportunity to set breakpoints:

```
(dlv) break modernc.org/memory.(*Allocator).mmap
```

You can also set conditions on the breakpoint, such as this one, to only
break after a certain number of mmaps have been allocated:

```
(dlv) cond 1 len(a.regs) > 14
```

You can then continue execution to run the program, and `dlv` will pause
execution at the `mmap` call when the condition is met:

```
(dlv) continue
```

Once execution halts at the breakpoint you can print the stack, and move
up/down it:

```
(dlv) stack
(dlv) up 1
(dlv) up 10
(dlv) down 10
```


It's possible to track the number of allocated mmaps from outside the
process using `make pmap`.

## forcing the leak

