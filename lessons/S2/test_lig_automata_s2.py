import unittest
from s2_automata import accepts

class AutomataTests(unittest.TestCase):
    def setUp(self):
        self.transitions={('q0','0'):'q0',('q0','1'):'q1',('q1','0'):'q0',('q1','1'):'q1'}
    def test_accepts_strings_ending_in_one(self):
        self.assertTrue(accepts(self.transitions,'q0',{'q1'},'101'))
        self.assertFalse(accepts(self.transitions,'q0',{'q1'},'100'))
    def test_empty_word_uses_start_state(self): self.assertTrue(accepts({},'q0',{'q0'},''))
    def test_missing_transition_rejects(self): self.assertFalse(accepts({},'q0',{'q1'},'0'))

if __name__ == '__main__': unittest.main()
