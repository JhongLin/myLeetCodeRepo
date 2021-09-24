/**
 * @param {string} preorder
 * @return {boolean}
 */
 var isValidSerialization = function(preorder) {
    let tokens = preorder.split(",")

    let len = tokens.length
    if(tokens[0]=='#'){
        if(len==1)  return true
        return false
    }
    let end = len-1
    //let temp = len-1
    let done = false
    while(!done){
        //console.log(tokens)
        done = true
        for(let i = 0; i<end-1; i++){
            if(tokens[i]!='#'&&tokens[i]!='$'){
                //temp = i+2;
                let j = i+1
                while(j<=end&&tokens[j]==='$')    j++
                if(j>end){
                    break
                }
                if(tokens[j]==='#'){
                    let k = j+1
                    while(k<=end&&tokens[k]=='$')   k++
                    if(k>end){
                        break
                    }
                    if(tokens[k]==='#'){
                        done = false
                        tokens[i] = '#'
                        tokens[j] = '$'
                        tokens[k] = '$'
                        i=k
                    }else{
                        i = k-1
                    }
                }else{
                    i = j-1
                }
                
            }
            
        }
        //end = Math.min(temp, end)
    }
    if(tokens[0]==='#'){
        for(let i = 1 ; i<len ;i++){
            if(tokens[i]==='$') continue
            else    return false
        }
        return true
    }
    return false
};