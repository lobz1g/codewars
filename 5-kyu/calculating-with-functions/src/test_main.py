from unittest import TestCase

from calculatingWithFunctions.src.main import seven, times, five, four, plus, nine, eight, minus, three, six, \
    divided_by, two


class Test(TestCase):
    def test_four(self):
        self.assertEqual(four(plus(nine())), 13)

    def test_six(self):
        self.assertEqual(six(divided_by(two())), 3)

    def test_seven(self):
        self.assertEqual(seven(times(five())), 35)

    def test_eight(self):
        self.assertEqual(eight(minus(three())), 5)
