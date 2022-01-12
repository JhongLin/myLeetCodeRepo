//https://leetcode.com/problems/robot-bounded-in-circle/

func isRobotBounded(instructions string) bool {
    //forwards := make([]int, 4) //N,W,S,E
    direct := 0 
    position := make([]int, 2)
    for i:=0 ; i<4 ; i++ {
        for _, c := range instructions {
            if c=='G'{
                switch direct {
                case 0:
                    position[1]++
                case 1:
                    position[0]--
                case 2:
                    position[1]--
                case 3:
                    position[0]++
                }
            }else if c=='L'{
                direct = (direct+1)%4
            }else{
                direct = (direct+3)%4
            }
        }
    }
    if position[0]==0 && position[1]==0 {
        return true
    }
    return false
}