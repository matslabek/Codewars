<?php
//https://www.codewars.com/kata/514b92a657cdc65150000006/train/php
function solution($number){
    if (! $number) return 0;
    $multiples = [];
    for ($i = 3; $i < $number; $i += 3 ) {
        $multiples[] = $i;
    }
    for ($j = 5; $j < $number; $j += 5) {
        $multiples[] = $j;
    }
    return array_sum(array_unique($multiples));
}