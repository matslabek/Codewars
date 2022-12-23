<?php
function loop_size(Node $node): int {
  $checkedNodes = [];
  $counter = 0;
  do {
      $checkedNodes[] = $node;
      $isFound = array_search($node, $checkedNodes, true);
        if($counter === 0)
        {
        //var_dump($node);
        var_dump($checkedNodes);
        echo 'CHECKED';
        var_dump($isFound);
        echo('FOUND');
        }
      $node = $node->getNext();
      //$isFound = array_search($node, $checkedNodes, true);
      $counter++;
      if ($isFound !== false)
         {
           return $counter - $isFound; 
         }
  } while ($counter <= 10000);
  echo count($checkedNodes);
  return 0;
}