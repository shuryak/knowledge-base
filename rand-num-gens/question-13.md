# Вопрос №13. Генерация случайных величин, имеющих распределение Пуассона с $M \left(X\right) = \lambda$

Дискретная случайная величина $X$ распределена по закону Пуассона с параметром
$\lambda > 0$, если она принимает значения $k = 0, 1, 2, \ldots$ с вероятностью

$$\mathrm{P}\left(X=k\right) = \dfrac{\lambda^k \cdot e^{-\lambda}}{k!}, \text{где } k \in \mathbb{Z}^+$$

Математическое ожидание и дисперсия случайной величины $X$ равны параметру
$\lambda$: $M\left(X\right) = D\left(X\right) = \lambda$.

Действительно,
$M\left(X\right) = \lambda e^{-\lambda} \cdot \sum\limits_{k=1}^{\infty} \dfrac{\lambda^{k-1}}{\left(k-1\right)!} = \lambda \cdot e^{-\lambda} \cdot \left(1 + \lambda + \dfrac{\lambda^2}{2!} + \dfrac{\lambda^3}{3!} + \ldots\right) = \lambda \cdot e^{-\lambda} \cdot e^{\lambda} = \lambda$.

> Сумма в скобках представляет собой ряд Маклорена для функции $e^\lambda$.

$M\left(X^2\right) = \sum\limits_{k=0}^{\infty} k^2 \cdot \dfrac{\lambda^k \cdot e^{-\lambda}}{k!} = \sum\limits_{k=1}^{\infty}k \cdot \dfrac{\lambda^k \cdot e^{-\lambda}}{\left(k-1\right)!} = \sum\limits_{m=0}^{\infty} \left(m + 1\right) \cdot \dfrac{\lambda^{m+1} \cdot e^{-\lambda}}{m!}=$

$= \lambda \cdot \left(\sum\limits_{m = 0}^{\infty} m \cdot \dfrac{\lambda^m \cdot e^{-\lambda}}{m!} + \sum\limits_{m = 0}^{\infty} \dfrac{\lambda^m \cdot e^{-\lambda}}{m!}\right) = \lambda \cdot \left(\lambda + e^{-\lambda} \cdot \sum\limits_{m = 0}^{\infty} \dfrac{\lambda^m}{m!}\right) =$

$= \lambda \cdot \left(\lambda + e^{-\lambda} \cdot e^\lambda\right) = \lambda^2 + \lambda$.

$D\left(X\right) = M\left(X^2\right) - M^2\left(X\right) = \lambda^2 + \lambda - \lambda^2 = \lambda$.

Известно, что для неотрицательных независимых случайных величин
$\left\{Y_j\right\}$, одинаково распределённых по показательному закону с
математическим ожиданием $M\left(Y_j\right) = \dfrac{1}{\lambda}$, случайная
величина $X = \max\{i; \sum\limits_{j=1}^{i} Y_j \leqslant 1\}$ ($X$ —
максимальное количество слагаемых $Y_j$, для которых сумма не превосходит
единицы) распределена по закону Пуассона с параметром $\lambda$.

Зная [алгоритм генерации показательного закона](./question-15.md), получаем:

$\sum\limits_{j=1}^{i} Y_j \leqslant 1 \Leftrightarrow \sum\limits_{j=1}^{i} - \dfrac{1}{\lambda} \ln{U_j} \leqslant 1 \Leftrightarrow \left(\ln{U_1} + \ldots + \ln{U_i}\right) \geqslant - \lambda \Leftrightarrow \ln{U_1 \cdot \ldots \cdot U_i} \geqslant -\lambda \Leftrightarrow U_1 \cdot \ldots \cdot U_i \geqslant e^{-\lambda}$

В соответствии с последним соотношением алгоритм генерации случайной величины,
распределённой по закону Пуассона, следующий:

1. Вычисляется $e^{-\lambda}$.
2. Генерируются равномерно распределённые случайные величины $U_1, U_2, \ldots$
   до тех пор, пока их произведение превосходит $e^{-\lambda}$. Как только
   неравенство $U_1 \cdot \ldots \cdot U_k < e^{-\lambda}$ становится верным,
   генерация заканчивается, и очередное значение $X = k - 1$.

> Следует отметить, что эффективность алгоритма (трудоёмкость — временн*ы*е
> затраты) существенно зависит от параметра $\lambda$. Чем больше значение
> параметра $\lambda$, тем больше времени, в среднем, требуется для получения
> очередного сгенерированного значения $X$, так как в этом случае нужно
> генерировать более длинную последовательность величин $U_1, U_2, \ldots$ до
> тех пор, пока их произведение станет меньше очень малого значения
> $e^{-\lambda}$
