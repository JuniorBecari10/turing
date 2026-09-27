; Unary increment: add one more '1' to a run of 1s on the tape
; Tape alphabet: '1' and '_' (blank)
; Start state: q0, head at leftmost 1

; Walk right across all the 1s until you hit the first blank
q0 1 * R q0

; Found the blank at the end — write a new 1, move left, and halt
q0 _ 1 L qf
