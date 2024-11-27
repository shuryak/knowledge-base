# Вопрос №20. Генерация нормально распределённых случайных величин — метод полярных координат

> Алгоритм метода полярных координат генерирует две независимые нормированные
> нормально распределённые случайные величины $X_1$ и $X_2$.

1. [Генерируются](./question-1.md) равномерно распределённые в
   $\left[0, \, 1\right)$ независимые случайные числа $U_1$ и $U_2$. По ним
   вычисляются случайные величины $V_1 = 2 \cdot U_1 - 1$ и
   $V_2 = 2 \cdot U_2 - 1$, равномерно распределённые в $\left(-1, \, 1\right)$.
2. Вычисляется $S = V_1^2 + V_2^2$. Если верно неравенство $S \geqslant 1$, то
   возврат к шагу 1.
3. Искомые случайные величины
   $X_1 = V_1 \cdot \sqrt{\dfrac{-2 \cdot \ln{S}}{S}}$,
   $X_2 = V_2 \cdot \sqrt{\dfrac{-2 \cdot \ln{S}}{S}}$

---

Покажем, что полученные $X_1$ и $X_2$ независимые, нормированные нормально
распределённые случайные величины.

Шаг 1 генерирует случайную точку $\left(V_1, \, V_2\right)$, равномерно
распределённую в единичном (т.е. $R = 1$) круге. Если перейти к полярным
координатам, то $V_1 = R \cdot \cos{\uptheta}$, $V_2 = R \cdot \sin{\uptheta}$,
а $S = R^2$. Тогда $X_1 = \cos{\uptheta} \cdot \sqrt{-2 \cdot \ln{S}}$,
$X_2 = \sin{\uptheta} \cdot \sqrt{-2 \cdot \ln{S}}$.

Из равномерного закона распределения $V_1$ и $V_2$ следует, что угол $\uptheta$
— равномерно распределённая в $\left[0, \, 2\pi\right)$ случайная величина
(функция плотности вероятности для неё —
$p_\uptheta\left(\varphi\right) = \dfrac{1}{2\cdot\pi}$ в интервале
$\left[0, \, 2\pi\right)$, а вне этого интервала —
$p_\uptheta\left(\varphi\right) = 0$). $S$ — равномерно распределённая в
$\left[0, \, 1\right)$ случайная величина, т.е. её функция распределения
вероятности $F_S\left(x\right) = \mathrm{U}_0\left(x\right)$. Из независимости
$U_1$ и $U_2$ следует независимость $S$ и $\uptheta$.

Обозначим $\rho = \sqrt{-2 \cdot \ln{S}}$ (очевидно, что $\rho$ и $\uptheta$
независимы). Найдём функцию распределения вероятности случайной величины $\rho$:

$$F_\rho\left(r\right) = \mathrm{P}\left(\rho < r\right) = \mathrm{P}\left(\sqrt{-2 \cdot \ln{S}} < r\right) = \mathrm{P}\left(-2 \cdot \ln{S} < r^2\right) = \mathrm{P}\left(S > e^{-\dfrac{r^2}{2}}\right) =$$

$$= 1 - \mathrm{P}\left(S \leqslant e^{-\dfrac{r^2}{2}}\right) = 1 - \mathrm{U}_0\left(e^{-\dfrac{r^2}{2}}\right) = 1 - e^{-\dfrac{r^2}{2}}$$

Перейдём к полярным координатам $\left(r, \varphi\right)$. Тогда
$X_1 = r \cdot \cos\varphi$, $X_2 = r \cdot \sin\varphi$.

Вероятность того, что $\rho$ попадает в бесконечно малый интервал
$\left(r, \, r + \mathrm{d}r\right)$ равна
$\mathrm{d}F_\rho\left(r\right) = r \cdot e^{-\dfrac{r^2}{2}} \, \mathrm{d}r$.
Вероятность того, что $\uptheta$ попадает в бесконечно малый интервал
$\left(\varphi, \, \varphi + \mathrm{d}\varphi\right)$ равна
$\dfrac{1}{2\cdot\pi} \, \mathrm{d}\varphi$. Тогда

<!-- TODO: здесь \mathrm{d} или d? и надо ли \,? -->

$$\mathrm{P}\left(\left\{X_1 < x_1\right\} \cap \left\{X_2 < x_2\right\}\right) = \iint\limits_{\left\{\left(r, \, \varphi\right) | r \cdot \cos\varphi < x_1; r \cdot \sin\varphi < x_2\right\}} r \cdot e^{-\dfrac{r^2}{2}} \cdot \dfrac{1}{2\cdot\pi} \, \mathrm{d}r\mathrm{d}\varphi$$

Переходя от полярных координат к прямоугольным, получаем

$$\mathrm{P}\left(\left\{X_1 < x_1\right\} \cap \left\{X_2 < x_2\right\}\right) = \dfrac{1}{2\cdot\pi} \cdot \iint\limits_{\left\{\left(x, \, y\right) | x < x_1; y < x_2\right\}} e^{-\dfrac{x^2 + y^2}{2}} \, \mathrm{d}x\mathrm{d}y =$$

$= \dfrac{1}{2 \cdot \pi} \cdot \int\limits_{-\infty}^{x_1} e^{-\dfrac{x^2}{2}} \cdot \left(\int\limits_{-\infty}^{x_2} e^{-\dfrac{y^2}{2}} \, \mathrm{d}y\right) \, \mathrm{d}x = \left(\dfrac{1}{\sqrt{2\cdot\pi}} \cdot \int\limits_{-\infty}^{x_1} e^{-\dfrac{x^2}{2}} \, \mathrm{d}x\right) \cdot \left(\dfrac{1}{\sqrt{2\cdot\pi}} \cdot \int\limits_{-\infty}^{x_2} e^{-\dfrac{y^2}{2}} \, \mathrm{d}y\right) =$

$$= \mathrm{N}_0\left(x_1\right) \cdot \mathrm{N}_0\left(x_2\right)$$

По определению функция распределения двумерной случайной величины
$\left(X_1, \, X_2\right)$ —
$F_{\left(X_1, \, X_2\right)}\left(x_1, x_2\right) = \mathrm{P}\left(\left\{X_1 < x_1\right\} \cap \left\{X_2 < x_2\right\}\right)$.

Из полученного выше
$F_{\left(X_1, \, X_2\right)}\left(x_1, x_2\right) = \mathrm{N}_0\left(x_1\right) \cdot \mathrm{N}_0\left(x_2\right)$,
что означает их независимость и подчинённость нормированному нормальному закону.
