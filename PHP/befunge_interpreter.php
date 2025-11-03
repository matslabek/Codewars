<?php
//https://www.codewars.com/kata/526c7b931666d07889000a3c/train/php
function interpret(string $code): string {
    //$code = file_get_contents('example.txt');
    $stack = [];
    $plane = array_map(fn($row) => str_split($row), explode("\n", $code));
    $pointer = [0, 0]; // We start in top left corner
    $last_row = count($plane) - 1;
    $possible_directions = ['>', '<', '^', 'v']; // For choosing the random direction
    $direction = '>';
    $symbol = '';
    $output = '';
    while ($symbol !== '@') {
        $symbol = $plane[$pointer[1]][$pointer[0]];
        switch($symbol) {
            case ' ':   // Whitespace does nothing
            case '@':   // End of program <- checked by condition in the while loop
                break;
            // Set Directions
            case '>':
                $direction = '>';
                break;
            case '<':
                $direction = '<';
                break;
            case 'v':
                $direction = 'v';
                break;
            case '^':
                $direction = '^';
                break;
            case '?':
                $direction = $possible_directions[rand(0,3)];
                break;
            // END OF DIRECTIONS
            case '"': //Start string mode
                do {
                    switch ($direction) {
                        case '>':
                            $pointer[0]++;
                            break;
                        case '<':
                            $pointer[0]--;
                            break;
                        case 'v':
                            $pointer[1]++;
                            break;
                        case '^':
                            $pointer[1]--;
                            break;
                    }
                    $symbol = $plane[$pointer[1]][$pointer[0]];
                    if ($symbol !== '"') {
                        $stack[] = ord($symbol); //Push the ASCII value of the char on the stack
                    }
                } while($symbol !== '"');
                break;
            case '#':   // Skip the next sign
                switch ($direction) {
                    case '>':
                        $pointer[0] += 1;
                        break;
                    case '<':
                        $pointer[0] -= 1;
                        break;
                    case 'v':
                        $pointer[1] += 1;
                        break;
                    case '^':
                        $pointer[1] -= 1;
                        break;
                }
                break;
            case '0':
            case '1':
            case '2':
            case '3':
            case '4':
            case '5':
            case '6':
            case '7':
            case '8':
            case '9':
                $stack[] = intval($symbol);
                break;
            // OUTPUT
            case '.':
                $output .= array_pop($stack); // Output the integer on the stack
                break;
            case ',':
                $output .= chr(array_pop($stack)); // Output the matching ASCII char of the integer on the stack
                break;
            case '$': //Discard
                array_pop($stack);
                break;
            // Arithmetic operations
            case '+':
                $sum = array_pop($stack) + array_pop($stack);
                $stack[] = $sum;
                break;
            case '-':
                $a = array_pop($stack);
                $b = array_pop($stack);
                $stack[] = $b - $a;
                break;
            case '*':
                $product = array_pop($stack) * array_pop($stack);
                $stack[] = $product;
                break;
            case '/':
                $a = array_pop($stack);
                $b = array_pop($stack);
                $stack[] = $a ? intdiv($b, $a) : 0;
                break;
            case '%':
                $a = array_pop($stack);
                $b = array_pop($stack);
                $stack[] = $a ? ($b % $a) : 0;
                break;
            case '!':
                $val = array_pop($stack);
                $stack[] = $val ? 0 : 1;
                break;
            case '`':   // Greater than
                $a = array_pop($stack);
                $b = array_pop($stack);
                $stack[] = ($a < $b) ? 1 : 0;
                break;
                // Conditional directions
            case '_':
                $val = array_pop($stack);
                $direction = $val ? '<' : '>';
                break;
            case '|':
                $val = array_pop($stack);
                $direction = $val ? '^' : 'v';
                break;
            case ':':
                if ($stack)
                    {
                        $stack[] = end($stack);
                    }
                else {
                    $stack[] = 0;
                }
                break;
            case '\\':
                $a = array_pop($stack);
                if ($stack) {
                    $b = array_pop($stack);
                    $stack[] = $a;
                    $stack[] = $b;
                } else {
                    $stack[] = $a;
                    $stack[] = 0;
                }
                break;
            case 'p': // Put
                $y = array_pop($stack);
                $x = array_pop($stack);
                $v = array_pop($stack);
                $plane[$y][$x] = chr($v);
                break;
            case 'g': // Get
                $y = array_pop($stack);
                $x = array_pop($stack);
                $stack[] = ord($plane[$y][$x]);
                break;
        }
        switch ($direction) {
            case '>':
                $pointer[0]++;
                break;
            case '<':
                $pointer[0]--;
                break;
            case 'v':
                $pointer[1]++;
                break;
            case '^':
                $pointer[1]--;
                if ($pointer[1] === -1) {
                    $pointer[1] = $last_row;
                }
                break;
        }
    }
    return $output;
}

//Some test cases
var_dump(interpret("08>:1-:v v *_$.@\n  ^    _$>\:^"));
var_dump(interpret(">987v>.v\nv456<  :\n>321 ^ _@"));
var_dump(interpret(">              v\nv\"Hello world!\"<\n> ,,,,,,,,,,,, @"));