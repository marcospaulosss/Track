# Track
Prova de conhecimentos em Go

## Desafio 1
### Discount Calculator
   Calculo do desconto de um produto.

#### Problemas Identificados

1. **Descontos Inválidos:**
   O código original retornava o valor preço quando lidava com descontos negativos ou maiores que 1. 
   Porem o contexto da funcionalidade é retornar o preço do desconto Para esses casos, o preço deveria ser 0, mas o código retornava o preço original.

#### Soluções Propostas

1. **Validação de Descontos:**
   Corrigimos a função `calculateDiscount` para retornar 0 quando o desconto é negativo ou maior que 1. 
   Isso garante que apenas descontos válidos (entre 0 e 1) são aplicados ao preço.

#### Explicação ao desenvolvedor
   Caro desenvolvedor indentificamos uma falha na função `calculateDiscount` que estava retornando o preço original quando o desconto era negativo ou maior que 1.
   Nestes casos o correto seria retornar 0, pois o desconto é inválido.
   Fizemos a correção e agora a função retorna 0 para descontos inválidos.
   Precisamos que os testes unitários sejam realizados para garantir a qualidade do código e evitarmos futuros problemas.

```bash
go run ./discount_calculator/main.go
```


## Desafio 2
### Calculate Avarage
   Calculo da média de um array de inteiros.

#### Problemas Identificados

1. **Divisão e conversão dos valores:**
   A função calcula corretamente a soma dos números, porem não divide pela quantidade de números para obter a média.
   Além disso, a função retorna um inteiro, o que pode resultar em perda de precisão.

#### Soluções Propostas

1. **Divisão e Conversão:**
   Corrigimos a função `calculateAverage` para dividir o valor total pela quantidade de numeros no array fazendo com que retorne a media corretamente. 
   E também convertemos o valor total e o valor de numeros no array para float64, garantindo a precisão das informações.

#### Explicação ao desenvolvedor
   Caro desenvolvedor indentificamos uma falha na função `calculateAverage` que não estava dividindo o valor total pela quantidade de números no array.
   Além disso, a função estava retornando um inteiro, o que pode resultar em perda de precisão. Lembrando que a tipagem do retorno estava em float64.
   Fizemos a correção e agora a função retorna a média corretamente e com precisão.
   Tenho certeza que a pratica dos testes unitários irá garantir a qualidade do código e evitar futuros problemas.

```bash
go run ./calculate_average/main.go
```

## Desafio 3
### Total dos valores
   Calculo da soma de um array de inteiros.

#### Problemas Identificados

1. **Divisão do total pela quantidade de items no array:**
   A função calcula corretamente a soma dos números, porem divia pela quantidade de números no array trazendo uma media.
   Além disso, o chunkSize não tratava o caso de um array vazio, o que poderia resultar em um deadlock.

#### Soluções Propostas

1. **Remoção da divisão e adição da condicional do chunkSize:**
   Corrigimos a função `totalSum` para a remoção da divisão do valor total pela quantidade de numeros no array fazendo com que retorne o valor total corretamente.
   E também adicionamos uma condicional para se o array estiver vazio, o chunkSize seja 1.

#### Explicação ao desenvolvedor
   Caro desenvolvedor indentificamos uma falha na função `totalSum` que estava dividindo o valor total pela quantidade de números no array.
   Além disso, a função estava travando na execução do go routines no caso do array vir vazio.
   Fizemos a correção e agora a função retorna a soma corretamente.
   Tenho certeza que a pratica dos testes unitários iria encontrar o problema e garantir a qualidade do código.

```bash
go run ./total_sum/main.go
```

## Desafio 4
### Revisão das Estimativas
   Primeiramente, é importante avaliar se as estimativas atuais são realistas. 
   Considerando o contexto e a complexidade das tarefas, aqui estão algumas sugestões de ajustes:

#### Task 1: Criar Rota de Get para Produto

Estimativa Atual: 2 dias Revisão: 2 dias (mantida, pois geralmente é uma tarefa direta)

#### Task 2: Criar Rota de Post e Patch para Usuário

Estimativa Atual: 4 dias Revisão: 5 dias (considerando a complexidade adicional de lidar com diferentes métodos HTTP)

#### Task 3: Criar testes automatizados para fluxo de notificações

Estimativa Atual: 3 dias Revisão: 3 dias (mantida)

#### Task 4: Criar front end da tela de times

Estimativa Atual: 2 dias Revisão: 2 dias (mantida)

#### Task 5: Fazer integração do front com o back na tela de times

Estimativa Atual: 3 dias Revisão: 4 dias (integração pode ser complexa, dependendo das APIs e lógica de negócio)

#### Task 6: Correção de erro 500 retornado na rota de Get dos clientes

Estimativa Atual: 1 dia Revisão: 1 dia (mantida, mas com prioridade para evitar bloqueios)

### Revisão de Alocação de Recursos
Com três desenvolvedores disponíveis, a alocação deve ser otimizada com suas respectivas prioridades:

#### Desenvolvedor 1:

Task 3 (To Do)

Task 6 (In Progress)

total 4 dias

#### Desenvolvedor 2:

Task 5 (To Do)
Task 4 (In Progress)

total 5 dias

#### Desenvolvedor 3:

Task 2 (To Do)
Task 1 (In Progress)

total 6 dias

### Ações Complementares:
- Reuniões diárias para acompanhamento do progresso e ajustes rápidos.
- Revisões de código frequentes para garantir a qualidade e evitar retrabalho.
- Comunicação contínua com o time de produto para alinhar expectativas e gerir riscos.

![Captura de Tela 2024-08-25 às 23.23.12.png](Captura%20de%20Tela%202024-08-25%20%C3%A0s%2023.23.12.png)