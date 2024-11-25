# Вопрос №18. Генерация нормально распределённых случайных величин с заданной корреляцией

Если требуется сгенерировать нормально распределённые случайные величины с
заданными

- математическими ожиданиями $M\left(X_1\right) = a_1$,
  $M\left(X_2\right) = a_2$
- дисперсиями $D\left(X_1\right) = \sigma_1^2$,
  $D\left(X_2\right) = \sigma_2^2$,
- коэффициентом корреляции $R_{X_1, \, X_2} = \rho$,

то сначала
[генерируются две независимые нормированные нормальные случайные величины](./question-16.md)
$Y_1, Y_2$.

Требуемые случайные величины получим из соотношений:

- $X_1 = \sigma_1 \cdot Y_1 + a_1$,
- $X_2 = \sigma_2 \cdot \left(\rho \cdot Y_1 + Y_2 \cdot \sqrt{1-\rho^2}\right) + a_2$.

Математическое ожидание и дисперсия $X_1$ находятся аналогично
«[Вопрос №17. Генерация нормально распределённых случайных величин с заданными математическим ожиданием и дисперсией](./question-17.md)»
и удовлетворяют требуемым значениям.

Математическое ожидание и дисперсию $X_2$ находим из соотношений

- $M\left(X_2\right) = M\left(\sigma_2 \cdot \left(\rho \cdot Y_1 + Y_2 \cdot \sqrt{1 - \rho^2}\right) + a_2\right) =$

  $= \rho \cdot \sigma_2 \cdot M\left(Y_1\right) + \sigma_2 \cdot \sqrt{1 - \rho^2} \cdot M\left(Y_2\right) +a_2 = 0 + 0 + a_2 = a_2$,

- $D\left(X_2\right) = D\left(\sigma_2 \left(\rho \cdot Y_1 + Y_2 \cdot \sqrt{1 - \rho^2}\right) + a_2\right) =$

  $= \sigma_2^2 \cdot \rho^2 \cdot D\left(Y_1\right) + \sigma_2^2 \left(1 - \rho^2\right) \cdot D\left(Y_2\right) = \sigma_2^2 \cdot \rho^2 + \sigma_2^2 \cdot \left(1 - \rho^2\right) = \sigma_2^2$.

По определению
$R_{X_1, \, X_2} = \dfrac{M\left(\left(X_1 - M\left(X_1\right)\right) \cdot \left(X_2 - M\left(X_2\right)\right)\right)}{\sigma_{X_1} \cdot \sigma_{X_2}}$.

Так как
$\left(X_1 - M\left(X_1\right)\right) \cdot \left(X_2 - M\left(X_2\right)\right) = \sigma_1 \cdot \sigma_2 \cdot \left(\rho \cdot Y_1^2 + Y_1 \cdot Y_2 \cdot \sqrt{1 - \rho^2}\right)$,
то

$M\left(\left(X_1 - M\left(X_1\right)\right) \cdot \left(X_2 - M\left(X_2\right)\right)\right) = \sigma_1 \cdot \sigma_2 \cdot M\left(\rho \cdot Y_1^2 + Y_1 \cdot Y_2 \cdot \sqrt{1 - \rho^2}\right) =$

$= \sigma_1 \cdot \sigma_2 \cdot \left(\rho \cdot M\left(Y_1^2\right) +M\left(Y_1 \cdot Y_2\right) \cdot \sqrt{1 - \rho^2}\right) = \rho \cdot \sigma_1 \cdot \sigma_2$.

---

$M\left(Y_1^2\right) = 1$, т.к. $Y_1$ — нормированная случайная величина. В силу
независимости случайных величин $Y_1, Y_2$ имеем
$M\left(Y_1 \cdot Y_2\right) = R_{Y_1, \, Y_2} = 0$. Это означает, что
$R_{X_1, \, X_2} = \dfrac{\rho \cdot \sigma_1 \cdot \sigma_2}{\sigma_1 \cdot \sigma_2} = \rho$.
