; 0 is check left-side a
; if it doesn't exist anymore, end the script
0 a . r 1
0 x x r halt ; or 6

; 1 is go to x
1 a a r 1
1 x x r 2

; 2 is go to next a in right-side
2 . . r 2
2 a , r 3
2 = = l 5 ; consumed all a in right-side

; 3 is put one a for them after =
3 a a r 3
3 = = r 3
3 _ a l 4

; 4 is go back to next a in right-side
4 = = l 4
4 a a l 4
4 , , r 2

; 5 is fill all , with a again and go back to next a in left-side
5 = = l 5
5 , a l 5
5 x x l 5
5 a a l 5
5 . a r 0
