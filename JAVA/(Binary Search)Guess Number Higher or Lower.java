//https://leetcode.com/problems/guess-number-higher-or-lower/

/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * int guess(int num);
 */

public class Solution extends GuessGame {
    public int guessNumber(int n) {
        long mid = ((long)n+1)/2, upper = (long)n, lower = 1, last=0;
        while(guess((int)mid)!=0){
            last = mid;
             if(guess((int)mid)>0){
                 lower = mid;
                 mid = (upper+lower)/2;
                 if(mid==last)  mid++;
             }else{
                 upper = mid;
                 mid = (upper+lower)/2;
                 if(mid==last)  mid--;
             }
        }
        return (int)mid;
    }
}