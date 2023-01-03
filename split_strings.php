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