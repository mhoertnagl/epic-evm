// 1
===
mov r0 1
mov r1 r0
---
100 111 00000 0 00000000 00000001 1111
000 111 00001 00000 00000 00 00000 1111
---
r0 = 00000001
r1 = 00000001


// 2
===
mov r0 1
mov r1 r0 << 1
---
100 111 00000 0 00000000 00000001 1111
000 111 00001 00000 00001 00 00000 1111
---
r0 = 00000001
r1 = 00000002


// 3
===
mov r0 0x4000
mov r1 r0 <<> 18
---
100 111 00000 0 01000000 00000000 1111
000 111 00001 00000 10010 01 00000 1111
---
r0 = 00004000
r1 = 00000001


// 4
===
mov r0 0x4000
mov r1 r0 >> 4
---
100 111 00000 0 01000000 00000000 1111
000 111 00001 00000 00100 10 00000 1111
---
r0 = 00004000
r1 = 00000400


// 5
===
mov r0 0xffff << 16
oor r0 0x8000
mov r1 r0 >> 4
---
100 111 00000 1 11111111 11111111 1111
100 111 00000 0 10000000 00000000 0101
000 111 00001 00000 00100 10 00000 1111
---
r0 = ffff8000
r1 = 0ffff800


// 6
===
mov r0 0xffff << 16
oor r0 0x8000
mov r1 r0 >>> 4
---
100 111 00000 1 11111111 11111111 1111
100 111 00000 0 10000000 00000000 0101
000 111 00001 00000 00100 11 00000 1111
---
r0 = ffff8000
r1 = fffff800


// 7
===
mov r0 1 <<> 0
---
001 111 00000 00000 0000 00000001 1111
---
r0 = 00000001


// 8
===
mov r0 0x80 <<> 0
---
001 111 00000 00000 0000 10000000 1111
---
r0 = 00000080


// 9
===
mov r0 1 <<> 0
---
001 111 00000 00000 0000 00000001 1111
---
r0 = 00000001


// 10
===
mov r0 0x8000
---
100 111 00000 0 10000000 00000000 1111
---
r0 = 00008000


// 11
===
mov r0 1 << 16
---
100 111 00000 1 00000000 00000001 1111
---
r0 = 00010000


// 12
===
mov r0 0x8000 << 16
oor r0 r0 1
---
100 111 00000 1 10000000 00000000 1111
001 111 00000 00000 0000 00000001 0101
---
r0 = 80000001
