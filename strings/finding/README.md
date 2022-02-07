# 字符串查找
## 定义
给定一段长度为N的文本和一个长度为M的模式（pattern）字符串，在文本中找到一个和该模式相符的字符串
![](assets/字符串的查找.png)
### 特点
N >> M

## 衍生问题
- 找出文本中所有和该模式相符的子字符串
- 统计该模式在文本中的出现次数
- 找出上下文（和该模式相符的子字符串周围的文字）

## 应用
- 在文本编辑器中查找某个单词
- 在截获的通信内容中寻找某种重要的模式
## 暴力字符串查找算法
在文本中模式可能出现匹配的任何地方检查是否存在

### 实现一
用一个变量`i`跟踪文本，一个变量`j`跟踪模式。对于每个`i`，首先将`j`重置为`0`，并逐渐增大，直至找到了一个不匹配的字符或是模式结束
![](assets/暴力字符串查找.png)

见[violent_search.go:ViolentSearch1(txt, pattern string) int](violent_search.go#ViolentSearch1)

### 实现二
用一个变量`i`跟踪文本，一个变量`j`跟踪模式。首先将`j`重置`0`，对于每个`i`，逐渐增大`j`，直至找到了一个不匹配的字符或是模式结束。如果`i`和`j`位置指向的字符不匹配了，那么需要回退这两个变量的值：将j重新指向模式的开头（0），将`i`指向本次匹配开始位置的下一个位置

见[violent_search.go:ViolentSearch2(txt, pattern string) int](violent_search.go#ViolentSearch2)

## KMP字符串查找算法
## BM字符串查找算法
## RK指纹字符串查找算法

## 总结
<table border="4">
  <caption>各种字符串查找算法的实现的成本总结</caption>
  <thead>
    <tr>
      <th rowspan = "2">算法</th>
      <th rowspan = "2">版本</th>
      <th colspan = "2" style="text-align:center" >操作次数</th>
      <th rowspan = "2" style="text-align:center" >在文本中回退</th>
      <th rowspan = "2" style="text-align:center" >正确性</th>
      <th rowspan = "2" style="text-align:center" >额外的空间需求</th>
    </tr>
    <tr>
      <th >最坏情况</th>
      <th >一般情况</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>暴力算法</td>
      <td style="text-align:center">-</td>
      <td>MN</td>
      <td>1.1N</td>
      <td>是</td>
      <td>是</td>
      <td>1</td>
    </tr>
    <tr>
      <td rowspan = "3">Knuth-Morris-Pratt算法</td>
      <td>完整的DFA</td>
      <td>2N</td>
      <td>1.1N</td>
      <td>否</td>
      <td>是</td>
      <td>MR</td>
    </tr>
    <tr>
      <td>仅构造不匹配的状态转换</td>
      <td>3N</td>
      <td>1.1N</td>
      <td>否</td>
      <td>是</td>
      <td>M</td>
    </tr>
    <tr>
      <td>完整版本</td>
      <td>3N</td>
      <td>N/M</td>
      <td>是</td>
      <td>是</td>
      <td>R</td>
    </tr>
    <tr>
      <td >Boyer-Moore算法</td>
      <td >启发式的查找不匹配的字符</td>
      <td>MN</td>
      <td>N/M</td>
      <td>是</td>
      <td>是</td>
      <td>R</td>
    </tr>
    <tr>
      <td rowspan = "2">Rabin-Karp算法</td>
      <td >拉斯维加斯算法</td>
      <td>7N*</td>
      <td>7N</td>
      <td>是</td>
      <td>是</td>
      <td>1</td>
    </tr>
    <tr>
      <td>蒙特卡洛算法</td>
      <td>7N</td>
      <td>7N</td>
      <td>否</td>
      <td>是*</td>
      <td>1</td>
    </tr>

  </tbody>
  <tfoot>
    <tr>
        <td colspan = "7">* 概率保证，需要使用均匀和独立的散列函数</td>
    </tr>
  </tfoot>
</table>