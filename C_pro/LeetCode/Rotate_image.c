
// https://leetcode.com/problems/rotate-image/

void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void rotate(int** matrix, int matrixSize, int* matrixColSize){
    if(matrixSize == 1) return;
    for(int i = 0 ; i<=(matrixSize-1)/2; i++) //pass number
    {
        if(i==(matrixSize-1)/2&&matrixSize%2) return;
        for(int j = i ; j<matrixSize-i-1 ; j++) 
        {
            swap( &matrix[i][j], &matrix[matrixSize-j-1][i]);
            swap( &matrix[matrixSize-j-1][i], &matrix[matrixSize-i-1][matrixSize-j-1]);
            swap( &matrix[matrixSize-i-1][matrixSize-j-1], &matrix[j][matrixSize-i-1]);
        }
    }
}