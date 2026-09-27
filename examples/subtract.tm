; 0 - detect a
0 a x r 1
0 - - r 3
0 = = l 9

; 1 - put a in front
1 a a r 1
1 - - r 1
1 = = r 1
1 _ a l 2

; 2 - go back until x and go to 0
2 = = l 2
2 a a l 2
2 - - l 2
2 x a r 0

; 3 - detect a to start subtracting
3 a x r 4
3 = = r halt_end ; end - positive or zero
3 - - r halt_end ; end - negative

; 4 remove one a from there
4 a a r 4
4 = = r 4
4 - - r 4
4 _ _ l 5

; 5 - verify if there is a
5 = = l 7 ; negative (a < b; a - b < 0)
5 a _ l 6

; 6 - go back until x and go to 3
6 = = l 6
6 a a l 6
6 - - l 6
6 x a r 3

; 7 - negative. put -
7 x a r 7
7 a a r 7
7 = = r 7
7 _ - l 8

; 8 - go back to - to start the converse operation (b - a)
8 = = l 8
8 a a l 8
8 x a l 8
8 - - r 0 ; go back to 0!

; 9 - go back to the start to start removing
9 = = l 9
9 a a l 9
9 - - l 9
9 _ _ r 3
