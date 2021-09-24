var findLUSlength = function(strs) {
    var seen = {};
    var arr = [];
    var max = -1;
    var index = -1;
    for(var i = 0; i < strs.length; i++){
        seen[strs[i]] = (seen[strs[i]] || 0) + 1;
        if(seen[strs[i]] > 1){
            if(max < strs[i].length){
                max = strs[i].length
                index = i;
            }
        }
    }
    if(index === -1) {
        strs.forEach(el =>{
        if(el.length > max) max = el.length;
		})
        return max;
    }
    for(var i = 0; i < strs.length; i++){
        if(seen[strs[i]] === 1) arr.push(strs[i]);
    }
    max = -1
    for(var i = arr.length - 1; i >= 0; i--){
        var l = arr[i];
        var d = 0;
        for(var j = 0; j < strs[index].length; j++){
            if(strs[index][j] === l[d]){
                d++;
            }
        }
        if(d === l.length){
            var temp = arr[i];
            arr[i] = arr[arr.length - 1];
            arr[arr.length - 1] = temp;
            arr.pop();
        }
    }
    arr.forEach(el =>{
        if(el.length > max) max = el.length;
    })
    return max
};