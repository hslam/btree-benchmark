**hslam-btree**

degree 100
```
Insert:	total 10000000,	c100,	time 2.189217935s,	218 ns/op,	4567841 items/s
Insert:	total 10000000,	degree 100,	time 2.135339489s,	213 ns/op,	4683096 items/s
Insert:	total 10000000,	degree 100,	time 2.088410656s,	208 ns/op,	4788330 items/s
Search:	total 10000000,	degree 100,	time 2.02869588s,	202 ns/op,	4929275 items/s
Search:	total 10000000,	degree 100,	time 2.011450449s,	201 ns/op,	4971536 items/s
Search:	total 10000000,	degree 100,	time 2.015015238s,	201 ns/op,	4962741 items/s
Delete:	total 10000000,	degree 100,	time 2.197786746s,	219 ns/op,	4550031 items/s
Delete:	total 10000000,	degree 100,	time 2.142481673s,	214 ns/op,	4667484 items/s
Delete:	total 10000000,	degree 100,	time 2.251089457s,	225 ns/op,	4442293 items/s
Iterate:	total 10000000,	time 63.804624ms,	6 ns/op,	156728452 items/s
Iterate:	total 10000000,	time 60.687127ms,	6 ns/op,	164779591 items/s
Iterate:	total 10000000,	time 59.832898ms,	5 ns/op,	167132135 items/s
```


degree 1000
```
Insert:	total 10000000,	degree 1000,	time 2.094750373s,	209 ns/op,	4773838 items/s
Insert:	total 10000000,	degree 1000,	time 2.027024892s,	202 ns/op,	4933338 items/s
Insert:	total 10000000,	degree 1000,	time 1.956000608s,	195 ns/op,	5112472 items/s
Search:	total 10000000,	degree 1000,	time 1.915358123s,	191 ns/op,	5220955 items/s
Search:	total 10000000,	degree 1000,	time 1.912072899s,	191 ns/op,	5229926 items/s
Search:	total 10000000,	degree 1000,	time 1.916595257s,	191 ns/op,	5217585 items/s
Delete:	total 10000000,	degree 1000,	time 4.554162777s,	455 ns/op,	2195793 items/s
Delete:	total 10000000,	degree 1000,	time 4.781057831s,	478 ns/op,	2091587 items/s
Delete:	total 10000000,	degree 1000,	time 4.52973869s,	452 ns/op,	2207632 items/s
Iterate:	total 10000000,	degree 1000,	time 36.810853ms,	3 ns/op,	271659013 items/s
Iterate:	total 10000000,	degree 1000,	time 35.478006ms,	3 ns/op,	281864769 items/s
Iterate:	total 10000000,	degree 1000,	time 36.184541ms,	3 ns/op,	276361112 items/s
```