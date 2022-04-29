**google-btree**

degree 100
```
Insert: total 10000000, degree 100,     time 1.953285244s,      195 ns/op,      5119579 items/s
Insert: total 10000000, degree 100,     time 1.795027763s,      179 ns/op,      5570944 items/s
Insert: total 10000000, degree 100,     time 1.830557033s,      183 ns/op,      5462818 items/s
Search: total 10000000, degree 100,     time 1.721107286s,      172 ns/op,      5810213 items/s
Search: total 10000000, degree 100,     time 1.670545249s,      167 ns/op,      5986069 items/s
Search: total 10000000, degree 100,     time 1.722480451s,      172 ns/op,      5805581 items/s
Delete: total 10000000, degree 100,     time 1.922569885s,      192 ns/op,      5201371 items/s
Delete: total 10000000, degree 100,     time 1.87509982s,       187 ns/op,      5333049 items/s
Delete: total 10000000, degree 100,     time 1.921985125s,      192 ns/op,      5202953 items/s
Iterate:        total 10000000, degree 100,     time 43.217064ms,       4 ns/op,        231390082 items/s
Iterate:        total 10000000, degree 100,     time 44.120885ms,       4 ns/op,        226650032 items/s
Iterate:        total 10000000, degree 100,     time 44.996003ms,       4 ns/op,        222241962 items/s
```

degree 1000
```
Insert: total 10000000, degree 1000,    time 1.952061737s,      195 ns/op,      5122788 items/s
Insert: total 10000000, degree 1000,    time 1.689028979s,      168 ns/op,      5920561 items/s
Insert: total 10000000, degree 1000,    time 1.717608153s,      171 ns/op,      5822049 items/s
Search: total 10000000, degree 1000,    time 1.683942292s,      168 ns/op,      5938445 items/s
Search: total 10000000, degree 1000,    time 1.647018045s,      164 ns/op,      6071578 items/s
Search: total 10000000, degree 1000,    time 1.663677455s,      166 ns/op,      6010780 items/s
Delete: total 10000000, degree 1000,    time 3.797230736s,      379 ns/op,      2633498 items/s
Delete: total 10000000, degree 1000,    time 3.768688919s,      376 ns/op,      2653442 items/s
Delete: total 10000000, degree 1000,    time 3.791164428s,      379 ns/op,      2637712 items/s
Iterate:        total 10000000, degree 1000,    time 35.79993ms,        3 ns/op,        279330155 items/s
Iterate:        total 10000000, degree 1000,    time 32.55252ms,        3 ns/op,        307195879 items/s
Iterate:        total 10000000, degree 1000,    time 31.617216ms,       3 ns/op,        316283381 items/s
```

100 rand
```
Insert:	total 10000000,	degree 100,	time 14.73154281s,	1473 ns/op,	678815 items/s
Insert:	total 10000000,	degree 100,	time 14.669807688s,	1466 ns/op,	681672 items/s
Insert:	total 10000000,	degree 100,	time 14.464650177s,	1446 ns/op,	691340 items/s
Search:	total 10000000,	degree 100,	time 17.863390795s,	1786 ns/op,	559804 items/s
Search:	total 10000000,	degree 100,	time 18.651310126s,	1865 ns/op,	536155 items/s
Search:	total 10000000,	degree 100,	time 17.953943351s,	1795 ns/op,	556980 items/s
Delete:	total 10000000,	degree 100,	time 17.603552631s,	1760 ns/op,	568067 items/s
Delete:	total 10000000,	degree 100,	time 17.807744062s,	1780 ns/op,	561553 items/s
Delete:	total 10000000,	degree 100,	time 17.938793589s,	1793 ns/op,	557451 items/s
Iterate:	total 10000000,	degree 100,	time 60.852936ms,	6 ns/op,	164330608 items/s
Iterate:	total 10000000,	degree 100,	time 62.246169ms,	6 ns/op,	160652457 items/s
Iterate:	total 10000000,	degree 100,	time 57.406867ms,	5 ns/op,	174195188 items/s
```

1000 rand
```
Insert: total 10000000, degree 1000,    time 20.027466104s,     2002 ns/op,     499314 items/s
Insert: total 10000000, degree 1000,    time 22.151479474s,     2215 ns/op,     451437 items/s
Insert: total 10000000, degree 1000,    time 19.880276837s,     1988 ns/op,     503011 items/s
Search: total 10000000, degree 1000,    time 16.455086434s,     1645 ns/op,     607714 items/s
Search: total 10000000, degree 1000,    time 16.375941199s,     1637 ns/op,     610651 items/s
Search: total 10000000, degree 1000,    time 16.387806127s,     1638 ns/op,     610209 items/s
Delete: total 10000000, degree 1000,    time 23.207065198s,     2320 ns/op,     430903 items/s
Delete: total 10000000, degree 1000,    time 23.155711989s,     2315 ns/op,     431858 items/s
Delete: total 10000000, degree 1000,    time 23.109427881s,     2310 ns/op,     432723 items/s
Iterate:        total 10000000, degree 1000,    time 35.790994ms,       3 ns/op,        279399895 items/s
Iterate:        total 10000000, degree 1000,    time 34.302738ms,       3 ns/op,        291521918 items/s
Iterate:        total 10000000, degree 1000,    time 30.869165ms,       3 ns/op,        323947861 items/s
```