# Вопрос №14. Генерация случайной величины, равномерно распределённой в $\left[a, \, b\right)$

Последовательность чисел $\left\{X_i\right\}$ равномерно распределённых в
$\left[a, \, b\right)$ получается линейным преобразованием последовательности
$U_1, U_2, \ldots$, равномерно распределённых в $\left[0, \, 1\right)$:
$X_i = \left(b - a\right) \cdot U_i + a$.

Функция распределения $X_i$:

$F_{X_i}\left(x\right) = \mathrm{P}\left(X_i < x\right) = \mathrm{P}\left(\left(b - a\right) \cdot U_i + a < x\right) = \mathrm{P}\left(U_i < \dfrac{x-a}{b-a}\right) = \mathrm{U_0}\left(\dfrac{x-a}{b-a}\right) = \dfrac{x-a}{b-a}$.

> $\mathrm{U_0}$ — это функция распределения равномерно распределённой на
> интервале $\left[0, \, 1\right)$ случайной величины $U_i$

Последнее равенство выполняется при $0 \leqslant \dfrac{x-a}{b-a} < 1$. Это
неравенство равносильно неравенству $a \leqslant x < b$.

При $x < a$ получим
$F_{X_i}\left(x\right) = \mathrm{P}\left(X_i < x\right) = \mathrm{P}\left(U_i < 0\right) = 0$,
а для $x \geqslant b$ — $F_{X_i} = 1$, так как тогда
$\dfrac{x-a}{b-a} \geqslant 1$.

Эти соотношения показывают, что функция распределения $X_i$ — это функция
распределения равномерной в $\left[a, \, b\right)$ случайной величины:
$F_{X_i} = \mathrm{U_{a, \, b}}\left(x\right)$.
