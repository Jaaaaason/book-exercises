Write an HCL expression for a signal xor, equal to the EXCLUSIVE-OR of inputs a and b. What is the relation between the signals xor and eq defined above ?

```shell
bool xor = (!a && b) || (a && !b);
```

and the relation between the signals xor and eq is:

```shell
xor = !eq
```

