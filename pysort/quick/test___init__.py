from unittest import TestCase

from . import quick_sort


class Test(TestCase):
    def test_quick_sort(self):
        to_sort = [119, 281, 762, 21, 1, 99, 0, 615, 2275, 8, 4, 2]
        sorted_ = [0, 1, 2, 4, 8, 21, 99, 119, 281, 615, 762, 2275]

        quick_sort(to_sort)
        self.assertListEqual(to_sort, sorted_)
