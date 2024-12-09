# Вопрос №12. Генерация случайных величин, имеющих Геометрическое распределение

Случайная величина $X$ подчинена геометрическому закону распределения, если она
принимает целые числа из интервала $\left[0, \infty\right)$ с вероятностью
$$\mathrm{P}\left(X=k\right) = q^k \cdot p \text{, где } k = 0, 1, 2, \ldots$$

На практике такой величиной является количество независимых испытаний по схеме
Бернулли относительно некоторого события $A$ до первого его появления.

> Другая интерпретация: случайная величина $X$, являющаяся количеством
> независимых испытаний по схеме Бернулли, проведённых после появления события
> $A$ до следующего появления события $A$.

Математическое ожидание случайной величины $X$ —
$M\left(X\right) = \dfrac{q}{p}$, а дисперсия —
$D\left(X\right) = \dfrac{q}{p^2}$.

> Действительно,
>
> $M\left(X\right) = \sum\limits_{k = 0}^{\infty} k \cdot q^k \cdot p = p \cdot q \cdot \sum\limits_{k = 1}^{\infty} k \cdot q^{k-1} = p \cdot q \cdot \left(\sum\limits_{k = 1}^{\infty} q^k\right)' = p \cdot q \cdot \left(\dfrac{q}{1-q}\right)' = \dfrac{p \cdot q}{\left(1-q\right)^2} = \dfrac{q}{p}$.
>
> (использован факт, что $\left(q^k\right)' = k \cdot q^{k-1}$).
>
> $M\left(X^2\right) = \sum\limits_{k = 0}^{\infty} k^2 q^k p = p \cdot q \cdot \sum\limits_{k = 1}^{\infty} k^2 \cdot q^{k-1}$.
>
> Найдём сумму ряда, составив уравнение:
>
> $s = \sum\limits_{k = 1}^{\infty} k^2 \cdot q^{k-1} = \sum\limits_{k = 0}^{\infty} \left(k + 1\right)^2 \cdot q^k = \sum\limits_{k = 0}^{\infty} \left(k^2 \cdot q^k + 2 \cdot k \cdot q^k + q^k\right) =$
>
> $= q \cdot s + 2 \cdot \dfrac{q}{p^2} + \dfrac{1}{1-q} = q \cdot s + \dfrac{2 \cdot q + p}{p^2} \Rightarrow s = \dfrac{2 \cdot q + p}{p^2 \cdot \left(1-q\right)} = \dfrac{q + 1}{p^3}$.
>
> $D\left(X\right) = M\left(X^2\right) - M^2\left(X\right) = \dfrac{q \cdot \left(1 + q\right)}{p^2} - \dfrac{q^2}{p^2} = \dfrac{q}{p^2}$.

Алгоритм генерации геометрического распределения следующий:

1. [Генерируется](./question-1.md) последовательность чисел
   $U_1, U_2, \ldots, U_n$, равномерно распределённых в интервале
   $\left[0, \, 1\right)$.
2. Требуемая целочисленная последовательность $X_1, X_2, \ldots, X_n$ получается
   как наибольшее целое из частного натурального логарифма каждого $U_i$ и
   натурального логарифма $q$ без единицы:
   $X_i = \left\lceil \dfrac{\ln{U_i}}{\ln{q}} \right\rceil - 1$.

---

Рассмотрим случайную величину
$Y = X_i + 1 = \left\lceil \dfrac{\ln{U_i}}{\ln{q}} \right\rceil$. Покажем, что
сгенерированная величина $X_i$ подчинена геометрическому закону распределения
вероятностей. Очевидно, что для этого должно выполняться соотношение
$\mathrm{P}\left(Y=k\right) = q^{k-1} \cdot p$.

Событие $Y = k$ равносильно следующим неравенствам:

$Y = k \Leftrightarrow\left\lceil \dfrac{\ln{U_i}}{\ln{q}} \right\rceil = k\Leftrightarrow k - 1 < \dfrac{\ln{U_i}}{\ln{q}} \leqslant k \Leftrightarrow \text{(так как } \ln{q} < 0 \text{)}$

$\Leftrightarrow \left(k-1\right) \cdot \ln{q} > \ln{U_i} \geqslant k \cdot \ln{q} \Leftrightarrow \ln{q^{k-1}} > \ln{U_i} \geqslant \ln{q^k} \Leftrightarrow q^{k-1} > U_i \geqslant q^k$.

Тогда
$\mathrm{P}\left(Y = k\right) = \mathrm{P}\left(q^k \leqslant U_i < q^{k-1}\right) = q^{k-1} - q^k$.

> Последнее равенство вытекает из равномерного закона распределения $U_i$.

Окончательно,
$\mathrm{P}\left(Y = k\right) = q^{k-1} - q^k = q^{k-1} \cdot \left(1 - q\right) = q^{k-1} \cdot p$.
Следовательно, $X_i = Y - 1$ подчиняется геометрическому закону распределения
вероятностейю

> [!NOTE] Примечание
>
> Для случая $p = 0.5$ на двоичных компьютерах величина $Y$ реализуется
> значением на единицу большем количества старших нулевых разрядом в двоичном
> представлении $U_i$.
>
> Действительно,
> $Y = \left\lceil \dfrac{\ln{U_i}}{\ln{0.5}} \right\rceil = -\left\lceil \log_2 U_i \right\rceil$.
> Если $U_i = 2^{-k} \cdot z$, где $\dfrac{1}{2} \leqslant z < 1$, то
> $\log_2 U_i = -k + \log_2 z$.
>
> Так как $-1 \leqslant \log_2 z < 0$, то
> $-k - 1 \leqslant -k + \log_2 z < -k \Leftrightarrow k < -\log_2 U_i \leqslant k + 1$.
>
> Соответственно, $X$, равное количеству старших нулевых разрядов в двоичном
> представлении $U_i$, будет подчиняться геометрическому закону распределения
> вероятностей.
