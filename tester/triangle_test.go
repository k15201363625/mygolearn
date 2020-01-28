package main

import (
	"fmt"
	"testing"
)

//必须命名为Test(首字母大写string)
/*
可以测试coverage
命令行：
go test --coverageprofile=c.out
go tool cover -html=c.out

*/

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{30000, 40000, 50000},
	}
	//normal case
	//edge (corner) case
	//chinese support case

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("res:%d,exepected:%d\n", actual, tt.c)
		}
	}
}

//性能测试 benchmark 命名大小写小心
//testing.T testing.B 都是测试签名 signature
//go test -bench .
func BenchmarkTriangle(b *testing.B) {
	fmt.Println(b.N) //自动设置
	a, bb, c := 30000, 40000, 50000
	b.ResetTimer() //隔离准备输入数据的时间
	for i := 0; i < b.N; i++ {
		actual := calcTriangle(a, bb)
		if actual != c {
			b.Errorf("res:%d,exepected:%d\n", actual, c)
		}
	}
}

//性能测试深入
//go test -bench . -cpuprofile cpu.out
//go tool pprof cpu.out
//web
