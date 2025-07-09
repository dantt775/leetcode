# Documentação das Funções `lengthOfLongestSubstringSlice` e `lengthOfLongestSubstring`

Esta documentação detalha duas funções em Go, `lengthOfLongestSubstringSlice` e `lengthOfLongestSubstring`, que têm como objetivo encontrar o comprimento da **substring mais longa sem caracteres repetidos**.

---

## Função: `lengthOfLongestSubstringSlice`

### O que a função faz?

A função `lengthOfLongestSubstringSlice` recebe um *slice* de strings (`s`) e utiliza uma abordagem de **janela deslizante** (sliding window) para identificar a substring (ou sub-slice, neste contexto) mais longa que não contém strings repetidas. Ela itera sobre o *slice*, expandindo a janela para a direita e ajustando a janela para a esquerda sempre que uma string repetida é encontrada dentro da janela atual.

### Responsabilidade de cada variável

Abaixo, descrevemos o papel de cada variável dentro da função `lengthOfLongestSubstringSlice`:

#### Parâmetro de Entrada
* **`s []string`**: Representa o *slice* de strings no qual a função buscará a substring mais longa sem repetições.

#### Variáveis de Controle da Janela
* **`stringIndexMap map[string]int`**: Um **mapa (hash map)** que armazena a última posição (índice) de cada string encontrada dentro da janela atual.
    * **Chave**: A `string` em si.
    * **Valor**: O `int` representa o **último índice** em que essa string foi vista dentro do *slice* `s`.
    * **Responsabilidade**: Essencial para verificar rapidamente se uma string já existe na janela atual e, em caso afirmativo, qual foi sua última ocorrência, permitindo o ajuste da `windowStartIndex`.
* **`windowStartIndex int`**: Representa o **índice inicial da janela deslizante**.
    * **Responsabilidade**: Define o limite esquerdo da janela atual. Quando uma string repetida é encontrada, este índice é ajustado para a direita para garantir que a janela não contenha repetições.
* **`windowEndIndex int`**: Representa o **índice final da janela deslizante**.
    * **Responsabilidade**: É o índice que itera sobre o *slice* `s`, expandindo a janela para a direita. O `range s` no laço `for` atribui automaticamente o índice atual a esta variável.

#### Variáveis de Processamento e Resultado
* **`maxLength int`**: Armazena o **comprimento máximo** da substring sem repetições encontrada até o momento durante a iteração.
    * **Responsabilidade**: Manter o controle do maior comprimento válido encontrado, que será o valor de retorno da função.
* **`currentString string`**: Armazena a **string no `windowEndIndex` atual**.
    * **Responsabilidade**: Facilita a leitura e manipulação da string que está sendo processada na iteração atual.
* **`repeatedStringIndex int`**: Armazena o **índice da última ocorrência** de `currentString` caso ela já exista no `stringIndexMap`.
    * **Responsabilidade**: Se `found` for verdadeiro e `repeatedStringIndex` for maior ou igual a `windowStartIndex`, isso indica que a string repetida está dentro da janela atual, e `windowStartIndex` precisa ser ajustado.
* **`found bool`**: Indica se a `currentString` **foi encontrada** no `stringIndexMap`.
    * **Responsabilidade**: Ajuda a determinar se a `currentString` já foi vista na janela.
* **`currentLength int`**: Calcula o **comprimento da janela atual** (substring sem repetições).
    * **Responsabilidade**: É usada para comparar com `maxLength` e atualizar `maxLength` se a janela atual for maior. O cálculo é `windowEndIndex - windowStartIndex + 1`.

---

## Função: `lengthOfLongestSubstring`

### O que a função faz?

A função `lengthOfLongestSubstring` é similar à `lengthOfLongestSubstringSlice`, mas opera diretamente em uma **string** (`s`) em vez de um *slice* de strings. A principal diferença é que ela converte a string de entrada para um *slice* de `rune`s (`[]rune`) para garantir o tratamento correto de caracteres Unicode (que podem ocupar mais de um byte). Ela também utiliza a abordagem de janela deslizante para encontrar a substring mais longa sem caracteres repetidos.

### Responsabilidade de cada variável

Abaixo, descrevemos o papel de cada variável dentro da função `lengthOfLongestSubstring`:

#### Parâmetro de Entrada
* **`s string`**: Representa a string na qual a função buscará a substring mais longa sem repetições.

#### Variáveis de Processamento de Caracteres
* **`runeIndexMap map[rune]int`**: Um **mapa (hash map)** que armazena a última posição (índice) de cada `rune` encontrada dentro da janela atual.
    * **Chave**: O `rune` em si.
    * **Valor**: O `int` representa o **último índice** em que esse `rune` foi visto dentro do *slice* `runes`.
    * **Responsabilidade**: Essencial para verificar rapidamente se um `rune` já existe na janela atual e, em caso afirmativo, qual foi sua última ocorrência, permitindo o ajuste do ponteiro `left`.
* **`runes []rune`**: Um **slice de `rune`s** criado a partir da string de entrada `s`.
    * **Responsabilidade**: Permite que a função itere sobre os caracteres Unicode da string de forma correta, garantindo que cada caractere seja tratado individualmente, independentemente do seu tamanho em bytes.

#### Variáveis de Controle da Janela
* **`left int`**: Representa o **índice inicial da janela deslizante**. É análogo a `windowStartIndex` na função anterior.
    * **Responsabilidade**: Define o limite esquerdo da janela atual. Quando um `rune` repetido é encontrado, este índice é ajustado para a direita para garantir que a janela não contenha repetições.
* **`right int`**: Representa o **índice final da janela deslizante**. É análogo a `windowEndIndex` na função anterior.
    * **Responsabilidade**: É o índice que itera sobre o *slice* `runes`, expandindo a janela para a direita.

#### Variáveis de Processamento e Resultado
* **`maxLength int`**: Armazena o **comprimento máximo** da substring sem repetições encontrada até o momento durante a iteração.
    * **Responsabilidade**: Manter o controle do maior comprimento válido encontrado, que será o valor de retorno da função.
* **`currentRune rune`**: Armazena o **`rune` no `right` atual**.
    * **Responsabilidade**: Facilita a leitura e manipulação do `rune` que está sendo processado na iteração atual.
* **`index int`**: Armazena o **índice da última ocorrência** de `currentRune` caso ele já exista no `runeIndexMap`. É análogo a `repeatedStringIndex`.
    * **Responsabilidade**: Se `found` for verdadeiro e `index` for maior ou igual a `left`, isso indica que o `rune` repetido está dentro da janela atual, e `left` precisa ser ajustado.
* **`found bool`**: Indica se o `currentRune` **foi encontrado** no `runeIndexMap`.
    * **Responsabilidade**: Ajuda a determinar se o `currentRune` já foi visto na janela.
* **`currentLength int`**: Calcula o **comprimento da janela atual** (substring sem repetições).
    * **Responsabilidade**: É usada para comparar com `maxLength` e atualizar `maxLength` se a janela atual for maior. O cálculo é `right - left + 1`.