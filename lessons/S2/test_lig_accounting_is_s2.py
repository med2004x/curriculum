import unittest
from s2_accounting_is import validate_transaction, reconcile

class AccountingISTests(unittest.TestCase):
    def test_complete_valid_transaction(self):
        event={'event_id':'INV-1','amount':125.0,'source_document':'invoice-1.pdf','approved_by':'manager'}
        self.assertEqual(validate_transaction(event),(True,[]))
    def test_missing_controls_are_reported(self):
        valid,errors=validate_transaction({'amount':-1})
        self.assertFalse(valid)
        self.assertGreaterEqual(len(errors),3)
    def test_reconciliation_delta(self): self.assertEqual(reconcile(125,120),5)

if __name__ == '__main__': unittest.main()
