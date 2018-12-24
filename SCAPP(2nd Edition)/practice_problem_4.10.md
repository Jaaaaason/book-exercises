Write HCL code describing a circuit that for word inputs A, B, and C selects the median of the three values. That is, the output equals the word lying between the minimum and maximum of the three inputs.



```c
int Median3 = [
		A <= B && A >= C : A;
    	B <= A && B >= C : B;
    	1                : C;
];
```

