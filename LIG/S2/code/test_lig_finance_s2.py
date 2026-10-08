import unittest
from lig_finance_s2 import future_value_compound, present_value

class FinanceTests(unittest.TestCase):
    def test_future_value(self): self.assertAlmostEqual(future_value_compound(1000,.05,2),1102.5)
    def test_present_value(self): self.assertAlmostEqual(present_value(1102.5,.05,2),1000)
    def test_zero_periods(self): self.assertEqual(future_value_compound(250,0,0),250)
    def test_invalid_periods_rejected(self):
        with self.assertRaises(ValueError): present_value(100,.05,-1)

if __name__ == '__main__': unittest.main()
