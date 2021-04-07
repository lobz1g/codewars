from unittest import TestCase

from ruleDivisibilityBy7.src.python.main import seven


class TestSeven(TestCase):
    def test_seven(self):
        self.assertEqual(seven(1603), (7, 2))
        self.assertEqual(seven(371), (35, 1))
        self.assertEqual(seven(483), (42, 1))
        self.assertEqual(seven(2340029794923400297949), (14, 20))
        self.assertEqual(seven(477557101), (28, 7))
