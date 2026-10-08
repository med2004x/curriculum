import unittest
from lig_analysis import derivative_polynomial, evaluate_polynomial

class AnalysisTests(unittest.TestCase):
    def test_derivative_highest_power_first(self):
        self.assertEqual(derivative_polynomial([3, 2, 1]), [6, 2])
    def test_constant_derivative(self):
        self.assertEqual(derivative_polynomial([7]), [0])
    def test_horner_value(self):
        self.assertEqual(evaluate_polynomial([2, -3, 4], 3), 13)
    def test_empty_polynomial_rejected(self):
        with self.assertRaises(ValueError): evaluate_polynomial([], 2)

if __name__ == '__main__': unittest.main()
