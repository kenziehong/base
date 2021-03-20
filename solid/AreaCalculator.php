<?php

class AreaCalculator {
  protected $shapes;

  public function __construct(array $shapes = []) {
    $this->shapes = $shapes;
  }

  public function sum() {
    $shapes = $this->shapes;
    
    if (empty($shapes)) {
      return 0;
    }
    
    $area = []; 
    foreach ($shapes as $shape) {
      switch (is_a($shape)) {
        case 'Square':
          $area[] = pow($shape->length, 2);
          break;
        case 'Circle':
          $area[] = pi() * pow($shape->radius, 2);
          break;
      }
    }

    return array_sum($area);
  }

  public function output() {
    return 'Sum of the areas of provided shapes: '. $this->sum();
  }
}