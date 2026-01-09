# Indiana Jones and the Great Circle Puzzle Solver
This is a puzzle solver for the ancient relic puzzles in the game Indiana Jones and the Great Circle!

![Indiana Jones and the Great Circle backdrop](assets/IJGCBackdrop.png)

## Ancient Relic Puzzle Overview

Throughout the game you collect ancient relics: 

![Ancient Relic](assets/AncientRelic.png)

These relics can be used in 3 locations throughout the game on logic puzzles that use the ancient relics to fill in a grid. Turning relics in different ways on the grid fill in areas differently. Here is an example of the first puzzle that can be found in Gizeh, Egypt:

![Gizeh Puzzle](assets/GizehPuzzle.png)

This puzzle can be represented as a grid with 8 rows and 4 columns. Here is such a grid where 0 represents a slot that is not activated, 1 represents a slot that is activated, and . represents the lack of a slot:

```
00..
00..
0000
000.
.0..
000.
.000
.0..
```

In the image all the slots are activated but for the purpose of the example, all of the slots are deactivated (marked as 0).

## Mechanics

To fill in the puzzle, you have to insert one of the relics into a slot. From there you have two options:

1. Turn the relic left to fill all 8 slots surrounding it
2. Turn the relic to the right to fill in all slots in the same row and column until it hits a wall (spot without a slot)

## Goal
At the end of the game, you are at a ziggurat (a temple) in Iraq. Inside the ziggurat there is a huge ancient relic puzzle. The goal of this project is to be able to solve that puzzle using as little relics as possible.