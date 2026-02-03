def factorial_r(x):
    if x == 0:
        return 1
    print(f"x: {x}")
    factorial_r(x - 1)
    return x
