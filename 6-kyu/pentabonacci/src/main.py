def count_odd_pentaFib(n):
    return len(set([x for x in fib(n) if x % 2 == 1]))


def fib(n):
    arr = [0, 1, 1, 2, 4]
    for _ in range(n):
        arr.append(sum(arr[len(arr) - 5:len(arr)]))
    return arr[:n + 1]
