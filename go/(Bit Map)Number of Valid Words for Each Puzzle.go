//https://leetcode.com/problems/number-of-valid-words-for-each-puzzle/

func findNumOfValidWords(words []string, puzzles []string) []int {
    freqWords := make(map[int]int)
    // Count how many words are represented by a specific bitmask
    for _, word := range words {
        // get mask for each word
        bs := bits(word) 
        if _, ok := freqWords[bs]; !ok {
            freqWords[bs] = 1
        } else {
            freqWords[bs] += 1
        }
    }

    result := make([]int, len(puzzles))
    for i, puzzle := range puzzles {
        // get mask for puzzle
        mask := bits(puzzle)
        num := 0 // number of words matched in a puzzle
        
        // set first char bit
        fb := 1 << (puzzle[0] - 'a')
  
        // Iterate subsets of current puzzle's character set
        // cur = (cur - 1) & mask runs through (0..(1<<N)) & mask
        for cur := mask;cur != 0;cur = ((cur-1) & mask) {
            if (cur & fb) == 0 {
                continue
            }
            if v, ok := freqWords[cur]; ok {
                num += v
            }
        }
        result[i] = num
    }
    return result
}


func bits(s string) int {
    var bits int
    for _, l := range s {
        // set the bit in the mask
        bits |= 1 << (l - 'a')
    }
    return bits
}