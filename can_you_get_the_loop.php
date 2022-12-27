<?php
function loop_size(Node $node): int {
    $checkedNodes = [];
    $counter = 0;
    do {
        $checkedNodes[] = $node;
        $node = $node->getNext();
        $isFound = array_search($node, $checkedNodes, true);
        $counter++;
        if ($isFound !== false)
        {
            return $counter - $isFound;
        }
    } while ($counter <= 10000);
    return 0;
}