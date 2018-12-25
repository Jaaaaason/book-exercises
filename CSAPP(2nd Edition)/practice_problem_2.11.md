Armed with the function inplace_swap from Problem 2.10, you decide to write code that will reverse the elements of an array by swapping elements from opposite ends of the array, working toward the middle. 

You arrive at the following function: 

```c
void reverse_array(int a[], int cnt) {
    int first, last;
    for (first = 0, last = cnt - 1;
         first <= last;
         first++, last--)
        inplace_swap(&a[first], &a[last]);
}
```

When you apply your function to an array containing elements 1, 2, 3, and 4, you
find the array now has, as expected, elements 4, 3, 2, and 1. When you try it on
an array with elements 1, 2, 3, 4, and 5, however, you are surprised to see that
the array now has elements 5, 4, 0, 2, and 1. In fact, you discover that the code
always works correctly on arrays of even length, but it sets the middle element to
0 whenever the array has odd length.

**A. For an array of odd length cnt = 2k + 1, what are the values of variables
first and last in the final iteration of function reverse_array? **

k+1

**B. Why does this call to function xor_swap set the array element to 0?**

当 first 等于 last 时，&a[first] 和 &a[last] 指向相同的地址，即此时在 inplace_swap 中，

x 和 y 指向相同的地址，*x 和 *y 始终相同

<pre>
Step        *x        *y
------------------------
Initially    a         a
Step1        0         0
Step2        0         0
Step3        0         0
</pre>

**C. What simple modification to the code for reverse_array would eliminate this problem?**

```c
void reverse_array(int a[], int cnt) {
    int first, last;
    for (first = 0, last = cnt - 1;
         first < last;
         first++, last--)
        inplace_swap(&a[first], &a[last]);
}
```





