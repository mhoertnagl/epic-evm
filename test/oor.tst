// 1
===
mov r0 6            | 100 111 00000 0  00000000 00000110 1111 | r0 = 00000006
mov r1 3            | 100 111 00001 0  00000000 00000011 1111 | r1 = 00000003
oor r2 r0 r1        | 000 111 00010 00000 00000 00 00001 0101 | r2 = 00000007


// 2
===
mov r0 6            | 100 111 00000 0  00000000 00000110 1111 | r0 = 00000006
oor r1 r0 3 << 2    | 001 111 00001 00000  0001 00000011 0101 | r1 = 0000000e


// 3
===
oor r2 0x8040       | 100 111 00010 0  10000000 01000000 0101 | r2 = 00008040


// 4
===
oor r2 0x8040 << 16 | 100 111 00010 1  10000000 01000000 0101 | r2 = 80400000
