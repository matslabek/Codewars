<?php
//https://www.codewars.com/kata/52a89c2ea8ddc5547a000863
//The solution is based on the uniquness of every node
function loop_size(Node $node): int {
    $checkedNodes = [];
    $counter = 0;
    do {
        $checkedNodes[] = $node;
        $node = $node->getNext();
		//check if the node has already appeared
        $isFound = array_search($node, $checkedNodes, true);
        $counter++;
		// if yes, we reached the loop
        if ($isFound !== false)
        {
			//index of the first node in the loop = length of the dangling piece
			//thus the loop size can be determined
            return $counter - $isFound;
        }
    } while (true);
    return 0;
}