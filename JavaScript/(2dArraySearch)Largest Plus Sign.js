/**
 * @param {number} n
 * @param {number[][]} mines
 * @return {number}
 */
 function orderOfLargestPlusSign(n, mines) {
    const arr = [...Array(n)].map(r => Array(n));
    const up = [...Array(n)].map(r => Array(n));
    const down = [...Array(n)].map(r => Array(n));
    const left = [...Array(n)].map(r => Array(n));
    const right = [...Array(n)].map(r => Array(n));
    
    for(let item of mines){
        arr[item[0]][item[1]] = 0;
    }
    
    for(let i = 0 ; i<n ; i++){
        if(arr[0][i]===0)   up[0][i] = 0;
        else                up[0][i] = 1;
        if(arr[n-1][i]===0) down[n-1][i] = 0;
        else                down[n-1][i] = 1;
        if(arr[i][0]===0)   left[i][0] = 0;
        else                left[i][0] = 1;
        if(arr[i][n-1]===0) right[i][n-1] = 0;
        else                right[i][n-1] = 1;
    }
    
    for(let i = 0 ; i<n ; i++){
        for(let j = 1 ; j<n ; j++){
            if(arr[j][i]===0)   up[j][i] = 0;
            else                up[j][i] = up[j-1][i]+1;
            if(arr[n-j-1][i]===0)   down[n-j-1][i] = 0;
            else                    down[n-j-1][i] = down[n-j][i]+1;
            if(arr[i][j]===0)   left[i][j] = 0;
            else                left[i][j] = left[i][j-1]+1;
            if(arr[i][n-j-1]===0)   right[i][n-j-1]=0;
            else                    right[i][n-j-1]=right[i][n-j]+1;
        }
    }
    let ans = 0;
    for(let i = 0 ; i<n ; i++){
        for(let j = 0 ; j<n ; j++){
            let horizontal = Math.min(left[i][j], right[i][j]);
            let vertical = Math.min(up[i][j], down[i][j]);
            ans = Math.max(Math.min(horizontal, vertical), ans);
        }
    }
    return ans;
}