<p align="center"><img src="assets/art.png" alt="Chzcake" width="300"></p>
<p align="center">декларативный мета-язык с поддержкой плагинов</p>

---

## Оглавление
  * [Примеры синтаксиса](#синтаксис)

## Синтаксис

``` rust
package model

use (
    "golang" as go
    "std"
)

go::struct People {
    pub name: string = "Name"
    pub age: uint {
      std::v(max:150)
    }
}
```
