package main

func getInput01() string {
	return `##################################################
#.OOO#.#O..O.O...OO.O......OO....O#O............O#
#O.......#....O.OO..O...O...#.OO..OO.OO....O.O..O#
#OO..#......O...OOOO...O.O...O...#O.O..O...O..OO.#
#...O.......#...........O..O..O........O.O.O..O..#
#OOOO#..OO#.O#.#..O.OOO..OO...O...O..O.OO#.....#.#
#........#O#.......OOO#.....#.....#...O....O#.O.##
#...#..OOO...O.#.......O#OO..O.O...O...OO.O#O....#
#.O..#...O.....#.O#O.O...#.#O##....O#....O...O.O.#
#.O...O..OO.O..O....O.OO.OO.O##..O...#.OO.O#OO..##
#..O..O#.#O.OO..#.....#...O..#O..O.O.#OO.O....#.O#
#...OO.O.O..OO#...O.OOO.....O.....O.O..O..O.O....#
#O....O...O#.O.#.#..OO#.#......#..O..O.....OO....#
#.OOO....O..O#..O....O.....O.O#OO........#..O.#..#
#.....O.O..#.#.....#OOO....#.#..O.O..............#
#.O#.OO....#......O.#.....O.O...OO..O.OO.O..OO..##
#.O.O.....O..#O.O..OOO.O#..O#..O.O...O.O...O.OO..#
#.........O#.##O.O..O.O......O..#....##.O.OO...OO#
#.......OO..O.O.O.OO.O.O......OO#..O.O....O....OO#
#...O.....OO....#.OO#.........O...O...#.#OOOO...O#
#.....O.OO.#.O......O#O..O#.O.....#......O...O..##
#...O.O#....OO.......O....O.#OOO.O..O......OO##OO#
#.........O..O..O.#.#.#.O#OO..........O....OOO...#
#..#O.......#O...O....O.O.#O.O...O..O.O...OO..OO.#
#O.........O.....OOOO...@.OO#OO.....O..O#...#..O.#
#O###OOO.....O..O...O.OO.O....O#...OO..O.O.OOO...#
#.....O..#..OO#.O.OO.....#....OO...........O..OO.#
#O..OO.#.....##..OO...O....................#....##
#O...O.............O.O#O..OO.......OOO...O....O..#
#.OOO.#..OOOOO.....#.O.....O.....#..............O#
#....OO..#...OO.#O..#...#.#O..OO#..#.#.....O.....#
#...O.#...#.#..O....O#......O.OO..O.#...OO.O....O#
#..O...O##O....OO.O..O......#O..O..O#.O...#......#
#O.OO..O........O#.O.#.......O..#O...#.##O....O..#
#.O....O..O.#..O#.O.OO.....O...O.O.O...#.O.....O.#
#O....#.O...OO.....O.OOO......#O..#..O..#.OOO.#..#
#..#.OO#O#.O.#O..O.#.O.O#.OOO.#..#.....O.OOO.....#
##OO..#......#...#..O.O.............OO#O.O...#.O.#
#.OO.O..O.O.#O..#...O..OO.OOOOO..#..#.O...OO#O#O##
#......#...O.#.OO#...O............O.O.OO#..OOO#..#
#..OO.O.....O.....O.O...O#.#.....#.#.OOO.##.OO..##
#....O#O..O..O#.O.....O.....O..........O..OO.OOO.#
#O..O#....#O...O.#....OO.OO....O...O..OO#........#
#O#.O......OO..O...O..#..#O..O...#OO.O...O..O....#
#.#....O...OOO.........#O..O.O#O.O...O.O..#.O.#O.#
#.......O....OO.....OOOO.....O...O....OO....O.O..#
##...O.O.......#.OO................OO.....OO..OO.#
#O..#O#.#OO.....O......O..O.OO.O......#.OO.......#
#O#.#...O.O.#.O.......O.OOO.......O......O.OO.O..#
##################################################

>v>^^<v>vv><>^<<^^<>v^>vvv<^v^><v^v^<>>v^v>v<vv><<^vvv^vv>^<^^^v<>^v^^v>^><<<^>>vv>^<>v^v^^>v^^<^><^^>v<<v>>^<^v>^v>^>><vvv>^>^><vvv<^^<v^>^v>v>>^vv<<v^>>^v<<>vv>>>>^<>^^<<vvv^<><><>>^<vv>v>^^<^v<^v<<<v>><><<>v^v>>v<<><<<>>^><<^<^^v^v>v^vv<v^<v^^v^>^><v>v>^^>>>><>v<<^>v<^^^<v>>>^v^v<v^^^<^>^vv>^>^v<v<v<^<<v<^<>v^v<v>vv^><vv<<>^^^<v<^><<^<<v^vvv><^><<<<^^><>><v^^^^>>^<^^^>^v>>^<v<^<<<<v^<v^>^<v>v^>^>>^><<><^v^v>v^<^^>><^><<<^v<<^>^vv^v^v<v>v>^^><<vv^<>^^v>v^^v^<<<>vv>v>^v>vvv>v^v^v>^>^^vv>>v>>v>^<<v>^v><^^>v^v<>>^vv<^^^<<^v^><>vv>^^^<><v^v<^>>v^^<><vvv<v<vv^>^>><<^^>v<<v<<><<^^>>v<<<<^v>^>v^vv>^v>vvvvvv^><vv><v^<^<^^v<><^^><v<>><^<v^^vv^^^^v<<vv>^^^vvv<^>^^vvv<v>><<^^>><<^>vv^>>^v>vvv><v^v<v<v<^v^>vv>^<<<>^^><><<v>><^><>v^>v^<vv<^v>^<v<<v>>>>^>^^>>^^<vv<><><<<v>v<>^v<<vv^>v>v<<vvv<>>><v^^v>>>>>vv^vv^<v><><>>vv^>v^>>^>>>^>>^>v>><v>>>v<^^>>>^v^vv>>^>^^v<>^>^>^<<>v<^v^>^>^><<>>><>v>^<v>^^vv^v^<<vvvvv^<^>v<>^vv^>>^v^>^v<>><v>v>v^^<vv<><<^v^><^><<>^v><>>^><v^>v^^vv<v>vv>>^>v<vv<><<v<>vv^>v>^
>><>^>^<^^>>><><>^vvvv^><^>^^<><>^<v>^^v^vv<^^^<^^>^>^>vv^<^>>v>><^v^<>>v><^v><^<<<>>>^^^v^^>>><><v>><><v<v<^^<><><<v>^<^^<^<v>vvv^v<^<^v<v><v<><<<><<>^>^<^vv>>v>v<^^v>^^^^<v^>>v^^<>v^<<^v<>^^^^v<^v^^v<>>^v<vv^<v>v^><^v<>vv<v<<^v<<^<<v><>^<^><v<<<<>>^<v>^^vv<<v<<<^<<>>v^>>>>>>>vv>>v^<>v>^>v^vv^>v<><>v>>v<^^v<vvv^><>><><vvv>v<v^^><^v^^v^><v<>>>vvv^><>^><<vv^^>>v><v^v>>^<v<>>^^^^<>v<<v^^<^v<>v<>^v^^>>v<>^>^^<>vv<<v<<^<<vv<>>^^vv><vv<>^<^^>><v>v<^v>^v<v>><^vvv>^v<^<v>>>^^>v<vvv^<^^>>^<>>v<^<v^^v><>^^^v<vv<^>vv>v^^v<v>^>v<><^<^<v>^^<v>>^>^v<>^<<^<<<>^><><<<^^<>v<>^^^v^>v^><>>^<v>vv>>^>^><<<v<^<<<<vv><><^<^<<^^^>>^<^<<^<^^>v<vv<^>>v>><^^>^vvv^^^vvv<><v><vv^v>vv>><>><<^v<v^vv<>v<^v^^>>^><^>><<v>v<>>v<>v<>v^^<<v<^^>^<<<<^v^<>v<v>^<><v^vv^^v^v<><><^>>^>vv>>v<>^<<>><^>^>>^><>v^<^vv<<^v>>v><<<>v>><v^v<v<>^<^^><>>>^<^<<^>>>><>>^<>^<v>v><<<><<v><>v>^<>>v^vvv<^>v^<v>^>^^><>v>v>v><<v^v^<^>^^><vv<><^>><^>^^^^>^<>v>vv>^^^><<>v^^^^>><^v>^v^^>^<^<v>^<>v<v^^>>v^>^>v<<><><v^^vv^^^<>><<^<>v>>v>>><^<<^v<><v
^<^<>>vv^^<v<^<<<^vv>v>>^vv<^^^v><>><<>^vv^><<<<>>>vvv<^^v<v>>vv^v>v<<>^><>v^<^^^>v^<v><v<^>>>^^v><<>^<<<vv<><>v^v<><^>v<^>>^vv<v<v>^><>>^<^>>^^<<>>>><>^>v^^^v^><<v<^v>>^>>><^<^>>><>^v>>v^>^<<>>>v^v><<>^v>v<^v<>v<vv^<v<^><>v^vv^^^vv^^v>vv<<<^v>>v<><<^<<>>>v^vv<^<><v^><^v>v>^v^>>v^v^>vv>>>>>>>v><<>v<><>>>^><><^v^^^^>vv>^>vvv>v<^^^^<>>><v>>^<<<>vv^<>>>^<v>^^^>v><^v<<v>^^^>^>^<<><><>^>v><^<v><^<>v^<^>>^<^<v<^v^^v^<^<<^>vv<>v^^v^vv^v><<<^>v<^>^>>^^>^<><^^>>>^>><<v<><><<<><<v^<^v<^vv<v>>>>^<>^^<><><>v>^^v><<v>v<>^^>^<v>v^^^^>><v<>^>^<><^<^v>v>^^<v^v^v<>^v>><^><>>v^^><v^v<>>v^v<>^^^^^<>>vvv^^v^<><>><vv^><<<><>^vv^<vv<vv><^^<<v<v>>^<><v><<<vv>^>>v>vv^<v>>>>^>v^<^v><>vv^^<v>>vvv<>vv<^v<v^^v<vv^v>vv>vv><^<^^<v^<>v><<>v<<^<v<vv^vv<>^vvv^>^<^^^<^vv<v><<v<<><v^><>><v^<>>^^^^>>>>^^^^<<vv<<v>v<<><v>v>v<^v>^^>v<>><<v<^^^<v<^<<<<>vv><v<>>v^>><>><<v^>>^^<^vv><><v<>v>v^>>v^^<v^<>^>v^^<v<v<^<>^^^^<^v>v<<v>^<>^^<<v<^^^v^^>>v>>v<v>><<v><v>>^><v<>vv^>^<v>><<<^>v>^^v<><<v<>>vv^^>>><v^<v^<^v>v>v<^<^v>v^v<><>>
>>>vvv>v>^>^<v>v^v^v>><>v^^>^v<v^>^v^>^vv^v<<<v^^^>^^^v<>^v<<v<><<>>vv<^v<^>vv>>^<<>>^><v^>^><^><<^v<v^^^vv^<v<^<v<^v>>^<v^^<^v^>v<^<v<>><^^vv<^v>^<>^>^>>v><>>^^^v^v^v>v>>>^<^v<^>v<v^<v^^^>vv>>^<v^v>^vvv^^^>v^^>vv^^<><^<v<^>^<^>^>><vv^^<<v<v<<<><<^v<^vvvv>>^<<v>vv<vv^^^^v<>^>v^>v><><<v><^>v^>^<><>^v<<^vv><<>><><^>vvv>^v<<>v<<v>^^>><<v>v<>>v><v^>v<<<<v>>^><v<v^^^v^^>v^>vv^>v>v^><v^v^^^>><>^>^>><>>>^v>^>^><^vv>^<>>^<<v^<^><^<v^v<^^<<>^<>v<v<^>>v>v^>>>^>^v<>^><<>v<<^vv^>v<^^<<^v^<v^>><^^<>>>>>v^^vvv^^<^<^<vv><>vv<><v^v<><^>v<<v<>^^^>vv^v<>>^>><>v<<<<>^v^>><vvv>vvv^>v><^v<>>v^<>><vv<^><v^<^v^vv>^^v^<^^><^^>>^vv>^><<v>v>>v>>v^^>><v^>v<v^<v>v^>>><^>v>^<^><^><>>v>><<>><><^^^<><^vv<vv<v>^<^vvv>><v><^v<^v^^^^<v^<v<>^>><^v>^^^v>^v><vvvvv^<^vv<>>^>>^^>^vv<<^<>v>v<v<><>v^^^<>v<<>>>>^>v>>v><^v^^<<v<^<<^>^v<>><<^>^v<<<^><v>^^v>v>v<^<<v<<>>v>^v<<v<v<<<vv^^^^<v^<vvvv^vv^v^<>^<v<<v^v<^^vv^v^<^v^<vv^v^v>>v^<v^<vv<vv^^^^<>^^^v^^<<^<>v>^^^><^>^>^^<<^>>^v><<<^><vvv<v<vv>vvv<v^^v<<<>>^<<^^v<v>v>vv^vv><<>v>>
>^^<v^<>>^>v^<^v<>v^^v>v<><<v<^^>>^><^>>^>>><>^<^v^>v^^<><<vv<<>>v<^<^>v<v<^>>v>v>v^<<>vv^v<<v>>>>vv<^^^>>^>v^vv^<^v<>vvv>^>^v>^>^v><^^>vv<vv>vvvv<<<vv>><v>^v>v^v^v<^<>v>^>>^<^^>>v>v><^<>^<<v<^vv<><v^>^^<^^<^^<<v^>v<v^<^v<^><<^v>vv^^<<v<^<vv^^<v<>vv>v><vv<vvv>v^>v<<v^v^>>^<v^>v^^<>><v^<v^^>>v><>v>>v^vvvv>v<v<vv>^v>>^^<<v^>^><v<^v<^^<^>v<vvv^^v^>v<<^v<^<>^v<^>^^v><>>>vv<^<^<<^<v><^>^><<^^v^>v>v^^^^<v>>^>><^^vv<^v<<><^^v^v><><<<^<vv><>v<<^v>^<^><><v^><v>v^vv^><<<<^v<>v^^^<^^^v^<<<>>^^<<^><<>>><v^><>^^^v<^^<^>><v>^><v<<vvv>>^><^v<>^^<^vv<^^^v<^v><vv<v^<v>^>vv>^v<<^><^<^^>v^><<^^v>>vv^><<>>v<>^>^<v<<<>v>^^><vv<<<^<^v>v<>^v<^>^<<vv>^<v^^v>v<>v<>vvv>>>><><<v<>^>v^>v^>v^>vv^>v^vv<<^<<^<^^><^^<v>^<^vvv><<vvvv^vv<^v<<>vv^<^<^>^v><v<^^<>^v<^v<<v<><v><<><>v><v^>v^^<<<^^<>vv><^><^^^v^^^>^^<^vv^<^v<^v<v<<>>>>^>v><^^^^^<^^<>^^>^<>vv^>v<^v<v<v>v<<><v>v><<v<<>^>^^<<>>^^^v>^<v<v<><>v<vv^^^>v^^v^><>^><v<<^>><<v<>^<>><vvv^><>>>^v>^^>><^>v<>^<v^^^v>>^<v>><^<^<<>^>v<<^v>^><v^v>^v>^v<><v^<v>vv<>^v<^^<^^vv^<
>^^^<>v^><^>><>^<^vvv>v^^>v>>^<^<>><^vvv^^>v<>>^>v<><<^v^^<><v><^<^<<v<>^<<<<<^><>v<^v^v<><vvvv^^v<<v^v>v><>v><v>vv<<^<>v^^v<^v^^^^<v>>^v^><^v<<>>v<<>v^^^v^^^^<v<^>v<>><^>^^v><>>>^<v^vvv>v<vv>^>v<vv<v<v^^>>v<vv>v<v<v<v><<>>vv>v<v<^v><>^<<^<^<<v^><<<vv^^v^v>>^>v<>^vv<^v^>^>><<>v<^>v^^vv^>>^<^v<^v>><^<v^><>^^^<v><v><>>^v^v<>>vv<vvvv>v>v<^<v<<>v^^^^<<vv>vv<<v^vvvvv<vvv^v^><^<>^<<><^v^^>v>^^><v^>><<>v<<vvvv>^<^^>^^<<vvv>v<<v^^>><v^vv>vvvvv>^^>>^><<<><v<<v^<^<<<>vv><<^<<v^^^v<<><^vv>vv><<<>v<><v<^>>>^><<^^v<v<^v^v<>^v<^v^^^<vvvvv>^>v<v>>^<^^<>^v<v><>>^v>>^^<v>>^^<<>v<vv<^>>>vv><^>^>v>><<<^v^v<>^^<^vvv>^v^><<vv><><>>^<vv<^<^<^v<<<<vvvv<<v^^v>><v<<>^>>^<v<<v<v<^^^vv<vv>^^v^>vv<^vv^>^<^<>^^v><v^^<^^><>^<<<<<>^<v^<^<>v<<<v>v^>^v<>^vv^v^v^^>>>>^<v<v<>^><>vv<^<v<^^><<v>v<><<v<>><<><>>><>v><v<v>^v<^<v^<^v>v<<<<<>>^^vv^vv<^<>vv^v<<v>v^vv^v^<^<<<^vv^>^^>v>v^v>v>>^^^v>>v>><<^<>v<>v<^^><<v^<^<<>v^><<v<^^<>vv>^^vv>^^<>^<>^^v^v^^v><<v^>^<><v^<<<<<^<><^>><^^^v<v^<<<^^<><<>>^^vvv^<vv>^<v<vv<><v^^v^>><^<v^
v<v^<^<^v><<>>>v<v^<<^^>^v^v<v>>>^^><^v<^><v>^<v>^^<^^>v<<<^<v<vv^^v>>^>>><v^v<v^<^^<v>v<<^^<^^^<><<>v^>v><<^>v><>^<v<^v>>v>v>v>v>v^>>>>^v^<>>^<vvvv^><>>^v^^<v><>v^^>v^>^>>><^<>^>^>v<>><<^v<^^>v><v>v>>v<<v<<^^<<vv><<>v^^>^<>><^><^vv<>v^vvv>v^v>vv^^^^^^<v>^><>>^<<<^vv^<<>v^vv>v>><>>^v>v<v^<v<v<>^v^<^<v^^<<v<^v<><<vv<v^^vv>^v<v^^v>>>^>vv>v>>^^^^^<^v^^^>><>>><<v<^>v^vv>>^^^v<^>^v^><>v^^>vv^<^v^<^v><<>v><<^v<v^<<><vv^<>v^>>^>v<^<<>v<<vv>vvv<v>vv<<v>>v>^>v^^v><>v><^v>>>>^>^^v>^vv^<v^<^^^v><><^^v>v^<^>vvv><^v^<<<>>^^^><v^>>v<<^>>><^vvv^>v<v><vv^^v^>^<^<<^v>>v<><<^^<^v^v<>^><vv>v><>v>>>>>^><^<^<>>v^>^>^^<<><^^v<vv<v^^<>^^^<^^vv^v><>v>>^<^<^^>v^>^^v>v<v^>^^^<vv^<<>v>v<>><<v>vv<<^<>>v>><<^^^^<^v<>^^vvvv^<>^v^v<>^<<<v<><<><<^v>^>v^^v^<<>v<<vv^>>v<v>>>v>vv>>^v<>^>^v>^v><^^><^>v>v<v^^v><^^<<<<^^^<>>v^>^<^<<v><<^>><v>><vv>^^<^v^^v<>><^>vv<><vv^^<v>^^v^>^v<v<v^<><^><<^^vvv^vv^v<>>^>><^>v>^^^<<v<vv^^<vv>>>^v>>>><^v>vvv^>>>vv<^<><v><v><^><><<^v<>>^^v^^<>>>>><><<^^v^vv<v^^<<vv>^><^<>v<><>>>^^^>>>v^>v>^
<>>^><>>>v<<v<^^v><^v>>vv^v<>^>>vvv<<>><<^<>^v>>v<>><>v<>vvvv^v<^>>^vv>v^v<^>^^v^<v^<^>v<^>v<^><v^<v>v<<<^<>^>^^v^><^>^<>v>v><<<^<^<>vv><>>^vv^<>^>><>><^v<<><<^^v<<v^v<v^>^vv<<<^^>>^vvvvv>^^^vv^>>>^<>v<^^>>^v^^>><v>^<^vv>^><>v>vv^v><v^v^^>^<><^>^v^<>v>>v>>^^>>^^v^^<>v^<>><^^<v<v^<><v<>><<v<v^<><^<<^v><<<<^^<<<<>^^>><^<<>vv>vv^><<vv<>v<vv<^<<>^>><<vv>>^^>^>>>v>^<<<<<>^vv^vv<v>^>v^>vv^><>v>>^v<>^<><<^<^^<><^v^vv<>v<v>>^<vv><><^vv<<^<>>v><<^>>>>vv><>v^><<<^v^v>><^<<v^><^^^>>^v^>>vv<<v^v>^>>^>>v^^>vv^>v^^<<^^v<<<<^<<v^<^>>^^^<v^><<<v<^v><>v><>v>^<><^^vv>>v>v^^><^^v<^>^>^>v><^>^^>v<vv>>vvv<^>^^v>>v^><v<v>><^^><^>^>v^<v^<<v<<>><v<>>^v>>>^^>^>^^v<^>^<vvv<^^>^^^v>^v^<^^><<<vv^<><<vv>vv>^>>>>>>v><>^<><>v^<>>>^<v^>v^>v<^><^<<^>^vv<>v<^vv>>><v><>v>^<<>^>^>vvv^vv<>^<<>^v<^<>^v<v^vv<<>v<v^<><^v<v<><^^<^<<<<>>>>v>><^>>^^v<^<v<v>^v>^^>^><v^<>^^<^v>^^><v>^>><>^v>><><^^><v^v>v>v>>^<><vvv<^^v<<>>vv^vvvvvv^<<^><>>^<>v^v<vv^>^><^vv>v<<<v<<>>><v>>^>>v<<>>v^><>^>><^^^<<<><^^<^vv>>^v<>vv<>v<<><^v<^>^<><v<v>v
vv>>><<><<<v<^><<<>>>^<^^<<^^v^^^^v^>^^^>>^>>><>^^^><<^v<^>>>>>^vv<^<>v^v^>^<>>>^^^<><>v<>^>v^><<<<^v^^^<><^<>>v>><<^v<<><><<^>>vv^>^vvv><vv<<<<v^^<<v>v>v>v^<<v>^><v^^^^><v><^<v^^>^<>>v^>><v^<<<^v^<^^<v^v^<<^><><vvvv>^^vv>^vv^v^^<>^^^^v^>vv^><^v^<<>^<<^^<v>^^^^^>^<v<<v><^<<v^>>^<<^>><<<^<<>^^^^v^>^^v<>^^v<>vv<<>>vv<^^<^^>>v^v>>>v>^^^><<>^^<<<vv><<v>v^<^v>^^^v<><^>>^<vvv><v^v<>v^<^>>><<^^>^>v<<v^<v^>^<<v<><v<>v^<v^><>vv>^<^<<^^^<>vv>^><<v^^<>><^>><<><>^^vvv^^><<v>v^vv>^^vv^><^^v<>>vvv<>^v^v^^<vv>v>>vv^^<>><><<>>v<>>v<v^^><^<v^^vv<<<v<>v<^><><^^^^^^v<^v><^^>^<<^^><^vv<<<>^v^<>><<>>>^v^v^v>^^>v>^><<^><>vv^v>v>v^<<v>>v>vv<><<v<>>^^v><^^>v^v^<v>>><<^^>^><>>>^^<^v^^<>^v>>>^^<>vvvv^<<<^^vv<^^^v>^vvv><v<^<>>^<v<<<><v>>><v<^>v>>^>v<><v^<<^^^><<<^v<<<v><<^<^>^<>v>^>>>v^^<<>>><^^v^>^v^^<^>>v>v<v^>v>^^<>v<>v^^^v<<>^^v>v<><vvv^>vv^<^^v<>^>v^^<>>^<^<><v>^^<<^vv><^^v>v<^^^<<<v<^v>v<>>vvvv>>>^<<>^v><<vv<><><>v^>^<>>^^^><>^^^^>>>v<<>>>>>>vvv>>>vvvv^<>^<^><>^v<<^>^v<v<><>^^^>vv<<v^^<vv>^^>^<v^>^<^>>^vv>
v<^vv>^>>>v>^v<<^^<v>^>v<<<<v<v>v>v>^>><^v><>v>v^vv<>>^>^^v>v^<<>^^>>><vv^<>>v>v<><<>^><v<v<><^<>v><<^>^>>>^<vv^^<<^vv^v^>>v>>^<>>v>^<^vv>^<>><>vvvvv>^v^>^<>>>>v^vv>v^<^<v<<v<>^>^><>^<v^<^<><<<<^<^v^v<<>^>^^v>><<>vv<<vv^>v<<v^<<^<vv<><v>v<v^<<v<^^v>>>^>^v<^>v^<^^v<v^v>^^>v>>^>v>^v><>vv<^>^v<<^v<^^>>^^<v<vv<v^<v>^v<v>v<v>^^<^>v^^^v>^><^<v><>><><><>^>>v^>v^><^^<><>>>>>>v^v<<v^v<^>>>^<><v^<^>^v><<<<v^>^<>>v>v>>v>v>vv<<v>>^><<v>^v^^><v<v^^>v>><^v<vv^<<v><^^v^<^>v<>><v<<>>^<^vv<^^>^^>v>>^v^>><vvv<<^^<^^>>^>>v>v^v>^v^<>vv>^v>^>^^^^<>v>>^>^^><>>><>^^<v<^<<<v><<<><^<v<>>v<<><^>v^v^v<v<v<v<<>v>>vv<<^^^v>><<><v><<<>^>^<>>><><^v^^>>>^^>^v>v^<>vv^>^>^v<<<v^<^vv^vv>>^v>>><^v>^^^<>>>v>>^>^><><<>>^><vv>^>^<>><v<<>>>v<^><v^v>><<^^<vvv^<<<<v<v^^>^><>>v<><<^v<<^<^<<>v<v>vv><<>v^^^<>>vv>>^^v^^<>>^>><>^^vv^<^^v<>^^v>v>><^vv><v<v<><^v^^>v<vvv^>v^>v>^^vv<<<>>^<>vvv<v^<><><<<^<>^v<>v^<v^<^^vvvv<<<^^<>^v^^vv<v><^><v^^^>>^v<v<<^>vv>^^><^>^v^v>^>^^<<^v^v^v^<v^>vv^<>v^^><v>^^^><vv<vv^^>v>^<^>><v>>>^v<^><>^vv^>^v
>v^vvv^^>><^^^>^<><^<>v><^<^v^^<<^^<vv<^^>>v^<^<><>^^>^<>^<vv<<^vv><v<^<v<>vv>>^^^<^><^v<><^^>^v>v<<>v>><v>>vvv<^v>><><>^v<^>>vv<^<vv^^<v^>^^vv<^^<v^><<^><^vv<>>v<><^>>^vv<<<^>^<>v<<<<>^<^v<vvv<v<^<vvvv^>^<^v<<^<^^<<^^^v<>>>v<v>vvv<><v<^v<>^><<>>>^vv^>^^v^^<v<>><<<vv^vv^v<>>>>v>^<^>>>^<>>v^^>^>v^<^^>^<<v^><<<<>>>v<v>>>v<>v^<<>^<<>^^^<^v>>^^v>vv><^vv^>v^v^v<v^<>>v<><>v><>^<><>>v^v^<<v^<^>^<<v>>^v>>^v<^^>>^v>><^>^><v>^^^v^<>v<vv><^>>^^v^<^^^><<^<^vv><>vv<<<<^^>^^>^v<>v<>v<v>^><>^<><<>vv<^<vv<^vv>^<^vv^<vv^^<^vv^^><^<>><>vv^>^<>>v^<<>v<^v<>v^<^<v<^<<<v>>^v><^><<>v<<^<v<>v^v>^<v^><>^<^<>>^<^<<><>>^^<^<<<<>>>vv<>^><>v><<^v>vv>^>><v^vvv>v^><>v>v^^v^v<<^>>^>^v<><><v<<<^v^>v<<><<<^<v<^^v>v^v>>^<<^v<>>><v<^<>>>v<>^>>v^<>>^>vv^v>^vv^>^<v<><v>^^>><<^v^v^v>^v>><<v<>>v<>>>>^vv^<v^<v<^<v^^<<<><v^<^>>v^>vv>^><>><vv^>^>><^^v<^>^<^v>^v>^^><>^><>^v>><v>>vv<<^<^>>^^<><>v<>^vv^^>>^><<v>^<v^>>>>v^<>^^^^v>v^>>^v^v^v<vv>^^^v^<<v^<^^>^><<<>>^^><v<>^>vv<^^<v^v>>^<>v>vvv>v<<^>>^><><<v><v<v<v>^^>^>vv>^^><<<>v>>^
><^<vv<^>>^^>^^vv<<>v^><^<><^<v><v><>^>^v^<<v>v<><^v<^<v><vv><>>>>><v^<><<^v><<<^<<<>>v>^>^<<v^v>v^v><^>v^^v><>vv^^<>>^v^vv<<v>>>^<v^>><>>vv<<><v><v>>^v^>^^<^^>^<vvv>vvv^<^<<<>>v>v>^<>v^<v<^v<>v^><v^^vv>>v><v>^^^<^v<^^>>v<vvv^><v^v>v<vvv>^>>v^>^^^^><<^vv<v<^^^>^v<^^>>vv><<<vv^<vvvv><v<><^^>^>v<^vv>v>>v>v>^^<^>vvvv>>^v>v>>^>v<v^<>>^v^<^<<^<^v<<<^v>vv^^v<v>>>><>^><v>^vv^^v^v^>v<^>><^<<>^>>>v^<<><v<<v^<^^>>^v<v<^<^>v^v^><v>v<^v^>>>>>^>^^^^<<>>>vv<>>vv><<<<v>v^^<^<^^<^><^v<<>vvvv<v<^^vv<v>^<^<>v>vvvv<^<>^<<vv^^^v>^<<>^>^vvv<<vv^^<^>^>^v><<>^<v<<>>>vvvv^vv^>>^^^<<<^v^^><<>>><>v>vv>><<>>v^>vv<<v^><v><^<>><<>v<v<<vv^vv>vv^>^^^v^^<><^<^>v>^>v<v<>v<v<^vv^^^^>v<>^<>^^><^vv^v>vv^^>>>vv^<^vv>^^>v^v^^<>v<^><v><^^>v<vv<^><>^^^<^^^><v>v^>><<<v^<>>^<^v><<v^v^>>>>^>^>^><>>^>^>>>vvv<><<^>v<<>v^^>v><v><^v>vv^<^v>^>>^>^<v<<<<^<<^^<<><<^^<^>v>v<^>>^v<<<^<>v>><v^^^v^>v>>^v<<<v^<v^v><>><vv<>v>>^<<<<^v>^>v<v><>^><>vvv^v<^^v>><<^<v<>v>^>^>^>v>v>v<^>>v^^<>v>v>^^>>><^>>><>>v<>v^>^^^>>^^^^<<v<>vvv>>^>^^<v<v<<v<>>
^<>>vv>v<^v^^v<>v<><^<^>vvvv<<vv><<><>>^^>v><^^><<>^^v>^<<v<^<^v>>^<>^^vv>v^>v^^^v>v<<v>v^^^<>^<v<vv<^<>>^^v<<vv><v<><vv^><^>^><<>vv^<<^>^<^>^<v>><^v<<^v>^vv^<vv<vv>v<<^v><>>^^<>><><<^v^v^<>vv>>^v^^vv>^^vv<<vv^><^>^v<>v>^vv^^>^>v^v<>>vv<v^><<><vv^^>^^^<<>^<><<vv><>><><v<v<v<>^<>^^^<><v>^<><^v<><<^>>>vvv^<<^<>>^<>>^^><<>v>^>>><>v><<>v>v^>v><vv<<vv^v>vvv<vv<v^^^v><<>^<^<<<<><<^^^>v<v<>><<vv>>v<v><><>vv^<^>v<v>><<>>^^<v^>vv^>>>v>^>>v>^<><<vv<^>><^<^vvvv>>vv^^vv><^v^<v^>^^<<^>>v>v^v^^^^v^v>^>v<^<v><v<v^><^>^v<^^<<^v<<v^^>v<^^<>>><vv>vv^>>^<<v>^<<>>>v>><>>>v^<<<^^^<v<v>^><^<<<^>vv^^v>^v>vvv<<>^^^v><vv>v>v<><<^vv>^>>v>v>^<^<>^>^^vvv<<<^vv><vv<<<<v>^^^v^v^<^v<^v^^^v^^^^vv^<>><^<vv^vv<vv<^><^>vv^v<v^v^<<<^v^vv^>>^<<>vv>vvv^vv<>^<^v^<<v>v^^v^v^>^v<<>v^v<^^<^<>><>^^^v>^<^^>^>vv<^v>>>><<>vv<>v^<^<^>v<v>^>^>^<>v>v^v^<v<vv<v>^^v^>vvv<>^>>v<^>>>v^^^v>v>^<>^^vvvv^v^^vv>>>vv>>^^>^^<<>^>^^<^<<>><^>vvv<><v^v>>^v>vvvv<>>>v>^<><^><vvv^v^>vv<<vvv<<><^v<^v<^<<^v^v>^>>>v>v^v<^<<>>>>^v^vv>^>><>v^^>vvv^^v^>>>v
vvv<>><^v^^<^^^^vv>v>>>><<^vv^v><^<<v<^v><^v^v^<<v<<<^v^v^<<<>v>^^><v<^^><^v^v^^>^>>vv<<>>><>>v<^^<v>^<<<<<^<v>><>^^^^<^<^^v<>v<^<<^v^v^>>^^v><<^v^v<^v<^^v^<<^^^v>vv^>>v<>^><v^^v^>^^v^v^>><>^vvv<>^^<>^>v^v<>>^v^^^v<^<v^<vvvv^<v^vv<v<<<v<v^>>^><^>>v^<v^vv^v<><^v<^v<^>^<<>>>^^v<v<vv<<>vv<><v>>^v^<<<>^>^^<<<><<v<><>v><<<vvv^v><<v>v><v^^<^^^>><^<^>^v<>v<^v><vv<v>^^^<><>^<^v^>>v<>^v>vvv^>^^><v>vvvv<><vv<^>>vv<^<>>^>v>v^^>v><<>vvv<<^^vvvv>^<>^>v<^><>^^<<vv>v<^><v<^v<v^v>><<v<^v>v><>><><^<<><vv<v^vv>><<^<^^>^>vv^^v^<<<^v<<v<^v^<>>v^<<v>vv^>^<vvv><>><^<v<v^v<<^v>>>>>^^^v>^v^vv<<^>>><><vv^>v^>^>v^>>^^<vv>^vv><<>>v^^<^>^><<^v^<>^^>><v^vv^><v><^v^^<v<>>v^v^^v>v^v>>v<<^<v>^>>v>v>vv><><<vvvv<>^<v^vvv>^<><>v>^>>^<^v<v<v<^<^>v>^vv<^vvv<v>>><>^>>^>^<>^v>v><v>>vv<<v^^>v>^^<^<vv>vv^v>>^v>>v><^<<>>^^vvv^>^<v><<v><>^v^<^v^<><<^^<<<^<^v>><<vv>v^>>>><^v>v>^v<^v>^v^<>>vvv<^>>>^>^^><<v<v^v<<v^^^>>>><v<>^^^vv>>v^>v>>^<><><<v>>^<v^^>v<<<<^<^>^^><<<^>v^vvv<v^><<<^>v>^<>^^v^^>^v<^>^>^vv<v^<><^<v^><v^^>^><v<^vv<^>
><v>^>>>>v<^<><><>>><<>v^^vv>vvv^<^^v>>v<v^<v^<v^<^>>^>v^^^v^v^>><v<v^<^>^<^>^^<^v^v<v<><^v>><>^>v>><v>^<^^^<<^vv^>>><<<^v>v><>v<<>vv^<^>><v<><v>v<<^>vv<vv>>v<v^<>^vv>^<v<v^^>vv><><^>v<vv>vv><^<><v<v<>v<<<>>><><v>^vv<^<<v^<<^<^<v>^><><vv>>^^>>>v>><<^v>vv<>>^v<v>><<v>^v^><v^<^v>^v<v><v>^vv^<v>vv^>>v<vv<>>^vvv^<v^vv^>v>vv^<^<v<^<<v^v^v>vv^><<<^<>>vv>>^^>v<^>v^<<>^<>>v<v<<^<v>>^v>v>><<<^^^>v^v^v^<v>>^<^^<>^>>v^>>>v^><<^vv>v^^v>^v^^v^<^vvvv<>v<>v^<v^<><^^>v<^v<>><v>vv><v>><<^<<<v>v^v>vvv^<v><v>><<^>^v>>vv<v^>v>^><v><^<v<v<^v^>^vv^^v<^<<<v>^^v<^<>>v><<<<^v<<>^><>>>>v^vv>v^vv>v>^><><^><vv^><^><v^<v^<v>v<>vv<v><vvv>^v<^^^<vv><v>>>>vvv>>v>><<><^v<>>><v>v>^>^>>v^^><>v^>^<vv><>>>^^v^^>v^>>v>vv>^^^<^>v^<<>^<>v>v<v<v^v^vv>>>vvv>^<vv^<v^^v<vv^<>>^<^v>v>>>>>^<>^<^v>v^>>>>>^v^<v><^v^><^^>^<v^^v^>v^v>^>>><v<^v>>>>>><<<^^^^<^v^^<<<<<^>>^<>v><^<^<<v^>>v>><v>v^>^>>^v^^^<>vv^><v>^v^<<^>^<<^v<>>^<v<>v<v^v^<>^>vv><>^^>>>>>^^v>^^^vv<<v^>v>^<><<><v>><^v^<>v<<>^vvv^^^v<>v>>^<^^^vv<><^>v^>^^><>>^^>v<v<<v>>>>^vv
^vv>><<v<^>v<v>>^v^<>^vv<v^><vv<>v<<v<<^>v<^v^^<>v>v^>^<v<v>>vvv^><<^v>>vv<><>^><^^v<<<>>^<^>v^<v^>>vv>^v>><<vvvvv>>^^>v>>v>v^>>v^<<^v^vv^^<<<<><<^^^^<^<v^v<<v><>><^>><<<v>^<<vv<^<>^>vvv^vvv^^^v<^<<><v<<v^^v>^<^<>^><v^<v><<>>>><>><^v>>^<<v><>>vvv<v<vv^>^>v<>v<v^<v^^<<^v^<vvv<>>v^<<^>><^v<^>>v^v<<<v><vv>><^<<>^>v>v^<^v^><^<<<v<^><>v^v>v<^^^^^v<v>v^><^v^^<^><^vv^^<^>>><^^v>>>>v<<>v<><v><<vvv>^^<>>v>>v^v>><>v><^^>>vv^><v<^v<<>v>vv^<v>vv<><<^^^^><vv^<<><v<<^v^><<>v^<vv<>>^^>^<<>><<v<^v^^>vv<<^v^v<>>vvv<><>^<>v<vv>^v<v>^vv^>>>vvvvv^v^<^vv>^>^><<v><>^^v>>>^>v<^>v><v><<<^^>^<v^<<<<<^^<v<><^>^v<<<^^>^^^vv^^v>v<v><^^^^>^>><vv<<<vv^v<vv>>vv<>>><<v>v><><>vv>^>><>^v^^<v^v<><<<<^<>><^v^><v<v<v^><>^>v><^<^^v<^><>v^<>>^v<v^<<<v>><^v>^v^^v<><v><^<<vv^<^>^>^v>v^<vv^v^^<^v^>>^v>vv<<>^<v^<<>>v<^v<v^^^v^>><<v<v^v<>v<v><<>v^>>>v>^<v><^^>vv^v<v>vv^v^v^>vv>v>^>v><vv<v^>><<<^<><<>^^^><vv><<^v<^<>>^v>>v<>v<><<><>>>v^<><>v<<<v<><><v^<^v><>vvv>v<^^v<^<^v<vv<vv><>^v>^>>><<^vvv<>^v<>>v<<<><<v>>^<^>^v^>vv>><vvv^^^v
<<^><>^><^^^v<v^v>>>^<^<>^<^^^<^vv>>><<v<<<<v^>><><^<>v<<<>>v^<<^^>>v^^vv^<>><v>vv<>^v>><><v^^<<<^^^v^<^^^^<<<<>>^^<<^vv<<>v^><>>^>v^>vv^v><<^vv^v<^>>>^<^<>v<^>vv<>v<><^<^v>v^>^v<<v^>>>>^<<v>>^>^vv^<^v<v<v>>>><<^v<^vv<>><<^vvvvvvv^^><>^vv>^vv^<^^^^>^v><<<<>^v>><^>><^^>^<>v><^v<><^<<<v>^><v<^v>v<^^>^^<v>v<^<<>^^>^<vv^>vvv>^>^v^^v^>vv<<^v>vv<<^v<<v>^v<^<^v<^><>>>>vv>>>>^v^>^^><v^^>>>v^><><v^v>>v^^v>^^^^^vv>>v<v>>v>^^v^<^vvvv><>v>v<<^<<>>v>>>><>><v^^^>^>^>^><^>>>^vv>^<v><v<<<>v^><>^v><>vvv>>>^^<<<>^>^vv>v^^v^<^>vv>^<^>>^><<<<<<<>v^<>v^v>^vv>v><>^>><v<v>vv>><v>v><<<>>^v<>vv><<v^>^v<v^>v>^>^<<<v^vv^v^<v><>>v^>><<>>^v<^v^<<><^v<^vvvv<<^^v<<<^vv^><v<<vv^>^vv^><^^^<^^>>v^>>^<^>>><v^<><v<><^>^vv>v<^<<^>vv>><>>^vvv>>^^<<>v>v^>><v<<<<<<v>vv<v>>>v^>^^>>>vv><><^><>v><<>^^^><^>v><><<^v<>><v<><^<<^v><^<>v<v><<<<<<v<^>^v><<v^^^>>>>>^>>vv^^^v<<>>^<<v<<>^>v>>^><<v<<<<>>^vv^>vvvv^>^v<v>v>v>^>v<><^vv>>vv<><v><^>>v^<v^^vv<vv^>^>>v^vv<><^vv><<><<<vvvv>^<<v><>>^vv<^<>><>>v>v^v<<^^><>^<<><><>v>v>>^>v^<^^<v^<v
^v>^v<>^>vv><<vvv^^vv^<^^v<><<v^<v>><v<^^^><<<vvv>v^^>>^v<^>^^^>vvv^<<^<vvv>^v>><<<^<^vvvv>v<vv^^v^><>v^^v>v<^vv><>^>^v>^<^^vv<v<^>^<^><>>>^<><>^vvv^>v><v^<<>>>>^>^v<v><<<<v^^^<>v<<v>>^<v><^><^><v<<^vvv><>^<><^><>v<v<><>>^>^<vvv^>^^><^^<<<<<><^^v><>v>>>v^^<>^>>v>^<<v^><^^>>^>v^>><>>^^^>v><<<<^>^<v><^><^v<^<>^^>^vv>v<v<^^>^vv<<>^^<v^v^<v^v^vv^><^vv^^<>>v^v<v^>v>^v^^>^v<>^^v<>>^>^^<<<v><^^>>^<>>^><^v<><<><^<<^^<>><<vv^v<<>v^<<v^><>^<>^><<^v^<vv><>v^^<^v^v<>>>>^vv>vv>vv>^<^<^><>><><v<v^>>>^vv^<><>><><><v<>^>>>v^v>^<^><^^v^^>>><vv<^v>vv>^^v^v^^^<<>v<^v^<<<><<<<v^v<><v>>>>>^vvv>v^>v>v<^v<^vvv^>>v^><><>v<>^><^^vvv><vvv<v^<>^v>><><<>>>vv^>>vv<v<^^<>vv<^^<>><v^v^<v><<^v>^v><^v^^^v>v<v>v<^^>^<>v^<^>^<<>>>><vv<>v<>><>v<<^<v<<<^^^v^vv^>><v<><>^v^v^^v^>>v><v^><vvv>^v>>><vvvv<^>v^><v^>v>>>vv<vvv^>^^vv^vv<<>vvv<^^v^<v^^^>v<v^^><<<v^^>>>^<>v>^v<<v^v>>>v^<v<v^><^v><<>>^><^v><^<^<<^^>v^^<^<<>>v<^>vv^><><><vvv>v>v>^>>^^>>v^<^>>v^^>^<^v>^^>v>>>>>vv^>^v<><^>>^>>^>^v<>^v<^><<^<^^<vv^<vv^<v^<<>v<^<>>^^>v<>^
v^<><^<<>>vv<^<^>^^v><^<^v^>>>^<v^v>>>^>>v^v>v^>^^^>><^^^>^>^><>^>v>^>^>><^>^>>v^<><<v^<^<^<<>^>^^>^^<<><<^^><<v>^v<vv^>v^^^<<^<<v^vv<>^v<>v>^^v^>v>^v^^^>>^<<<>^^v>^>^>vv<<<<><v^>>v<><^v<v^>>v^^^<>^><v^v>vvv^v<>^>^<>^<>^<<<vv<<^<vv^<v<^<>>vvv<><>vv><v>^v^v<>^<v^<<<v^>^>v>^<<>v^<v^^v>v^<>^vv<v><>v^<<<^vv^v><<>><^v^><><<<<>vv>v<^<<><vv^>^<^<>v^v<^>v>vv<v<<v^v><^^^<^<^^v>vv>^>>v>>>v><vv>^v>vv><v^>vvv^v^>^>vv>>>vvvv^><^v>>>vv^^^<<>v^><v^>>vv^><>^^^<>^^><<<<^^>vvv<v><>><^vvvv<>^><<>>v<^^><>>v><<^<<^<v^><>><^^v<<<><^vv^^^^^><<v>vv>v>v>v<^^v>v^>><<^v<<<<>v<<>v><v<>^<^v>v<>^^v^<v>v<><v^<^>v^v>>>v><<v>v<><<^^>v^<<<>vv^>^<v<<<^<>><^^v>v<<<<<<^^v<^<^><^vv>^^<>^^<>v<<>>vv<><<<<>v>^v^>>v^<v<>v>^<v^v>>><v^>v^<^>^<<^<v^v><^>^^>>^<<v>v<>>><<v^v>v^vv>v<^<>>v^v^>^<^v>>><^<v<v<<><<>>v<>v>>>vv>v^^^>^v^>^^>>vvv^<>>^^>>^><^>>^^<^>>^<^^>vv<vvv^<<>>>v>>^^>>vv>vv<>^^^>^vv^<v>>vvv^<<<>vvv>^^v<^v<v<v<v>^>>^^v>^v>v^><vv<v>^<<>><v>^>v<v><v<<<^v<>><^v><^v>v<v^>v^>>v^<>>>^^<vv<<^><^<^^>^>>v^v>^<><^>^v<<>v^<><^<v^^<>
<>>>^v><v^^^>v<>v^^v<>>^>><>v<^>>><<>>>><>v^^>v<v<>^<^v<<^<^^^<<<<>^^^v<^>^>v^<<><<v<^<v<^<><>^<^v<>><v<vv^<v<>^^<vvvv<>^^vvv<vv^>^vv^<<<<v><><<<vv^^^<v<^^>^^^^^v<v^vv<^><v^v<<v<<v<^<<^<>vvv^v^><^vv^><vvv^<<vv<v^>^v<<<>v^<^^^^^<vv<>^<vv<v>v<<>v^vv^>>v<<>^^>>v>v^>v^v<<^^v^^v>>>vv><<>>>^<v<^^<vv^><<<><<<^v>^><^vv>>>vv>>>><<>vv>>>^^vv<<vv<<>v^>><^>^<<^>v<>v>^<^vvv<v><<<>v^<^^v<<^>>><v^^v^<v>><v>>^<^^>v>>^^^v>^>>>>vv<>>>>^^^^<<^<><>>vv^<>v<^v<<vv<^>>^<<>vv><^^v>vv<vvv<v<v<<<^^^v>v>>^>^<><<^^^^><v<>v><^^<>>>v<<^v><<>^<<^^>^<<<<>^^^v>v<^<v<v^^>v<^^>v<^<<>v>><^^>v<v<vv><^<^vv><v>^^>^v>^<^>v>>v>^<>vv><vv>^<^>^^^<^<v^<^<><v>^^<v<><>v>vv<^^<<<v>>>>v<>v>v<^^<>^<>^^^>^v^^v>>>v><v><v<>>v>><vv<^v^v<<<<vv>vvv^v><<><^>v<^<v<<>>>vv<<^^vv><v<<>^<v<v>^<^>^>v<v^vvvvv^v^vv>vv^<vv^^>^>>^^<v<vv>^^^^v^<>^<<>^><<><v<>v^^<<^v>v^vvv^v^>>>v>^^<><<v>^><<^^<>><<v^<>^^>^v>^<vvv<^>v<^><^<v>>v><^<^vv<v^<>>v^<<>v^^><v><v<<>^^>vvv<<^><^^<<<<vv^>^v<v^v>^^^^>v<<vv^^>>>v>^><v>^^<<^^v<>vv^>v>^>>>^<>^>v><v^<^>^v^v^^><^>v^>^<`
}

func getTestInput01() string {
	return `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`
}

func getTestInput02() string {
	return `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`
}

func getInput() string {
	return getInput01()
}