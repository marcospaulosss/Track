# Track
Prova de conhecimentos em Go

## Discount Calculator

Este repositório contém o teste pratico da Track na linguagem Go.

### Problemas Identificados

1. **Descontos Inválidos:**
   O código original retornava o valor preço quando lidava com descontos negativos ou maiores que 1. 
   Porem o contexto da funcionalidade é retornar o preço do desconto Para esses casos, o preço deveria ser 0, mas o código retornava o preço original.

### Soluções Propostas

1. **Validação de Descontos:**
   Corrigimos a função `calculateDiscount` para retornar 0 quando o desconto é negativo ou maior que 1. 
   Isso garante que apenas descontos válidos (entre 0 e 1) são aplicados ao preço.

```bash
go run main.go