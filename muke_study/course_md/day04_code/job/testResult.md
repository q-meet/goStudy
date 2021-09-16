goos: windows
goarch: amd64
pkg: muke_study/coding/course_md/day04_code/job
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkSyncMap_Store-4         	 4385259	       268.1 ns/op
BenchmarkMMap_Store-4            	 7843014	       154.1 ns/op
BenchmarkSyncMap_Load-4          	13031692	        95.96 ns/op
BenchmarkMMap_Load-4             	 4553112	       264.2 ns/op
BenchmarkSyncMap_Delete-4        	112388860	         9.171 ns/op
BenchmarkMMap_Delete-4           	11110904	       109.3 ns/op
BenchmarkSyncMap_LoadOrStore-4   	53512195	        22.37 ns/op
BenchmarkMMap_LoadOrStore-4      	 8759206	       146.5 ns/op
BenchmarkSyncMap_Range-4         	 3819045	       264.4 ns/op
BenchmarkMMap_Range-4            	 3896084	       312.2 ns/op
PASS
ok  	muke_study/coding/course_md/day04_code/job	16.688s
