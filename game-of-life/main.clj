(ns game-of-life
  (:require
   [clojure.pprint :refer [pprint]]))
;; реализация игры в жизнь из книги Clojure Programming



(defn empty-board
  "Создаёт игровое поле с указанной шининой и высотой."
  [w h]
  (vec (repeat w (vec (repeat h nil)))))


(defn populate
  "Включает значение :on в ячейках, определяемых координатами [y, x]."
  [board living-cells]
  (reduce (fn [board coordinates]
            (assoc-in board coordinates :on))
          board
          living-cells))

(def glider (populate (empty-board 6 6) #{[2 0] [2 1] [2 2] [1 2] [0 1]}))

(defn neighbours
  "Ищем соседей. Если позиция 0 везде, значит это наше местоположение.
  А возвращаем мы массив из координат соседей. принимаем координаты
  точки, и считаем для неё координаты соседей."
  [[x y]]
  (for [dx [-1 0 1] dy [-1 0 1] :when (not= 0 dx dy)]
    [(+ dx x) (+ dy y)]))

(defn count-neighbours
  [board loc]
  (count (filter #(get-in board %) (neighbours loc))))

(defn indexed-step
  "Возвращает следующее состояние игрового поля, используя индексы
  для определения координат ячеек, соседних с живыми клетками."
  [board]
  ;; считаем длину и ширину поля
  (let [w (count board)
        h (count (first board))]
    (loop [new-board board x 0 y 0]
      (cond
        (>= x w) new-board
        (>= y h) (recur new-board (inc x) 0)
        :else
        (let [new-liveness
              (case (count-neighbours board [x y])
                2 (get-in board [x y])
                3 :on
                nil)]
          (recur (assoc-in new-board [x y] new-liveness) x (inc y)))))))


(defn indexed-step2
  "Замена loop на reduce"
  [board]
  (let [w (count board)
        h (count (first board))]
    (reduce
     (fn [new-board x]
       (reduce
        (fn [new-board y]
          (let [new-liveness
                (case (count-neighbours board [x y])
                2 (get-in board [x y])
                3 :on
                nil)]
            (assoc-in new-board [x y] new-liveness)))
        new-board (range h)))
     board (range w))))


(defn indexed-step3
  "Сворачиваем вложенный reduce"
  [board]
  (let [w (count board)
        h (count (first board))]
    (reduce
     (fn [new-board [x y]]
       (let [new-liveness
             (case (count-neighbours board [x y])
               2 (get-in board [x y])
               3 :on
               nil)]
         (assoc-in new-board [x y] new-liveness)))
     board (for [x (range h) y (range w)] [x y]))))

(defn window
  "Возвращает ленивую последовательность окон с тремя элементами в каждом,
  цетрами в которых является элемент coll."
  ([coll] (window nil coll))
  ([pad coll]
   (partition 3 1 (concat [pad] coll [pad]))))

(defn cell-block
  "Создаёт последовательности окон 3х3 из троек последовательностей по 3
  элемента в каждой."
  [[left mid right]]
  (window (map vector left mid right)))

(defn liveness
  "Возвращает признак наличия живой клетки (nil или :on) в центральной
  ячейке для выполнения следующего шага."
  [block]
  (let [[_ [_ center _] _] block]
    (case (- (count (filter #{:on} (apply concat block)))
             (if (= :on center) 1 0))
      2 center
      3 :on
      nil)))

(defn- step-row
  "Возвращает следующее состояние центра строки."
  [rows-triple]
  (vec (map liveness (cell-block rows-triple))))

(defn index-free-step
  "Возвращает следующее состояние игрового поля."
  [board]
  (vec (map step-row (window (repeat nil) board))))


(defn step
  "Возвращает следующее состояние мира"
  [cells]
  (set (for [[loc n] (frequencies (mapcat neighbours cells))
             :when (or (= n 3) (and (= n 2) (cells loc)))]
         loc)))


(defn main
  []
  (println "hello world"))


(defn stepper
  "Возвращает функцию step для реализации клеточного автомата.
  neighbours принимает координаты и возвращает упорядоченную коллекцию координат.
  Функция survive? и birth? - это предикаты, проверяющие
  число живых соседей."
  [neighbours birth? survive?]
  (fn [cells]
    (set (for [[loc n] (frequencies (mapcat neighbours cells))
               :when (if (cells loc) (survive? n) (birth? n))]
           loc))))

;; (main)

(comment

  (->> (iterate step #{[2 0] [2 1] [2 2] [1 2] [0 1]})
       (drop 8)
       first
       (populate (empty-board 6 6))
       pprint)



  (let [board (empty-board 3 3 )]
    (populate board #{[2 0] [1 1]}))

  (pprint glider)

  (-> (iterate index-free-step glider) (nth 8) pprint)



(loop [x 10]
  (when (> x 1)
    (println x)
    (recur (- x 2))))

(defn increase [i]
  (if (< i 10)
    (recur (inc i))
    i))
(increase 11)


(for [x [0 1 2 3 4 5]
      :let [y (* x 3)]
      :when (even? y)]
  y)



(window [[5] [4] [1] [2]])

(mapcat reverse [[3 2 1 0] [6 5 4] [9 8 7]])


  )
