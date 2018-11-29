(define (f-recurse n)
  ;; 递归计算过程的函数 f
  (if (< n 3)
      n
      (+ (f-recurse (- n 1))
	 (* 2 (f-recurse (- n 2)))
	 (* 3 (f-recurse (- n 3))))))

(define (f-iterate n)
  ;; 迭代计算过程的函数 f
  (define (f-iter a b c counter n)
    (if (not (< counter n))
	c
	(f-iter b c (+ c
		       (* 2 b)
		       (* 3 a)) (+ counter 1) n)))
  (if (< n 3)
      n
      (f-iter 1 2 4 3 n)))
