; 0 is check first a
0 a a r 1
0 / / r halt

; 1 is check second a
1 a x r 2
1 / / r 4

; 2 is write one a for them
2 a a r 2
2 / / r 2
2 2 2 r 2
2 = = r 2
2 _ a l 3

; 3 is go back until x
3 a a l 3
3 / / l 3
3 2 2 l 3
3 = = l 3
3 x a r 0

; 4 is handle the odd number
4 / / r 4
4 2 2 r 4
4 = = r 4
4 a a r 4
4 _ , r 5

; 5 is write 5
5 _ 5 r halt

; 6/halt is end
