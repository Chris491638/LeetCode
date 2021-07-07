def two_sum(nums: list, target: int) -> list:
    hash_map = {}
    for i, e in enumerate(nums):
        if hash_map.get(target - e) is not None:
            return [hash_map[target - e], i]
        hash_map[e] = i


if __name__ == '__main__':
    print(two_sum([2, 7, 11, 15], 22))
