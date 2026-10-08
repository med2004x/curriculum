import unittest
from lig_multimedia_calcs import raw_rgb_bytes, pcm_bytes

class MultimediaTests(unittest.TestCase):
    def test_rgb_bytes(self): self.assertEqual(raw_rgb_bytes(1920,1080),6220800)
    def test_pcm_bytes(self): self.assertEqual(pcm_bytes(10,44100,2,2),1764000)
    def test_invalid_dimensions_rejected(self):
        with self.assertRaises(ValueError): raw_rgb_bytes(0,10)
    def test_negative_duration_rejected(self):
        with self.assertRaises(ValueError): pcm_bytes(-1,44100,2,2)

if __name__ == '__main__': unittest.main()
