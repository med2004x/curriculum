import unittest
from s1_stats_prob import mean, variance_population, conditional_probability

class StatisticsTests(unittest.TestCase):
    def test_mean_and_population_variance(self):
        self.assertEqual(mean([1,2,3,4]),2.5)
        self.assertEqual(variance_population([1,2,3,4]),1.25)
    def test_empty_sample_rejected(self):
        with self.assertRaises(ValueError): mean([])
    def test_conditional_probability(self):
        self.assertAlmostEqual(conditional_probability(.2,.5),.4)
    def test_zero_condition_rejected(self):
        with self.assertRaises(ValueError): conditional_probability(0,0)

if __name__ == '__main__': unittest.main()
