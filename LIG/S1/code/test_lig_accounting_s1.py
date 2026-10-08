import unittest
from lig_accounting_s1 import total_debit, total_credit, balanced

class AccountingTests(unittest.TestCase):
    def test_totals_and_balanced_journal(self):
        rows=[('Cash',100,0),('Capital',0,100)]
        self.assertEqual(total_debit(rows),100)
        self.assertEqual(total_credit(rows),100)
        self.assertTrue(balanced(rows))
    def test_unbalanced_journal(self):
        self.assertFalse(balanced([('Cash',100,0),('Capital',0,90)]))
    def test_empty_journal_is_balanced(self):
        self.assertTrue(balanced([]))

if __name__ == '__main__': unittest.main()
