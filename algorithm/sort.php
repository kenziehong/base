<?php

/**
 * - Chạy vòng lặp từ đầu đến cuối mảng (i=0; i<n; i++)
 * - Tại vòng lặp thứ i, nếu số đứng trước lớn hơn số đứng sau thì đổi chỗ
 * - Sau mỗi vòng lặp, số lớn nhất sẽ trôi xuống phía dưới
 */
function bubbleSort($a) {
    $n = count($a);

    for ($i=0; $i<$n; $i++) {
        
        for ($j=0; $j<$n-$i-1; $j++) {
            if ($a[$j] > $a[$j+1]) {
                $tmp = $a[$j];
                $a[$j] = $a[$j+1];
                $a[$j+1] = $tmp;
            }
        }

        print_r($a);
    }
}

/**
 * - Chạy vòng lặp từ đầu đến cuối mảng (i=1; i<n; i++)
 * - Tại vòng lặp thứ i, xem 0 → i-1 đã được sắp xếp, tìm vị trí thích hợp cho a[i]
 * - Sau vòng lặp thứ i, thì dãy 0->i đã được sắp xếp
 */
function insertionSort($a) {
    $n = count($a);
    
    for ($i=1; $i<$n; $i++) {
        $tmp=$a[$i];
        
        for ($j=$i-1; $j>=0 && $a[$j] > $tmp; $j--) {
             $a[$j+1] = $a[$j];
        }
        
        $a[$j+1] = $tmp;
        print_r($a);
    }
}

/**
 * - Chạy vòng lặp từ đầu đến cuối mảng (i=0; i<n; i++)
 * - Tại vòng lặp thứ i, tìm phần tử nhỏ nhất từ i+1 → n, nếu nhỏ hơn a[i] thì đổi chỗ cho a[j]
 * - Sau vòng lặp thứ i, thì dãy 0->i đã được sắp xếp
 */
function selectionSort($a) {
    $n = count($a);
    
    for ($i=$minIndex=0; $i<$n; $i++) {
        $tmp=$min=$a[$i];
        
        for ($j=$i+1; $j<$n; $j++) {
            if ($a[$j] < $min) {
                $min = $a[$j];
                $minIndex = $j;
            }
        }
        
        if ($min != $tmp) {
            $a[$i] = $min;
            $a[$minIndex] = $tmp;
        }
        
        print_r($a);
    }
}

$a = [5, 3, 2, 7, 8, 3, 1];
