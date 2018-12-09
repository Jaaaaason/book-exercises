### If the calendar time is stored as a signed 32-bit integer, in what year will it overflow?

32 位有符号整数范围为：-2³¹ ～ 2³¹-1，非负数部分为 0 ～ 2³¹-1，也就是最大存储的秒数为 2,147,483,647，一年有 365 x 24 x 60 x 60 = 31,536,000 秒

2,147,483,647 ÷ 31,536,000 ≈ 68.1（年），从 1970 年 1 月 1 日算起的话，2038 年会 overflow，还有 20 年😅