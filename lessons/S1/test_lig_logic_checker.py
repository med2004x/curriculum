import unittest
from s1_logic_arch import implication, build_truth_table, to_twos_complement, from_twos_complement

class LogicTests(unittest.TestCase):
    def test_implication_truth_table(self):
        self.assertEqual([implication(p,q) for p,q in [(False,False),(False,True),(True,False),(True,True)]],[True,True,False,True])
    def test_twos_complement_round_trip(self):
        for value in (-8, -1, 0, 7):
            self.assertEqual(from_twos_complement(to_twos_complement(value, 4)), value)
    def test_twos_complement_overflow_rejected(self):
        with self.assertRaises(ValueError): to_twos_complement(8, 4)
    def test_malformed_bits_rejected(self):
        with self.assertRaises(ValueError): from_twos_complement("10x1")
    def test_table_has_all_rows(self):
        self.assertEqual(build_truth_table(),[(False,False,True),(False,True,True),(True,False,False),(True,True,True)])

if __name__ == '__main__': unittest.main()
