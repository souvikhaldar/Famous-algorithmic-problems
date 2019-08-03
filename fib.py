def fib(n):
    if n <= 2:
        return n
    return fib(n-1)+fib(n-2)

n = int(input())
# print("[]th fibonacci number is {}".format(n,fib(n)))
print("fibonacci number is: ",fib(n))