def merge(left: list, right: list) -> list:
    result = []
    while len(left) > 0 and len(right) > 0:
        if left[0] > right[0]:
            result.append(right[0])
            right = right[1:]
        else:
            result.append(left[0])
            left = left[1:]

    if len(left) > 0:
        result.extend(left)
    if len(right) > 0:
        result.extend(right)

    return result


def merge_sort(lst: list):
    length = len(lst)
    if length <= 1:
        return lst

    mid = int(length / 2)
    return merge(merge_sort(lst[:mid]), merge_sort(lst[mid:]))

