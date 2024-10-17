# gpt-vetting coding assignment
## Table of Contents
- [Story](#story)
- [Solution](#solution)
- [Conclusion](#conclusion)

## Story
__macro-1__ uses _gpt-vetting_ in their interview selection process. It has 2 parts: set of questions on different relevant subjects and coding challenge (of course). You have to turn on your camera and answer within an allocated time (I think it's a minute or two) to the questions asked by a cheerful enthusiastic voice. Then you are given a blank screen with "hello world" in language of choice and 25 minutes to solve the coding puzzle. My puzzle was to traverse a maze and find or not find all the keys and then report whether all of them can be found. I can't do that in 25 minutes. Afterwords it took me about 45 minutes to come up with solution and about an hour to code and more or less polish it. So, needless to say - I failed their coding test.
Subjectively, this is a dehumanizing procedure. They call it "fair".
I easily can see a future development of this approach when there will be positions they would not be able to fill because gpt-vetting will reject qualified candidates based on a coding challenge that has nothing to do with daily tasks.

## Solution

Solution was implemented in python3 [here](maze-traversal.py). Maze walls are represented as _pound_(__#__), empty cells as _dot_(__.__), cells with keys as _asterisk_(__*__) and path as _semicolons_(__;__).
Solution was developed to search for keys and also to detect additional exit from the labyrinth in addition to the designated one. __mazes__ variable holds maze string, maze dimentions - 15 x 10, start cell and exit cell.

```python
mazes = \
[
(
'#.#############'+\
'#......##...###'+\
'###.#####.#.###'+\
'###...##.....*#'+\
'#....###.##.###'+\
'#*##.....##.###'+\
'###########.###'+\
'########..*.###'+\
'#........######'+\
'####.##.#######',
15, 10,
(0, 0), (9, 4)
)
]
```

### Output

```text
[0]-> python maze-traversal.py
KEY COUNT: 2
MAZE:
00 #.#############
01 #......##...###
02 ###.#####.#.###
03 ###...##.....*#
04 #....###.##.###
05 #*##.....##.###
06 ###########.###
07 ########..*.###
08 #........######
09 ####.##.#######

SOLUTION:
00 #;#############
01 #;;;...##;;;###
02 ###;#####;#;###
03 ###;;.##;;.;.*#
04 #...;###;##;###
05 #*##;;;;;##;###
06 ###########;###
07 ########;;;;###
08 #......;;######
09 ####.##;#######

KEY COUNT: 1
MAZE:
00 #;#############
01 #;;;...##;;;###
02 ###;#####;#;###
03 ###;;.##;;.;.*#
04 #...;###;##;###
05 #*##;;;;;##;###
06 ###########;###
07 ########;;;;###
08 #......;;######
09 ####.##;#######

SOLUTION:
00 #;#############
01 #;;;...##;;;###
02 ###;#####;#;###
03 ###;;.##;;.;.*#
04 #...;###;##;###
05 #*##;;;;;##;###
06 ###########;###
07 ########;;;;###
08 #...;;;;;######
09 ####;##;#######
```

## Conclusion
My solution is not perfect and one can find ways to improve it. Hope it helps everyone out there.