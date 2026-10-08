import unittest
from lig_algebra_s2 import matmul, rank_placeholder

class AlgebraTests(unittest.TestCase):
    def test_matrix_product(self):
        self.assertEqual(matmul([[1,2],[3,4]],[[5,6],[7,8]]),[[19,22],[43,50]])
    def test_incompatible_dimensions_rejected(self):
        with self.assertRaises(ValueError): matmul([[1,2]],[[1,2]])
    def test_rank(self):
        self.assertEqual(rank_placeholder([[1,2,3],[2,4,6],[1,1,0]]),2)
    def test_empty_matrix_rank(self): self.assertEqual(rank_placeholder([]),0)

if __name__ == '__main__': unittest.main()
