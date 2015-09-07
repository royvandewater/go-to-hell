# Challenge 2

A nearly prime number is an integer positive number for
which it is possible to find such primes `P1` and `P2`
that given number `N` is equal to `P1 * P2`. In other
words, it's a nearly prime number if its the multiple of
two primes. The test will call `Run` with a sequence of
positive integers. `Run` must return `true` for all nearly
prime numbers.

### Example

```
1 => false // 1 Is already prime
6 => true  // 2 * 3, where 2 & 3 are both prime
8 => false // 2 * 2 * 2, it took 3 primes
```

[source: acm.hit.edu.cn](http://acm.hit.edu.cn/judge/show.php?Proid=1015)
