def is_palindrome(x: int) -> bool:
    # 60ms   14.8M
    # return str(x)[::-1] == str(x)

    #  68ms 14.8M
    if x < 0:
        return False
    cur = 0
    num = x
    while num > 0:
        cur = cur * 10 + num % 10
        num //= 10
    return True if cur == x else False

if __name__ == '__main__':
    assert is_palindrome(12321)
    assert is_palindrome(-123)
    assert is_palindrome(12351)