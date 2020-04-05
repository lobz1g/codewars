from unittest import TestCase

from reduceButGrow.src.python.main import grow


class Test(TestCase):
    def test_grow(self):
        self.assertEqual(grow([1, 2, 3]), 6)
        self.assertEqual(grow([4, 1, 1, 1, 4]), 16)
        self.assertEqual(grow([2, 2, 2, 2, 2, 2]), 64)
