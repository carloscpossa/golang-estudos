package reflection

import "reflect"

func percorre(x any, fn func(entrada string)) {
	valor := obtemValor(x)

	quantidadeDeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		{
			quantidadeDeValores = valor.NumField()
			obtemCampo = valor.Field
		}
	case reflect.Slice, reflect.Array:
		{
			quantidadeDeValores = valor.Len()
			obtemCampo = valor.Index
		}
	case reflect.Map:
		{
			for _, chave := range valor.MapKeys() {
				percorre(valor.MapIndex(chave).Interface(), fn)
			}
			return
		}
	}

	for i := 0; i < quantidadeDeValores; i++ {
		percorre(obtemCampo(i).Interface(), fn)
	}
}

func obtemValor(x any) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Pointer {
		valor = valor.Elem()
	}

	return valor
}
