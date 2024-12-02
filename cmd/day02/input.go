package main

func getInput01() string {
	res := `35 37 38 41 43 41
64 66 69 71 72 72
45 47 50 51 52 53 55 59
16 18 19 20 23 29
36 39 41 43 44 41 44
42 45 46 44 42
82 85 86 87 88 86 86
42 45 46 45 47 51
57 60 62 65 63 65 70
88 90 90 93 95 97
54 56 58 59 60 60 62 61
64 66 68 68 70 73 73
5 8 11 11 15
82 84 84 87 89 91 97
62 65 69 71 73
46 48 50 51 55 58 59 57
10 11 15 18 18
26 28 31 34 38 41 45
40 41 42 45 47 51 53 60
19 22 25 30 32
77 78 84 85 88 85
25 27 32 33 34 37 37
77 79 84 87 89 91 92 96
10 12 18 19 25
52 51 54 57 59 62 64
65 62 65 66 63
71 69 71 72 73 75 76 76
70 69 71 72 73 75 79
9 7 10 11 12 17
18 15 18 19 17 20
24 22 23 26 29 30 29 26
94 91 89 90 91 91
18 15 16 15 16 20
72 69 67 68 73
66 63 64 66 66 67 70
78 76 76 77 79 76
59 56 57 57 58 61 61
10 9 12 14 17 17 21
3 2 5 7 8 8 9 16
78 76 79 81 85 87
31 29 30 32 35 36 40 39
75 73 75 77 80 84 84
65 63 65 69 70 74
75 73 74 75 79 80 81 88
55 52 53 55 61 62
68 65 66 69 75 74
6 5 8 15 15
70 68 69 72 79 81 85
11 9 12 19 20 27
33 33 35 37 39
88 88 89 92 93 94 96 95
69 69 72 74 77 78 78
1 1 2 3 7
58 58 61 64 66 69 70 76
9 9 11 13 14 16 14 15
40 40 42 44 43 46 44
26 26 24 26 26
74 74 77 79 80 81 78 82
72 72 74 72 75 78 80 86
45 45 47 47 48
48 48 51 52 53 53 52
86 86 87 90 90 91 93 93
64 64 66 68 70 70 74
46 46 46 47 48 50 57
39 39 41 45 47 48 49
2 2 5 9 10 8
36 36 37 41 42 45 48 48
8 8 11 14 17 20 24 28
13 13 14 18 24
50 50 51 54 61 63
54 54 55 58 60 67 64
19 19 21 27 29 29
10 10 17 19 23
31 31 34 36 37 44 50
34 38 39 40 43 46 47 50
12 16 18 21 24 27 29 26
75 79 82 84 87 87
36 40 41 44 45 47 51
58 62 63 66 69 71 72 79
33 37 35 36 37 38 40
95 99 97 99 96
20 24 26 23 23
81 85 87 90 88 89 91 95
37 41 40 41 43 45 47 54
17 21 21 24 25 27 30 31
9 13 13 15 17 18 15
8 12 13 14 14 14
31 35 36 36 38 42
78 82 82 84 85 88 94
25 29 31 32 35 39 41 43
32 36 40 42 40
22 26 30 33 33
54 58 62 65 69
61 65 69 72 73 80
52 56 63 64 66
76 80 86 89 86
6 10 11 14 19 19
8 12 13 14 15 22 23 27
48 52 53 56 58 65 72
6 13 14 17 20 21
65 71 73 74 77 78 81 80
80 85 88 91 92 95 96 96
79 85 88 89 93
36 41 43 46 47 49 55
54 59 57 58 59
77 84 86 84 85 83
2 9 11 12 9 11 11
82 88 87 90 94
22 27 30 33 35 38 37 43
72 77 79 79 81 83
82 89 92 94 96 96 94
65 72 73 73 73
42 49 52 52 54 57 61
76 81 81 82 83 89
20 25 26 30 31
6 11 12 16 15
20 26 27 31 33 36 36
7 14 17 21 25
37 44 47 50 51 52 56 61
53 58 60 63 69 71 72
40 47 54 55 54
58 65 72 75 77 80 81 81
13 18 20 21 27 31
51 58 59 66 73
64 61 60 59 62
87 84 81 80 78 77 74 74
98 95 93 91 89 87 84 80
42 40 37 35 33 26
77 75 72 70 72 69 67
86 85 84 86 88
68 66 63 66 66
63 60 61 60 59 57 54 50
86 83 84 83 82 79 74
81 80 79 79 77 74
76 75 75 73 72 69 71
20 17 17 15 15
77 76 73 73 69
63 62 61 61 60 55
89 86 82 80 78 75
92 91 90 89 88 84 82 85
87 86 83 80 79 75 75
86 83 81 77 73
79 76 75 73 69 62
88 86 81 80 78
88 86 85 82 77 80
57 55 54 51 49 42 40 40
47 44 41 35 33 31 27
29 26 20 17 16 14 7
49 51 48 45 42 40
62 65 62 59 56 54 55
54 56 54 51 51
83 85 83 80 77 74 70
37 38 37 36 35 28
39 42 39 40 39 38 37
19 21 20 17 16 13 15 18
8 11 9 7 8 8
70 72 70 67 68 64
85 87 84 86 79
25 28 28 26 24 21
58 60 60 59 56 58
71 72 72 69 69
88 89 87 87 85 83 81 77
37 39 36 34 34 28
50 51 49 47 46 42 40
60 61 59 55 52 49 48 50
95 97 94 93 91 87 87
92 93 92 88 87 83
84 86 83 79 78 76 74 69
95 98 92 89 87 85 83
27 29 27 25 20 17 20
44 47 45 39 38 36 36
76 79 73 72 68
94 95 88 85 79
89 89 86 83 81 80 77 76
74 74 71 68 67 65 63 66
70 70 69 68 68
85 85 83 82 78
58 58 57 54 53 47
18 18 20 19 17 14 12 11
49 49 47 44 42 43 44
43 43 41 40 38 35 36 36
25 25 27 26 24 21 17
50 50 53 51 46
92 92 89 89 86 83 81 79
65 65 65 62 65
91 91 91 89 88 88
75 75 75 72 68
77 77 74 71 71 66
29 29 25 24 21 18
52 52 51 47 46 44 45
99 99 95 94 91 90 90
38 38 36 33 31 27 23
41 41 40 36 30
81 81 80 74 73 72 71 68
46 46 44 41 40 34 31 32
55 55 53 51 50 47 41 41
66 66 65 64 59 55
22 22 21 19 17 16 10 4
17 13 10 7 6 4 3
25 21 19 18 15 14 16
38 34 31 28 26 25 25
75 71 68 67 64 62 60 56
83 79 76 73 72 70 69 64
51 47 44 45 42 41
46 42 45 42 45
89 85 84 86 84 81 80 80
23 19 16 14 15 11
61 57 54 57 56 54 49
85 81 81 80 79 77 74
12 8 7 5 3 2 2 5
97 93 90 90 89 89
75 71 70 67 67 64 60
36 32 29 29 28 21
40 36 35 32 29 25 24 22
98 94 90 89 88 85 87
19 15 11 9 6 6
40 36 33 31 27 24 20
78 74 70 68 61
49 45 44 41 36 33
67 63 62 56 53 54
83 79 72 70 67 66 66
30 26 20 17 14 13 9
65 61 55 52 46
91 85 83 81 78 76 75
80 73 70 68 65 62 61 63
35 30 28 25 22 22
50 43 42 39 35
50 44 43 41 36
98 92 95 92 91 90 89
43 37 36 33 30 28 30 32
47 40 38 41 40 39 37 37
48 42 39 37 38 34
49 43 41 40 42 41 35
34 27 26 24 24 22
89 83 81 81 78 80
97 90 88 88 86 86
84 78 76 76 74 73 69
80 75 72 72 70 65
19 14 13 9 8 6 3
81 75 73 69 66 63 65
58 52 48 46 44 44
64 59 58 56 54 52 48 44
56 49 46 44 40 37 30
52 47 41 38 35
80 73 71 70 68 63 64
77 70 63 60 57 57
63 57 51 48 47 45 41
35 28 26 20 18 16 9
79 75 75 73 71 68
96 96 90 87 89
42 37 36 33 35 33
46 50 52 55 57 61 64
28 24 20 19 17 17
33 31 32 35 36 40 42 47
49 46 44 45 48 49 55
14 19 20 20 22 21
19 20 16 13 10 9 6 4
29 25 19 17 10
42 35 31 28 24
73 70 67 67 65 65
39 38 39 40 43 44 51 52
57 59 61 65 67 68 68
87 94 93 94 97 98
91 94 94 93 92
59 66 67 70 73 74 79 86
28 31 28 23 22 18
55 62 65 67 68 71 73 76
29 28 27 24 19 18 16 12
85 89 90 92 90 91 93 97
36 31 28 31 24
85 82 84 87 93 95 99
36 37 34 36 42
62 62 61 54 52 51 50 45
87 89 83 82 80 78 76
22 22 24 22 19 19
29 36 37 38 38 41 43
55 58 56 52 48
38 38 41 43 43 43
63 61 59 56 53 52 49 42
56 63 67 70 73 75 75
67 63 61 59 56 53 49 47
76 77 79 81 82 87
62 58 56 55 56 55 52 52
62 62 63 66 70
52 51 55 56 59 60
86 79 78 73 71
79 79 75 73 70 69
38 39 40 38 38
62 66 67 70 73 75 79
36 30 29 27 22 20 19 14
86 86 86 88 90 96
28 28 25 23 20 20
31 37 41 43 45 49
88 87 84 86 87 89 89
82 82 81 78 78 76 77
43 41 46 49 52 55 56 56
62 64 66 70 71 74 75 74
69 76 73 75 76 78 78
94 88 85 82 82 79 79
37 30 28 27 21
11 11 9 7 2
31 27 26 23 22 22 19 21
52 48 46 45 47 45 43
59 53 51 47 44 41 39 41
66 64 67 69 71 73 73
68 68 71 72 73 74 72 71
64 57 60 57 56 57
55 61 64 65 72 75 77 78
30 29 29 27 30
45 45 47 49 52 56 60
56 62 63 66 72 76
36 43 45 45 45
83 87 88 91 94 95 96 97
33 26 25 22 18
81 78 75 73 71 68 71 65
27 25 28 29 32 35 34
59 59 57 57 56
84 84 87 88 93 94 94
49 52 52 55 61
87 90 88 82 80 77 74 74
4 9 10 11 11 18
82 86 87 89 91 94 97 94
88 84 87 86 82
12 9 9 12 13 16 18 16
51 47 44 41 39 36 34
34 34 39 41 44 46 50
22 22 20 20 19 18 17 17
76 78 76 73 70 69 65
7 8 9 12 12
21 23 24 25 27 26 29 28
37 36 32 30 29 28 28
83 80 79 76 70 68 68
13 17 17 20 17
23 23 20 17 16 11 9 5
26 26 30 32 35 38 38
32 31 30 29 27 24 24
30 34 40 43 47
62 68 68 69 71 72 76
60 57 57 59 60 63
24 20 19 19 18 14
42 39 36 35 38
64 68 69 73 71
42 35 34 32 31 28 27 24
53 55 54 51 50 48 41
66 66 62 60 60
78 82 79 81 79
53 53 52 51 50 49 50 51
11 9 12 15 19
45 49 51 52 53 57 64
17 20 24 25 26
74 77 75 73 72 67 62
81 80 74 72 70 68 66
88 87 90 90 90
45 42 40 33 30 24
54 61 64 62 65 66 69 75
35 39 39 40 42 44
85 90 92 94 97 95 98 96
48 47 44 43 42 45 44 47
30 33 35 37 36
84 84 83 80 78 72 72
3 5 6 7 10 10 14
80 79 82 79 76 74 70
28 34 38 41 47
66 61 58 60 56
94 90 89 86 83 77 76 76
86 84 86 87 91 91
80 83 80 73 70 69 71
77 76 78 80 77 79 82 86
62 57 56 51 53
67 70 71 73 75 77 83 83
28 22 15 13 11 7
50 50 52 53 56 53 56
32 32 29 25 18
16 17 16 15 11 11
34 35 32 31 34 33 33
64 62 60 59 57 55 55 48
29 30 28 27 25 23 21
14 11 14 20 26
4 7 11 12 13 14 17 22
94 96 93 89 88 85 80
79 79 80 83 83 87
48 47 48 51 52 53 55
45 50 51 52 54 52
1 1 4 5 7 10 8
41 43 41 39 39 34
54 56 60 62 66
87 83 81 78 76 73 70 73
84 78 76 74 70 70
71 70 76 78 79 80 81 80
36 39 36 37 36 34 31 29
67 71 74 75 72 73
68 68 66 63 59
39 35 33 32 29 29
35 36 38 40 43 41 43 47
34 34 32 30 32
49 52 50 53 51 46
89 86 86 88 94
35 37 35 33 32 29 27 29
16 15 13 15 16 14
14 14 16 16 19 20
41 37 35 31 27
9 13 16 19 21 21 23 23
95 95 91 89 86 84 82 83
11 14 21 24 25
30 30 24 22 19 16 13
36 40 42 47 52
52 56 58 65 65
47 42 40 37 34 36 36
78 78 80 82 83 82 86
62 62 58 55 51
87 80 77 77 75 73 67
90 86 84 81 80 79 74 71
98 96 96 95 93
11 12 15 20 23 26 30
97 90 90 89 87 85
20 20 22 20 27
78 80 81 81 83 83
96 92 93 90 87 86 80
71 70 71 73 75 79 77
38 38 39 40 43 44 45 51
74 67 61 58 55 53 53
31 38 41 42 43 45 52
19 19 17 19 17 11
81 83 81 78 76 76
93 89 87 85 85 84 84
13 13 12 10 12 11
40 36 34 32 25
10 14 18 19 23
33 38 35 38 41 42 46
62 58 52 49 45
9 13 14 11 12 12
82 83 84 85 90 96
74 70 68 66 65 63 59
31 38 42 43 46
43 47 51 54 55 55
48 48 52 54 55 53
29 29 30 36 39 45
20 20 21 23 26 27 27
24 31 34 36 39 41 45 44
84 83 87 88 92
45 44 39 37 35 33 32 33
28 21 19 19 21
24 20 19 15 17
93 92 89 88 85 84 83 79
32 32 32 30 27 26 22
21 21 21 22 21
19 16 12 11 6
32 36 36 38 45
92 92 90 88 88 81
88 85 82 78 74
67 70 68 65 65 63 64
18 24 29 32 31
55 48 45 42 40 38 37 37
22 23 24 25 27 27 28 30
2 6 9 11 13 10 15
65 67 69 72 75 80 78
60 60 62 66 73
61 61 62 61 61
6 9 12 10 7 3
83 81 83 81 81
98 95 98 95 92 89 87
82 85 87 89 90 94
64 68 69 70 73 75 77 83
68 64 63 61 59 55 52 46
67 71 73 76 76
64 63 59 58 61
81 80 78 76 75 71 70
38 45 48 49 52 53 54 54
31 32 30 33 34 37
48 48 50 47 44 43 42 38
24 20 18 15 15 12 7
52 47 46 46 43 41 37
57 61 68 70 67
38 40 38 37 37 37
60 53 51 50 48 44 42 40
69 67 64 64 63 59
59 58 57 55 52
29 30 32 33 35 38
41 42 44 46 47
54 57 58 59 60 62 65 66
26 27 28 31 33 34 36 37
34 31 30 29 28 27
29 30 32 34 37
52 53 55 58 61 64 65 67
31 32 35 36 37 38 39
85 82 80 79 76
56 54 52 51 50 49
16 15 13 10 7 6
2 4 6 7 8 11
67 70 72 73 76 79 81 83
27 24 23 22 20 18 15
8 10 12 13 14 16
23 22 20 17 14 11 8 5
50 49 48 46 45 44 41
37 35 33 31 30 27 25 23
19 17 16 14 13 10 8 7
78 75 74 72 69 67 65
10 11 14 15 18 19 21 23
55 53 51 48 46
71 74 75 78 81 83 84 86
21 24 27 30 31
46 49 50 51 53 54 55 58
19 18 15 14 12 10
13 14 17 20 21
63 62 60 59 56 54 52 49
73 76 78 79 82 84 85 87
54 51 50 47 45 42
57 55 52 50 47 44 41
34 32 29 26 23 21 18 15
89 88 85 84 81 78
55 58 61 64 66 68 71
67 69 72 74 77
32 29 28 25 23 20 19 16
29 30 31 33 34 37 39 41
70 67 65 62 59
66 64 62 59 58
84 86 89 91 92 93
69 71 74 77 78 79 81 83
31 28 26 25 24 22 19 17
96 95 94 91 90 88
82 85 87 88 89 92 94
30 28 27 25 23 22 20
68 65 62 60 58 57 56 53
52 53 55 57 58 60
38 35 34 31 29 28 25
40 38 36 35 34 33 31
31 32 34 37 38 40 41
35 37 40 42 44 46
45 43 40 39 38 37
46 48 51 53 55 56 59 60
26 27 29 32 34 37 40
90 91 93 96 99
78 79 80 83 84
24 25 27 28 30 33 36
97 94 93 90 89 88
58 60 61 64 66 67
82 81 80 77 74 72 69 68
82 83 84 87 90
72 73 76 78 81 84 85
25 26 28 30 33 36 38 39
87 86 83 80 79 76 73
23 24 27 29 30 31 32
60 62 65 67 70 71 72
54 56 58 61 63
45 43 40 37 35 32 31 30
41 42 43 44 46
31 34 37 40 43 45
66 68 69 72 75
78 80 83 86 89 92 94 95
86 84 82 79 78 76
94 93 92 91 88 85 82 80
18 16 14 12 11
71 69 66 64 62 60 59
88 91 92 93 96
54 57 58 61 63 64 67 69
96 95 94 92 91 90 87
28 25 23 20 18 17 14 12
88 91 92 94 97
51 54 56 58 60 63 66 68
26 27 30 32 35
70 71 73 76 77 78
14 12 10 7 4 2
4 7 9 11 13 16
64 63 62 61 58
26 29 32 33 35 37 40
81 84 87 89 90
29 30 31 33 36
31 29 26 23 20
91 90 88 87 84 83 80 79
31 33 34 37 39 41 42
73 71 70 67 64 63 61 58
90 88 85 84 82 81 80 79
14 11 10 9 8 7 4
50 52 55 58 61 62
39 36 33 32 30
84 87 88 89 90 93
20 22 23 24 27 28
94 91 88 86 83 80
21 23 25 28 29 32 33
27 26 23 21 20
70 71 73 75 76 79 82
24 27 28 29 30
11 13 15 17 19 21 22
58 57 56 53 50 49
73 74 76 78 79 80 82 85
49 48 46 45 43
25 28 29 31 33 34
79 77 74 72 70 68
78 76 74 72 69 67 66
86 84 83 80 78
69 67 66 65 64 63 62 59
26 24 22 20 18 15
63 66 67 69 71 73 76
45 44 41 40 38
92 91 90 88 87
16 17 20 21 22 23
12 14 15 16 19
76 79 82 83 86
28 26 25 23 20 19 17 15
47 46 44 43 40 39
97 94 93 90 89 88 85 83
8 11 13 16 17
4 6 9 12 15 18
72 69 68 66 63 62 60 57
67 68 69 71 73 74
36 35 33 32 30
59 56 55 52 51 50
8 11 14 15 17
57 60 63 65 67 70 73 75
72 75 76 78 79 81 84 85
9 11 14 15 18 20
71 72 73 75 77 78 79 81
60 59 58 55 54 51 49
46 44 42 40 39 36
70 71 72 73 75 76
70 69 66 64 62 61
63 66 67 70 73 74 75
22 25 28 29 31 34 35
97 94 91 90 87 85 84 83
95 93 90 89 86 85 84 83
92 89 87 86 85 83 81
75 72 70 68 67
85 82 81 79 76
34 36 38 39 40 41 42 44
29 27 24 22 21
36 35 33 30 28 26 23
63 64 67 68 70 73
76 78 81 84 86
35 33 32 30 27
31 28 27 24 22 19
46 43 42 39 37 34 31 29
71 72 75 78 79 80 83
4 5 6 7 9 10
66 67 70 71 73 74 77 80
75 74 72 71 68
84 82 81 79 78 77 75 72
34 33 32 30 29
47 48 51 54 56 59 60 63
93 90 87 86 84 83
24 25 27 30 32 35 37
75 77 80 81 83 85 87
82 79 77 75 72 70
74 75 78 80 81 84 86 88
82 79 78 77 76 73 71 68
69 66 64 63 60 57 56 54
21 18 16 13 11
38 39 42 44 45 47 49 52
73 74 76 78 81 82
12 15 18 19 20 23 25 27
30 33 35 36 39 40
79 81 82 85 86 87 89 90
15 16 18 21 24
34 35 37 39 42 43
10 11 14 15 18
35 37 38 40 43 46 47
50 53 55 57 59 61 62 64
90 92 93 94 96
2 4 6 7 9
22 21 18 16 15 13 10 8
1 2 4 7 10 13 14 17
25 23 21 19 16 14 11 10
36 33 32 31 28 26
57 60 63 65 66 69 71
35 36 38 39 41 43 44
13 14 16 17 18 20
33 30 29 27 24 23 21
48 45 44 42 41 39
97 95 92 91 88 85
83 80 79 76 73
73 71 69 67 64 61 58 56
78 79 82 85 86 87 90
12 14 17 19 21 22 23 25
25 26 29 31 32 33 36 38
29 26 25 24 23 22 20
48 46 44 41 38 35 33 32
74 73 72 71 70 67 64
90 89 86 83 82 81 79 76
30 32 33 36 39 40 41 43
12 14 16 18 21 22 23 25
21 19 18 16 15
41 42 43 44 45 48 50 53
25 27 30 31 34
29 31 34 35 36 39 40 43
78 75 74 73 70 67
85 84 82 81 80 79
77 76 75 73 72 69 66 65
22 24 26 27 30 33 35 37
11 12 14 17 20 22 23
78 79 81 82 84
62 65 68 69 71 74
30 31 34 37 39 42 43 46
22 23 25 28 31 32 35 36
2 4 6 7 8
14 11 8 7 6 5 4 1
44 45 48 49 51 54
37 38 39 42 45 46
98 96 94 93 90 87
51 52 54 56 57 58
13 14 15 18 20 22
49 46 45 44 43
60 61 62 63 64 67
90 87 84 82 81
62 60 59 58 55 53 50
92 91 89 87 85 82
66 64 63 60 58 55 54 53
84 86 88 90 93 94 96
23 24 25 27 28 30 32 35
55 58 59 61 62 63 64 65
52 53 54 56 57 58 61
84 83 81 80 78 75
29 27 24 21 20 19 17
45 46 47 49 50 51 53 54
34 36 38 40 41 42 44
74 76 78 81 84 85 87
6 7 10 13 14 17 20 21
37 35 32 31 29
77 79 80 83 86 87
84 83 81 79 76 75
31 32 33 35 36 39
47 50 51 53 55 57 60
6 9 10 11 14 16 18
22 23 26 28 29 30 33 35
95 93 92 89 88 86
30 33 34 36 37 39 41
59 56 54 52 51 49 46 44
49 50 53 56 57 58
10 12 13 14 17
77 78 81 83 85 87 90 93
89 88 85 83 81
97 96 94 93 92 91 90 88
40 42 43 45 48 49
96 93 90 87 84
46 49 52 53 54 55 57
70 67 65 64 63
60 58 55 54 53 51 49
78 80 83 86 89 91 93
42 39 37 35 32 31 30
69 71 74 77 79
46 48 50 52 54
1 2 4 5 8 10
15 12 11 9 7
70 72 75 76 79 82 85
66 67 70 72 74 75
3 6 8 10 13 16
5 8 11 13 15 18
83 82 81 78 76 75
22 21 18 16 14 11
86 83 82 81 80 78 77 76
99 98 95 92 90 88 86 84
69 67 65 64 63
14 15 18 21 23 25 27
77 78 80 81 84 86 87
97 94 91 88 85 84 82 81
7 8 9 10 11
42 43 44 46 48 51
97 95 92 91 90 87
81 82 83 84 86 87 89 92
69 70 71 74 77 80 82 83
68 70 73 76 79
29 26 23 22 20 17
4 6 8 11 13 14
36 38 39 40 41 42
4 5 7 9 11 14 17 20
35 33 30 28 27 24 21
59 61 62 63 65 68 70
5 8 11 13 14 15 18
56 58 59 62 65 67 70
65 66 69 70 71 74 75
15 17 18 19 22 23
46 45 43 42 39 37 34 31
65 62 60 59 58
29 26 23 21 19 18
65 63 62 60 58 56 53
16 13 11 8 5 2 1
75 73 71 70 69
89 86 84 82 81 79 78 76
85 83 82 81 80 79 78 77
95 92 90 89 87 85 82
56 59 61 64 66 69
84 85 87 90 91
10 11 13 15 18 19
67 65 63 60 58 56 55
82 85 86 89 91 92
80 79 78 75 72
54 51 49 46 45 43
81 79 76 75 74 71 70 67
58 60 62 63 66 69
85 84 83 81 78 77 75
59 61 63 66 68
38 36 34 31 30
59 62 65 66 68
70 69 67 66 65 64 62 60
31 29 28 25 24 21 19 16
49 46 44 43 40 39 38
20 18 17 16 14 11 9
80 78 75 73 72 70 67
71 70 69 68 66 65 63 61
49 52 54 57 59 60 61
38 41 43 46 47 50
44 41 39 37 34 33
80 79 76 75 74
20 23 26 29 31 32 33 36
76 78 79 82 84
29 27 26 23 21
22 25 28 31 32 33
30 28 26 23 21
10 12 15 18 19 22
57 59 62 64 65 67 68
85 82 80 77 76 73 71 68
29 28 26 24 21 18 17
11 12 15 16 17 20
53 52 49 46 45 42
64 65 67 70 72 74 77 78
4 5 8 11 14
43 44 45 46 47 50 52 55
88 89 92 95 96 99
1 3 6 9 10 13
44 43 41 38 36
72 69 66 63 61 58
69 70 73 76 79 81
85 87 89 90 93 96
64 62 60 57 56 53 51 50
75 77 79 81 83 85
6 8 9 11 12
57 58 61 62 64 67
14 13 10 7 6 4 1
66 69 71 72 75 76
22 20 19 16 14
57 54 53 50 47 46 43 40
18 16 15 14 12 9 8
80 79 76 75 72 69 66 63
9 10 11 12 15 16 17
27 28 31 33 36 39
37 36 33 32 31 30 27
71 69 67 64 61 58 55
46 45 44 41 39 36
21 18 15 13 10
38 36 33 31 30 27 26 24
11 13 16 17 19 21 24 27
21 24 26 29 30 33
94 92 89 88 85 84 81 79
13 15 16 17 18
58 59 61 63 66
23 25 27 29 32 35 36
56 59 61 63 66 69 71 74
84 85 88 89 92 95
78 81 84 87 90 91 93
26 25 22 21 19 16
44 41 38 37 36 33
80 79 78 76 74 72 70 68
27 29 30 32 35 36
75 73 72 70 67 64 61 60
59 61 64 66 69 72 73
39 38 36 34 31 30 27 25
3 4 5 7 10 11
73 75 76 77 80 81
58 60 62 65 68 69
76 79 80 83 86 89 90 91
12 9 6 3 2 1
52 53 54 56 59 61 62 64
72 73 74 77 79
75 72 71 70 69 66 63
48 46 45 44 43 42 39
70 71 73 74 76 79 81 84
90 93 96 98 99
45 46 47 50 53 55 58 61
65 68 69 72 73 75 78 81
47 46 43 40 37 34
13 12 9 8 6 5 4 1
80 77 75 73 72 69 67 66
40 42 44 46 47 50
64 67 68 69 72 73 74 76
1 3 6 8 9 11
24 21 20 18 17
53 56 58 60 62 65
77 78 80 83 85 87
8 10 13 16 19 21 23
30 28 26 24 22
27 25 22 19 18 17
79 76 73 71 69 68 65 62
79 76 74 72 70
72 71 68 66 63 62 61 59
51 53 54 55 57 60 61 64
82 85 86 87 90 91 93 95
32 31 28 26 24 23 22
80 77 74 71 68
63 60 59 56 53
97 96 93 91 88 87
19 17 14 11 9
36 38 41 43 46 49 51 52
53 56 59 62 63 65 66
78 81 83 85 88 89 92 94
11 12 13 16 17
34 37 38 39 42
51 48 45 44 41 40
37 35 34 32 30 29 28 26
79 76 75 73 71 69 67 65
74 71 68 66 63 62 60
87 88 89 90 93 95
33 30 29 26 25 23 22 19
27 26 24 22 19 17 15 13
39 38 37 36 35 33 32
52 49 46 43 42 40 37
33 36 39 41 43 46
29 28 27 26 24 21
66 68 70 71 72 73
38 37 34 33 31
90 89 88 86 84 81 80 78
23 25 27 30 31 33
49 47 44 42 41 38 37
88 85 84 81 78
42 44 45 46 47 49 51
35 37 38 40 41 42 45
94 95 96 97 98
39 36 33 32 29 26
58 60 61 63 64
52 51 49 47 45
23 22 21 19 16 15 13 11
4 5 6 8 9 11 12 15
6 7 9 11 14 16 19 21
30 33 34 35 36
16 19 21 22 24 25 27
64 65 68 71 73 74 76 78
68 66 64 62 59 58 56 54
56 53 51 48 45 44 41 39
17 15 12 11 9 6
53 54 56 57 60 62
21 19 16 15 14 13 10 9
12 11 8 6 4 2
94 92 89 88 86 84 83 82
18 20 23 24 25
39 41 44 47 49 52 55 58
83 82 79 77 74 73 72 69
81 82 83 84 85 87 90
53 54 55 57 58
33 30 27 25 24 22 19
87 88 90 93 94 95 98
77 76 73 72 71 68 65 63
93 92 89 88 87 86
67 64 63 60 58 55
93 91 90 88 86 84
80 82 83 86 88 91 94 97
19 22 25 27 28 31
66 69 70 71 74 76
86 84 81 79 76 74 72 69
21 22 25 27 28
64 67 70 71 72 73
45 42 39 36 35
80 81 84 85 87 90 92 94
41 40 37 34 32 30 29 26
25 22 19 16 13 11 10 7
7 9 12 14 15
51 54 57 59 62 64 67
28 29 31 32 35 36 39
70 73 75 78 80 81 83
88 90 91 93 94 96
58 57 55 52 51 49
31 30 27 24 22 19
43 44 45 47 49 51
84 82 81 79 76 74
58 57 54 52 49 46 45
45 46 48 49 52
16 18 21 23 26 27
24 26 27 29 32 34 35
64 66 68 69 70 72 75 78
28 27 24 23 21 18 16
63 66 67 68 71 73 75
99 96 93 92 89
19 21 23 25 26 27 30 33
40 42 45 47 50 53
61 64 65 66 67
61 58 56 54 52 49
82 83 85 88 89
49 46 45 43 42 41 39
65 66 67 69 72 74 75
92 93 94 96 97
38 39 42 45 47 48 50
49 46 43 42 40
51 50 48 45 42 40 39
66 64 62 60 57 54 51
66 65 63 62 60 58 55
76 78 80 83 86 88
69 70 72 74 75 76
70 72 74 75 76 79
96 95 93 90 87 85 83 81
95 94 91 89 87 84 82
13 15 16 19 20 23
14 16 19 20 23 26
13 10 7 6 5 3
58 57 56 53 52 50
16 14 12 9 7 6 3
74 71 70 69 67 64 62 61
75 72 69 67 66
80 77 76 74 72 71 70 67
79 82 85 88 89 91 93 95
42 41 39 36 34 33 32
10 12 13 16 17
14 12 10 7 5 3
47 49 50 52 53 54 57 59
28 29 32 33 34`

	return res
}

func getTestInput() string {
	res := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	return res
}

func getInput() string {
	return getTestInput()
}
