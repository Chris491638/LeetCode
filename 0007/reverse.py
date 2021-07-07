def reverse(x: int) -> int:
    if x < 0:
        num = -int(str(x)[1:][::-1])
    else:
        num = int(str(x)[::-1])
    return num if -2 ** 31 < num < 2 ** 31 - 1 else 0


if __name__ == '__main__':
    assert reverse(123) == 3215, '321'
