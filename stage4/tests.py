# coding: utf-8
from hstest.stage_test import StageTest
from hstest.test_case import TestCase, SimpleTestCase


class RegexTest(StageTest):
    m_cases = [
        # stage 1
        ("a", "a",          "True",     "Two identical patterns should return True!"),
        ("a", "b",          "False",    "Two different patterns should not return True!"),
        ("7", "7",          "True",     "Two identical patterns should return True!"),
        ("6", "7",          "False",    "Two different patterns should not return True!"),
        (".", "a",          "True",     "Don't forget that '.' is a wild-card that matches any single character."),
        ("a", ".",          "False",    "A period in the input string is still a literal!"),
        ("", "a",           "True",     "An empty regex always returns True!"),
        ("", "",            "True",     "An empty regex always returns True!"),
        ("a", "",           "False",    "A non-empty regex and an empty input string always returns False!"),
        # stage 2
        ("apple", "apple",  "True",     "Two identical equal-length patterns should return True!"),
        (".pple", "apple",  "True",     "The wild-card '.' should match any single character in a string."),
        ("appl.", "apple",  "True",     "The wild-card '.' should match any single character in a string."),
        (".....", "apple",  "True",     "The wild-card '.' should match any single character in a string."),
        ("", "apple",       "True",     "An empty regex always returns True!"),
        ("apple", "",       "False",    "A non-empty regex and an empty input string always returns False!"),
        ("apple", "peach",  "False",    "Two different patterns should not return True!"),
        # stage 3
        ("le", "apple",     "True",     "If the input string contains the regex, it should return True!"),
        ("app", "apple",    "True",     "If the input string contains the regex, it should return True!"),
        ("a", "apple",      "True",     "If the input string contains the regex, it should return True!"),
        (".", "apple",      "True",     "Even a single wild-card character '.' can produce a match!"),
        ("apwle", "apple",  "False",    "Two different patterns should not return True!"),
        ("peach", "apple",  "False",    "Two different patterns should not return True!"),
        # stage 4
        ('^app', 'apple',           "True",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('le$', 'apple',            "True",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^a', 'apple',             "True",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('.$', 'apple',             "True",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('apple$', 'tasty apple',   "True",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^apple', 'apple pie',     "True",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('^apple$', 'apple',        "True",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('^apple$', 'tasty apple',  "False",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('^apple$', 'apple pie',    "False",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('app$', 'apple',           "False",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^le', 'apple',            "False",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!")
    ]


    def generate(self):
        return [
            SimpleTestCase(
                stdin="{0}|{1}".format(regex, text),
                stdout=output,
                feedback=fb
            ) for regex, text, output, fb in self.m_cases
        ]


if __name__ == '__main__':
    RegexTest('regex.regex').run_tests()
