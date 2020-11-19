def quick_sort(lst: list):
    if len(lst) <= 1:
        return
    # получаем индекс вокруг которого будем производить сортировку дальше
    pivot = partition(lst)

    quick_sort(lst[:pivot])
    quick_sort(lst[pivot:])


def partition(lst: list) -> int:

    pivot = lst[int(len(lst) / 2)]

    left = 0
    right = len(lst) - 1

    while True:
        while lst[left] < pivot:
            left += 1

        while lst[right] > pivot:
            right -= 1

        if left >= right:
            return right

        lst[left], lst[right] = lst[right], lst[left]

