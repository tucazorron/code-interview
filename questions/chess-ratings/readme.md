# google

## problem

there is a chess tournament

each player has a unique ranking

the player with a higher ranking always win when playing with someone with a lower ranking

their rankings doesnt update while playing the tournament

you have access to some game results of that tournament

with the result of the games, can you find the final position of any player in the tournament?

if so, return a list of players and their positions

if not, return an empty list

## input

list of pairs of integers

each pair represents the winner and loser from a game, respectively

## output

list of pairs of integers

each pair represents the player and their final position at the tournament, respectively

## tests

input

    [[4, 2], [4, 3], [3, 2], [1, 2], [2, 5]]

output

    [[2, 4], [5, 5]]

input

    [[4, 2], [4, 3], [3, 2], [1, 2], [2, 5], [3, 1]]

output

    [[1, 3], [2, 4], [3, 2], [4, 1], [5, 5]]
