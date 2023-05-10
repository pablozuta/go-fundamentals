package main  
  
import "fmt" 

func Map[T any, U any](slice []T, fn func(T) U)[]U  {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
  
  
  
func main() { 
	// creamos un slice de integers 
	intSlice := []int {2, 31, 43, 411}
	// asignamos una funcion que toma un numero x como argumento y devuelve ese numero multiplicado por 2 como return y lo asignamos a la variable doubleFn
	doubleFn := func (x int)int  {return 2 * x}
	// usamos el Map generic para crear nuestra estructura
	doubleSlice := Map(intSlice, doubleFn)
	fmt.Println(doubleSlice) 
	
	stringSlice := []string {"Altered Carbon", "VALIS", "Neuromancer"}
	lenFn := func (s string)int  {return len(s)}

	lenSlice := Map(stringSlice, lenFn)
	fmt.Println(lenSlice)
} 