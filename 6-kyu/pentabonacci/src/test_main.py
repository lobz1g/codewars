from unittest import TestCase

from pentabonacci.src.main import count_odd_pentaFib


class Test(TestCase):
    def test_count_odd_penta_fib(self):
        self.assertEqual(count_odd_pentaFib(5), 1)
        self.assertEqual(count_odd_pentaFib(10), 3)
        self.assertEqual(count_odd_pentaFib(15), 5)
        self.assertEqual(count_odd_pentaFib(45), 15)
        self.assertEqual(count_odd_pentaFib(68), 23)
        self.assertEqual(count_odd_pentaFib(0), 0)
        self.assertEqual(count_odd_pentaFib(1), 1)
        self.assertEqual(count_odd_pentaFib(2), 1)
