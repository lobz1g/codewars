import math


def zero(operation=None):
    n = 0
    return result(n, operation)


def one(operation=None):
    n = 1
    return result(n, operation)


def two(operation=None):
    n = 2
    return result(n, operation)


def three(operation=None):
    n = 3
    return result(n, operation)


def four(operation=None):
    n = 4
    return result(n, operation)


def five(operation=None):
    n = 5
    return result(n, operation)


def six(operation=None):
    n = 6
    return result(n, operation)


def seven(operation=None):
    n = 7
    return result(n, operation)


def eight(operation=None):
    n = 8
    return result(n, operation)


def nine(operation=None):
    n = 9
    return result(n, operation)


def result(n, operation):
    return n if operation is None else math.floor(eval(str(n) + operation))


def plus(n):
    return '+' + str(n)


def minus(n):
    return '-' + str(n)


def times(n):
    return '*' + str(n)


def divided_by(n):
    return '/' + str(n)
