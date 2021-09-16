goos: windows
goarch: amd64
pkg: muke_study/coding/course_md/day04_code/job
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkSyncMap_Store-4         	 4322677	       275.4 ns/op
BenchmarkMMap_Store-4            	 6937386	       155.7 ns/op
BenchmarkSyncMap_Load-4          	13899619	        88.09 ns/op
BenchmarkMMap_Load-4             	 4988013	       247.6 ns/op
BenchmarkSyncMap_Delete-4        	100000000	        11.56 ns/op
BenchmarkMMap_Delete-4           	11545576	       101.8 ns/op
BenchmarkSyncMap_LoadOrStore-4   	55966233	        21.94 ns/op
BenchmarkMMap_LoadOrStore-4      	 9022854	       133.0 ns/op
BenchmarkSyncMap_Range-4         	 4537368	       261.5 ns/op
BenchmarkMMap_Range-4            	 3497974	       311.8 ns/op
PASS
ok  	muke_study/coding/course_md/day04_code/job	13.916s
