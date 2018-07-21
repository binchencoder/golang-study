# Array and Slice benchmark

## Run and result

```
cd ch4

go test -bench ch4/ -benchmem -gcflags "-N -l"
```

These is benchmark result:
```
BenchmarkArray-4          300000              4009 ns/op               0 B/op          0 allocs/op
BenchmarkSlice-4          300000              5962 ns/op            8192 B/op          1 allocs/op
```

> NOTE: 在测试 Array 的时候，用的是4核，循环次数是500000，平均每次执行时间是3637 ns，每次执行堆上分配内存总量是0，分配次数也是0 。而切片的结果就“差”一点，同样也是用的是4核，循环次数是300000，平均每次执行时间是4055 ns，但是每次执行一次，堆上分配内存总量是8192，分配次数也是1 。