class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t): return False

        freqS1 = dict()
        freqS2 = dict()
        for i in range(len(s)):
            freqS1[s[i]] = 1 + freqS1.get(s[i], 0)
            freqS2[t[i]] = 1 + freqS2.get(t[i], 0)

        return freqS1 == freqS2

            
