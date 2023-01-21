<?php
function solution($number)
{
    $romanNumerals = ['M', 'D', 'C', 'L', 'X', 'V', 'I'];
    $values = [1000, 500, 100, 50, 10, 5, 1];

    $romanString = '';
    for ($i = 0; $i < count($romanNumerals); $i += 2) {
        $quotient = intdiv($number, $values[$i]);
        switch ($quotient){
            case 1:
                $romanString .= $romanNumerals[$i];
                break;
            case 2:
                $romanString .= $romanNumerals[$i] . $romanNumerals[$i];
                break;
            case 3:
                $romanString .= $romanNumerals[$i] . $romanNumerals[$i] . $romanNumerals[$i];
                break;
            case 4:
                $romanString .= $romanNumerals[$i + 2] . $romanNumerals[$i + 1];
                break;
            case 5:
                $romanString .= $romanNumerals[$i + 1];
                break;
            case 6:
                $romanString .= $romanNumerals[$i + 1] . $romanNumerals[$i + 2];
                break;
            case 7:
                $romanString .= $romanNumerals[$i + 1] . $romanNumerals[$i + 2] . $romanNumerals[$i + 2];
                break;
            case 8:
                $romanString .= $romanNumerals[$i + 1] .  $romanNumerals[$i + 2] . $romanNumerals[$i + 2] . $romanNumerals[$i + 2];
                break;
            case 9:
                $romanString .= $romanNumerals[$i + 2] . $romanNumerals[$i];
                break;
        }
        $number = $number % $values[i];

    }

    return $romanString;
}