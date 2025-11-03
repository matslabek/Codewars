<?php
//https://www.codewars.com/kata/525caa5c1bf619d28c000335/train/php
function is_solved(array $board): int {
    $sign = $board[1][1];
    //Checking the diagonals
    if ($sign)
    {
        if (($sign == $board[0][0] && $sign == $board[2][2]) || ($sign == $board[2][0] && $sign == $board[0][2]))
        {
            return $sign;
        }
    }
    for ($i = 0; $i <= 2; $i++)
    {
        //Columns
        $col = array_map(fn($row) => $row[$i], $board);
        if (count(array_unique($col)) == 1 && $col[0]){
            return $col[0];
        }
        //Rows
        if (count(array_unique($board[$i])) == 1 && $board[$i][0]){
            return $board[$i][0];
        }
    }
    return in_array(0, array_merge($board[0], $board[1], $board[2])) ? -1 : 0;
}

