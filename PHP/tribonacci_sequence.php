<?php
//https://www.codewars.com/kata/556deca17c58da83c00002db/train/php
function tribonacci($signature, $n) {
    if (!$n) {
        return array();
    }
    $tribonacciArr = $signature;
    for ($i = 3; $i < $n; $i++) {
        $tribonacciArr[] = array_sum(array_slice($tribonacciArr, $i - 3, 3));
    }
    return array_slice($tribonacciArr, 0, $n);
}
