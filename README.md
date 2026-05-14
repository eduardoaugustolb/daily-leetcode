# ⚡ daily-leetcode

> Um problema por dia. Todo dia. Sem desculpa.

Não é mais um repositório de estudo. É minha preparação diária pra entrevistas em empresas de tecnologia. Cada solução aqui foi escrita do jeito que eu gostaria de apresentar pra um entrevistador: código limpo, eficiente e com cada escolha justificada.

## 🗂️ Estrutura

```
daily-leetcode/
├── src/
│   ├── dia-1/
│   │   ├── day-one.go
│   │   ├── two-sum.go
│   │   ├── valid-parentheses.go
│   │   └── best-time-to-buy.go
│   └── ...
├── main.go
├── go.mod
└── README.md
```

Cada dia vira um package novo dentro de `src/`. Dentro dele, os problemas são arquivos soltos — sem subpastas, sem firula.

## 📈 Progressão de Dificuldade

A ideia é simples: a cada 2 dias o bicho pega mais.

| Dias | Nível |
|------|-------|
| 1–2  | 🟢 Easy |
| 3–4  | 🟡 Medium |
| 5–6  | 🔴 Hard |
| 7+   | 🔴 Hard + variações |

## ✅ Problemas Resolvidos

### 🟢 Easy

| # | Problema | Solução | Complexidade |
|---|----------|---------|--------------|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [Go](./src/dia-1/two-sum.go) | O(n) / O(n) |
| 20 | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | [Go](./src/dia-1/valid-parentheses.go) | O(n) / O(n) |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [Go](./src/dia-1/best-time-to-buy.go) | O(n) / O(1) |

## 🧠 Padrões & Técnicas

Reconhecer o padrão é metade da solução. A outra metade é implementar sem errar.

| Padrão | Problemas |
|--------|-----------|
| Hash Map | Two Sum |
| Stack | Valid Parentheses |
| Sliding Window / Greedy | Best Time to Buy and Sell Stock |

## 🔑 O que busco em cada solução

- **Clareza antes de esperteza** — se o entrevistador levar mais de 30 segundos pra entender, eu refaço
- **Complexidade justificada** — tempo e espaço sempre na mesa
- **Edge cases explícitos** — array vazio, zero, negativo, duplicata, tudo que pode quebrar
- **Go idiomático** — nada de escrever Go como se fosse Java disfarçado

## 🚀 Como rodar

```bash
git clone https://github.com/eduardoaugustolb/daily-leetcode.git
cd daily-leetcode
go run main.go
```

## 👤 Autor

**Eduardo Augusto** — [github.com/eduardoaugustolb](https://github.com/eduardoaugustolb)

<p align="center">
  <sub>Feito com consistência — um commit por dia.</sub>
</p>
