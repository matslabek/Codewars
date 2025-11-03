<?php
//https://www.codewars.com/kata/52685f7382004e774f0001f7/train/php
function human_readable(int $seconds): String{
    $hours = floor($seconds / 3600);
    $minutes = floor($seconds / 60) - ($hours * 60);
    $secs = $seconds - (3600 * $hours) - (60 * $minutes);

    $time = [$hours, $minutes, $secs];
    $stringTime = "";
    foreach ($time as $item){
        if ($item < 10){
            $stringTime .= "0";
        }
        $stringTime .= (string) $item . ":";
    }
    return substr($stringTime,0,-1);
}