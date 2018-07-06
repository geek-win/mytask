package tpl

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					return {{ err . "value must equal " (durStr $r.Const) }}
				}
			{{ end }}

// ptypes.Duration()可以生成一个time.Duration变量，将从{{assessor .}}获取的值，通过这个函数，可以转化成Duration变量
			{{  if $r.Lt }}  lt  := {{ durLit $r.Lt }};  {{ end }} // durLit是将传入的Duration变量转成格式化成相应格式
			{{- if $r.Lte }} lte := {{ durLit $r.Lte }}; {{ end }} //这里是将r的所有的字段，格式化成相应的格式
			{{- if $r.Gt }}  gt  := {{ durLit $r.Gt }};  {{ end }} // 判断是否设置了这个校验，直接调用字段名，如果是获取校验字段的值，需要使用Get来获取。
			{{- if $r.Gte }} gte := {{ durLit $r.Gte }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
						if dur <= gt || dur >= lt {
							return {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							return {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							return {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
						}
					{{ else }}
						if dur >= lt && dur < gte {
							return {{ err . "value must be outside range [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						return {{ err . "value must be less than " (durStr $r.GetLt) }}
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							return {{ err . "value must be inside range (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
						}
					{{ else }}
						if dur > lte && dur <= gt {
							return {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							return {{ err . "value must be inside range [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
						}
					{{ else }}
						if dur > lte && dur < gte {
							return {{ err . "value must be outside range (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						return {{ err . "value must be less than or equal to " (durStr $r.GetLte) }}
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					return {{ err . "value must be greater than " (durStr $r.GetGt) }}
				}
			{{ else if $r.Gte }}
				if dur < gte {
					return {{ err . "value must be greater than or equal to " (durStr $r.GetGte) }}
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					return {{ err . "value must be in list " $r.In }}
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					return {{ err . "value must not be in list " $r.NotIn }}
				}
			{{ end }}
`
