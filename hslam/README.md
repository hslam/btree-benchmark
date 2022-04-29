**hslam-btree**

degree 100
```
Insert: total 10000000, degree 100,     time 1.413691466s,      141 ns/op,      7073679 items/s
Insert: total 10000000, degree 100,     time 1.258827655s,      125 ns/op,      7943899 items/s
Insert: total 10000000, degree 100,     time 1.314273661s,      131 ns/op,      7608765 items/s
Search: total 10000000, degree 100,     time 1.347351076s,      134 ns/op,      7421970 items/s
Search: total 10000000, degree 100,     time 1.362584117s,      136 ns/op,      7338996 items/s
Search: total 10000000, degree 100,     time 1.332764378s,      133 ns/op,      7503201 items/s
Delete: total 10000000, degree 100,     time 1.343990557s,      134 ns/op,      7440528 items/s
Delete: total 10000000, degree 100,     time 1.352372923s,      135 ns/op,      7394410 items/s
Delete: total 10000000, degree 100,     time 1.36409608s,       136 ns/op,      7330861 items/s
Iterate:        total 10000000, degree 100,     time 45.596724ms,       4 ns/op,        219314001 items/s
Iterate:        total 10000000, degree 100,     time 46.073481ms,       4 ns/op,        217044594 items/s
Iterate:        total 10000000, degree 100,     time 43.438133ms,       4 ns/op,        230212472 items/s
```


degree 1000
```
Insert: total 10000000, degree 1000,    time 1.336821036s,      133 ns/op,      7480432 items/s
Insert: total 10000000, degree 1000,    time 1.162807529s,      116 ns/op,      8599875 items/s
Insert: total 10000000, degree 1000,    time 1.132060367s,      113 ns/op,      8833451 items/s
Search: total 10000000, degree 1000,    time 1.397190512s,      139 ns/op,      7157220 items/s
Search: total 10000000, degree 1000,    time 1.330215309s,      133 ns/op,      7517579 items/s
Search: total 10000000, degree 1000,    time 1.37028977s,       137 ns/op,      7297726 items/s
Delete: total 10000000, degree 1000,    time 3.394181961s,      339 ns/op,      2946218 items/s
Delete: total 10000000, degree 1000,    time 3.409340616s,      340 ns/op,      2933118 items/s
Delete: total 10000000, degree 1000,    time 3.442806201s,      344 ns/op,      2904607 items/s
Iterate:        total 10000000, degree 1000,    time 23.167604ms,       2 ns/op,        431637212 items/s
Iterate:        total 10000000, degree 1000,    time 23.793789ms,       2 ns/op,        420277745 items/s
Iterate:        total 10000000, degree 1000,    time 23.574199ms,       2 ns/op,        424192567 items/s
```

100 rand
```
Insert: total 10000000, degree 100,     time 13.86646008s,      1386 ns/op,     721164 items/s
Insert: total 10000000, degree 100,     time 13.536325294s,     1353 ns/op,     738752 items/s
Insert: total 10000000, degree 100,     time 13.577167609s,     1357 ns/op,     736530 items/s
Search: total 10000000, degree 100,     time 16.232815321s,     1623 ns/op,     616036 items/s
Search: total 10000000, degree 100,     time 16.575422559s,     1657 ns/op,     603302 items/s
Search: total 10000000, degree 100,     time 16.158886707s,     1615 ns/op,     618854 items/s
Delete: total 10000000, degree 100,     time 16.254769663s,     1625 ns/op,     615204 items/s
Delete: total 10000000, degree 100,     time 16.275803306s,     1627 ns/op,     614408 items/s
Delete: total 10000000, degree 100,     time 16.260554689s,     1626 ns/op,     614985 items/s
Iterate:        total 10000000, degree 100,     time 54.610651ms,       5 ns/op,        183114462 items/s
Iterate:        total 10000000, degree 100,     time 53.919714ms,       5 ns/op,        185460924 items/s
Iterate:        total 10000000, degree 100,     time 55.54282ms,        5 ns/op,        180041272 items/s
```

1000 rand
```
Insert: total 10000000, degree 1000,    time 19.883643877s,     1988 ns/op,     502925 items/s
Insert: total 10000000, degree 1000,    time 19.638743319s,     1963 ns/op,     509197 items/s
Insert: total 10000000, degree 1000,    time 19.677427297s,     1967 ns/op,     508196 items/s
Search: total 10000000, degree 1000,    time 15.987205202s,     1598 ns/op,     625500 items/s
Search: total 10000000, degree 1000,    time 15.999300017s,     1599 ns/op,     625027 items/s
Search: total 10000000, degree 1000,    time 16.039837314s,     1603 ns/op,     623447 items/s
Delete: total 10000000, degree 1000,    time 24.486833594s,     2448 ns/op,     408382 items/s
Delete: total 10000000, degree 1000,    time 24.478826897s,     2447 ns/op,     408516 items/s
Delete: total 10000000, degree 1000,    time 24.439078376s,     2443 ns/op,     409180 items/s
Iterate:        total 10000000, degree 1000,    time 23.395185ms,       2 ns/op,        427438381 items/s
Iterate:        total 10000000, degree 1000,    time 23.482586ms,       2 ns/op,        425847476 items/s
Iterate:        total 10000000, degree 1000,    time 25.449614ms,       2 ns/op,        392933268 items/s
```