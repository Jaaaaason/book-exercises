Experiment to measure the dif ference in running time bet ween our potentially inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)

以下为 benchmark 的结果：

```bash
BenchmarkEcho1-4   	100000000	        16.9 ns/op
BenchmarkEcho3-4   	200000000	        7.41 ns/op
```

其中 BenchmarkEcho1 对应书中的 echo1，BenchmarkEcho3 对应书中的 echo3（使用 strings.join 的那个），结果显示使用 strings.join 的那个执行一次大概 7.41 ns，而另一个执行一次大概 16.9 ns，显然使用 strings.join 的更快些。

