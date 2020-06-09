// 1
===
mov r0 6
mov r1 3
and r2 r0 r1
---
100 111 00000 0 00000000 00000110 1111
100 111 00001 0 00000000 00000011 1111
000 111 00010 00000 00000 00 00001 0100
---
r0 = 00000006
r1 = 00000003
r2 = 00000002


// 2
===
mov r0 6
and r1 r0 3 << 2
---
100 111 00000 0 00000000 00000110 1111
001 111 00001 00000 0001 00000011 0100
---
r0 = 00000006
r1 = 00000004


// 3
===
mov r0 0x8040
and r0 0x8000
---
100 111 00000 0 10000000 01000000 1111
100 111 00000 0 10000000 00000000 0100
---
r0 = 00008000


// 4
===
mov r0 0x8040 << 16
and r0 0x0040 << 16
---
100 111 00000 1 10000000 01000000 1111
100 111 00000 1 00000000 01000000 0100
---
r0 = 00400000