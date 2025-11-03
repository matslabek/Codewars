<?php
//https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/php
function solution($str) {
  if (strlen($str) % 2) {
    $str = $str . '_';
  }
  $arr = [];
  for ($i = 0; $i < strlen($str); $i++)
    {
      if ($i % 2 === 0) {
        $arr[] = substr($str, $i, 2);
      }
    }
  return $arr;
}