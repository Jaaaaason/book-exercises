### If the process time is stored as a signed 32-bit integer, and if the system counts 100 ticks per second, after how many days will the value overflow?

32 位有符号整数范围为：-2³¹ ～ 2³¹-1，非负数部分为 0 ～ 2³¹-1，也就是最大存储的值为 2,147,483,647，一天的 process time 为 24 x 60 x 60 x 100 = 8,640,000

2,147,483,647 ÷ 8,640,000 ≈ 248.55 天